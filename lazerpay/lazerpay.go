package lazerpay

import (
	// "os"
	// "log"
)

const (
	BaseUrl = "https://api.azerpay.enginnering/"
)

type LazerPay struct {
	SecretKey string
	Env      bool
}

type GetBaseurl interface {
	GetBaseurl() string
}

type GetEndpoint interface {
	EndPoint(Type string, Point string) string
}

var Enpoints = map[string]map[string]string{
	"payments": {
		"initialize": "api/v1/transaction/initialize",
		"verify":     "api/v1/transaction/verify/",
	},

	"payment-link": {
		"link":      "api/v1/payment-links",
	},

	"transfers": {
		"transfer":  "api/v1/transfer",
	},

	"swap": {
		"crypto": "api/v1/swap/crypto",
		"amount": "api/v1/swap/crypto/amount-out",
	},

	"misc": {
		"coins":  "api/v1/coins",
		"rate":   "api/v1/rate?coin=USDT&currency=NGN",
		"balance": "api/v1/wallet/balance",
	},
}

func (LazerPay) GetBaseurl() string {
	return BaseUrl
}

func (LazerPay) Endpoint(Type string, Point string) string {
	return Enpoints[Type][Point]
}

