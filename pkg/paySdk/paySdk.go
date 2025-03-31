package paySdk

import (
	"github.com/JeffreyZhu0201/rainbow-pay-sdk-go.git/pkg/paymentConfig"
)

type PaySdk interface {
	// MakePayment initiates a payment request to the payment gateway.
	makePayment(config paymentConfig.Payment, payment paymentConfig.Payment) (string, error)
}
