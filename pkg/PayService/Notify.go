/*
 * @Author: JeffreyZhu 1624410543@qq.com
 * @Date: 2025-07-27 17:42:29
 * @LastEditors: JeffreyZhu 1624410543@qq.com
 * @LastEditTime: 2025-07-27 17:46:19
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
