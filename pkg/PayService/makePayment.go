/*
 * @Author: Jeffrey Zhu 1624410543@qq.com
 * @Date: 2025-03-31 13:41:43
 * @LastEditors: error: error: git config user.name & please set dead value or install git && error: git config user.email & please set dead value or install git & please set dead value or install git
 * @LastEditTime: 2025-07-28 14:44:08
 * @FilePath: \rainbow-pay-sdk-go\pkg\PaySdk\makePayment.go
 * @Description: File Description Here...
 *
 * Copyright (c) 2025 by JeffreyZhu, All Rights Reserved.
 */

package Payservice

import (
	"os"
	"strconv"

	"github.com/JeffreyZhu0201/rainbow-pay-sdk-go.git/internal/models"
	"github.com/JeffreyZhu0201/rainbow-pay-sdk-go.git/internal/utils"
	// "github.com/google/uuid"
)

// CreateOrder handles the creation of a payment order. It takes an Order model as input,
// processes the payment parameters, and makes a request to the payment gateway.
// Returns a Response containing the payment URL if successful, or an error message if failed.
// The response includes:
//   - Code: HTTP status code (200 for success/failure)
//   - Message: Success/failure message
//   - Data: Map containing the payment URL (key: "payment_url")
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

	var signStr string

	switch paymentParams["sign_type"] {
	case "MD5":
		signStr, _ = utils.SortMapAndSignMD5(paymentParams)
	case "RSA":
		signStr, _ = utils.SortMapAndSignRSA(paymentParams)
	default: // 默认使用 MD5 签名
		signStr, _ = utils.SortMapAndSignMD5(paymentParams)
	}
	paymentUrl, err := utils.Post(os.Getenv("PAY_URL") + "?" + signStr)

	if err != nil {
		// 处理错误
		return models.Response{Code: 200, Message: "Order created failed", Data: map[string]interface{}{"payment_url": paymentUrl}}
	}

	if paymentUrl == "" {
		return models.Response{Code: 200, Message: "Order created failed", Data: map[string]interface{}{"payment_url": paymentUrl}}
	}

	return models.Response{Code: 200, Message: "Order created successfully", Data: map[string]interface{}{"payment_url": paymentUrl}}
}
