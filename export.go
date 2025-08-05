/*
 * @Author: JeffreyZhu 1624410543@qq.com
 * @Date: 2025-07-27 17:13:53
 * @LastEditors: JeffreyZhu 1624410543@qq.com
 * @LastEditTime: 2025-07-27 17:47:23
 * @FilePath: /workspace/rainbow-pay-sdk-go/export.go
 * @Description: 这是默认设置,请设置`customMade`, 打开koroFileHeader查看配置 进行设置: https://github.com/OBKoro1/koro1FileHeader/wiki/%E9%85%8D%E7%BD%AE
 */
package rainbow_pay_sdk_go

import (
	"github.com/JeffreyZhu0201/rainbow-pay-sdk-go.git/internal/models"
	Payservice "github.com/JeffreyZhu0201/rainbow-pay-sdk-go.git/pkg/PayService"
	"github.com/gin-gonic/gin"
)

func CreateOrder(order models.Order) models.Response {
	return Payservice.CreateOrder(order)
}

func NotifyOrder(c *gin.Context) models.Response {
	return Payservice.Notify(c)
}
