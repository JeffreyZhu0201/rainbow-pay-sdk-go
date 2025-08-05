/*
 * @Author: JeffreyZhu 1624410543@qq.com
 * @Date: 2025-07-27 17:42:29
 * @LastEditors: JeffreyZhu 1624410543@qq.com
 * @LastEditTime: 2025-07-28 15:13:07
 * @FilePath: /workspace/rainbow-pay-sdk-go/pkg/PayService/Notify.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */

package Payservice

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/JeffreyZhu0201/rainbow-pay-sdk-go.git/internal/models"
	"github.com/JeffreyZhu0201/rainbow-pay-sdk-go.git/internal/utils"
)

/**
 * @description: 监听支付消息事件
 * @param {*gin.Context} c
 * @return {*}
 */
func Notify(c *gin.Context) models.Response {
	// 处理支付通知的逻辑

	queryParams := make(map[string]interface{})

	// 获取所有查询参数
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

	var sign string

	switch queryParams["sign_type"] {
	case "MD5":
		sign, _ = utils.SortMapAndSignMD5(queryParams)
	case "RSA":
		if utils.VerifyStringRSA(queryParams, c.Query("sign")) {
			sign = c.Query("sign")
		} else {
			log.Println("invalid sign")
			return models.Response{Code: 400, Message: "Invalid sign"}
		}
	default:
		return models.Response{Code: 400, Message: "不支持该签名类型"}
	}

	if sign != c.Query("sign") {
		log.Println("invalid sign", sign)
		return models.Response{Code: 400, Message: "Invalid sign"}
	}

	// 这里可以根据 out_trade_no 查询订单状态，并更新订单状态为已支付
	// 处理支付成功的逻辑
	log.Printf("支付成功: out_trade_no=%s, pid=%s, trade_no=%s, trade_status=%s, money=%s, sign=%s, sign_type=%s, type=%s, name=%s",
		queryParams["out_trade_no"], queryParams["pid"], queryParams["trade_no"],
		queryParams["trade_status"], queryParams["money"], queryParams["sign"],
		queryParams["sign_type"], queryParams["type"], queryParams["name"])

	return models.Response{Code: http.StatusOK, Message: "success"}
}
