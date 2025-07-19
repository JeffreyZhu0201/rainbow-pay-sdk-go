package rainbow_pay_sdk_go

import (
	"github.com/JeffreyZhu0201/rainbow-pay-sdk-go.git/internal/models"
	Payservice "github.com/JeffreyZhu0201/rainbow-pay-sdk-go.git/pkg/PayService"
)

func CreateOrder(order models.Order) models.Response {
	return Payservice.CreateOrder(order)
}
