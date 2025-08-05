/*
 * @Author: Jeffrey Zhu 1624410543@qq.com
 * @Date: 2025-03-31 13:41:43
 * @LastEditors: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @LastEditTime: 2025-08-05 17:02:20
 * @FilePath: \rainbow-pay-sdk-go\pkg\PaySdk\makePayment.go
 * @Description: File Description Here...
 *
 * Copyright (c) 2025 by JeffreyZhu, All Rights Reserved.
 */
package Payservice

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/JeffreyZhu0201/rainbow-pay-sdk-go.git/internal/models"
	"github.com/JeffreyZhu0201/rainbow-pay-sdk-go.git/internal/utils"
	"github.com/gin-gonic/gin"
	// "github.com/google/uuid"
)

func CreateOrder(order models.Order) models.Response {
	// 处理创建订单的逻辑
	// 这里可以使用 Stripe 或其他支付网关的 SDK 来处理订单创建请求
	// 例如，创建一个订单并返回给前端

	pid, _ := strconv.Atoi(os.Getenv("PAY_PID"))
	paymentParams := make(map[string]interface{})
	paymentParams["pid"] = pid
	paymentParams["notify_url"] = os.Getenv("PAY_NOTIFY_URL")
	paymentParams["return_url"] = os.Getenv("PAY_RETURN_URL")
	paymentParams["sign_type"] = os.Getenv("PAY_SIGN_TYPE")

	paymentParams["out_trade_no"] = order.OutTradeNo
	paymentParams["name"] = order.CommodityName
	// paymentParams["count"] = countUint
	paymentParams["money"] = order.Amount

	// 处理支付请求的逻辑
	// 这里可以使用 Stripe 或其他支付网关的 SDK 来处理支付请求
	// 例如，创建一个支付意图并返回给前端

	signStr, _ := utils.SortMapAndSign(paymentParams)
	paymentUrl, err := utils.Post(os.Getenv("PAY_URL") + "?" + signStr.String())

	if err != nil {
		// 处理错误
		return models.Response{Code: 200, Message: "Order created failed", Data: map[string]interface{}{"payment_url": paymentUrl}}
	}

	if paymentUrl == "" {
		return models.Response{Code: 200, Message: "Order created failed", Data: map[string]interface{}{"payment_url": paymentUrl}}
	}

	return models.Response{Code: 200, Message: "Order created successfully", Data: map[string]interface{}{"payment_url": paymentUrl}}
}

func Notify(c *gin.Context) models.Response {
	// 处理支付通知的逻辑

	queryParams := make(map[string]interface{})

	queryParams["out_trade_no"] = c.Query("out_trade_no")
	queryParams["pid"] = c.Query("pid")
	queryParams["trade_no"] = c.Query("trade_no")
	queryParams["trade_status"] = c.Query("trade_status")
	queryParams["money"] = c.Query("money")
	queryParams["sign"] = c.Query("sign")
	queryParams["sign_type"] = c.Query("sign_type")
	queryParams["type"] = c.Query("type")
	queryParams["name"] = c.Query("name")

	// out_trade_no := c.Query("out_trade_no")

	// 验证签名
	if c.Query("trade_status") != "TRADE_SUCCESS" {
		return models.Response{Code: 400, Message: "Invalid trade_status"}
	}

	_, sign := utils.SortMapAndSign(queryParams)

	if sign != c.Query("sign") {
		log.Println("invalid sign", sign)
		return models.Response{Code: 400, Message: "Invalid sign"}
	}

	return models.Response{Code: http.StatusOK, Message: "success"}
}
