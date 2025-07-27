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
