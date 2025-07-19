package Payservice

// import (
// 	""
// )

type PaySdk interface {
	// MakePayment initiates a payment request to the payment gateway.
	makePayment(payment paymentConfig.Payment) (string, error)
}
