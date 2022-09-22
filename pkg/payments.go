package lazerpay

import (
	"errors"
	"fmt"
	"net/http"
)

type PaymentService service

type paymentData struct {
	Reference            	string 	`json:"reference"`
	BusinessName            string 	`json:"businessName"`
	BusinessEmail           string 	`json:"businessEmail"`
	BusinessLogo            string 	`json:"businessLogo"`
	CustomerName         	string 	`json:"customerName"`
	CustomerEmail        	string 	`json:"customerEmail"`
	Address                 string 	`json:"address"`
	Coin                 	string 	`json:"coin"`
	CryptoAmount            float64 `json:"cryptoAmount"`
	Currency             	string 	`json:"currency"`
	FiatAmount            	float64 `json:"fiatAmount"`
	FeeInCrypto            	float64 `json:"feeInCrypto"`
	Network            		float64 `json:"network"`
	AcceptPartialPayment 	bool   	`json:"accept_partial_payment"`
}

type PaymentOption struct {
	// A unique reference
	Reference            	string 	`json:"reference"`
	// The name of the customer
	CustomerName         	string 	`json:"customerName"`
	// The email of the customer
	CustomerEmail        	string 	`json:"customerEmail"`
	// The to pay with e.g "USDT"
	Coin                 	string 	`json:"coin"`
	// The currency   e.g "USD"
	Currency             	string 	`json:"currency"`
	// The qmount to pay
	Amount            		float64 `json:"cryptoAmount"`
	// To accept partial payment defaults to false
	AcceptPartialPayment 	bool   	`json:"accept_partial_payment"`
}

type paymentResponse struct {
	apiStatus
	Data paymentData   `json:"data"`
}

const (

	InitializeEndpoint = "api/v1/transaction/initialize"
	VerifyEndpoint = "api/v1/transaction/verify"
)

// Initialize: This method initializes a new payment
// @params {options}: Options needed to initialize a payment
//@returns: returns a paymentResponse
func (p *PaymentService) Initialize(options PaymentOption) (*paymentResponse, *http.Response, error ) {
	url := fmt.Sprintf(p.baseUrl + "%s", InitializeEndpoint)
	if options.Reference == "" {
		options.Reference = RandomString(12)
	}

	var pay = new(paymentResponse)
	resp, err := newRequest(http.MethodPost, url, options, &pay)
	if err != nil {
		return nil, nil, err
	}

	return pay, resp, nil
}


type verifyPaymentData struct {
	Id 						string 		`json:"id"`
	Reference      			string 		`json:"reference"`
	SenderAddress			string 		`json:"senderAddress"`
	RecipientAddress 		string 		`json:"recipientAddress"`
	ActualAmount     		float64 	`json:"actualAmount"`
	AmountPaid       		float64    	`json:"amountPaid"`
	AmountPaidFiat      	float64    	`json:"amountPaidFiat"`
	FiatAmount        		int64 		`json:"fiatAmount"`
	AmountRecieved   		float64    	`json:"amountRecieved"`
	AmountRecievedFiat 		float64    	`json:"amountRecievedFiat"`
	Coin					string		`json:"coin"`
	Currency 				string   	`json:"currency"`
	Hash 					string  	`json:"hash"`
	BlockNumber 			int64  		`json:"blockNumber"`
	Type        			string 		`json:"type"`
	AcceptPartialPayment 	bool   		`json:"acceptPartialPayment"`
	Status 					string      `json:"status"`
	Network  				string  	`json:"network"`
	Blockchain 				string  	`json:"blockchain"`
	Customer 		  		struct{
		Id 					string 		`json:"id"`
		CustomerName 		string  	`json:"customerName"`
		CustomerEmail 		string 		`json:"customerEmail"`
		CustomerPhone 		string  	`json:"customerPhone"`
	}  									`json:"customer"`
	PaymentLink				string 		`json:"paymentLink"`
	PaymentButton  			string 		`json:"paymentButton"`
	FeeInCrypto 			float64 	`json:"feeInCrypto"`
}

type verifyPaymentResponse struct {
	apiStatus
	Data verifyPaymentData   `json:"data"`
}

// Verify: This verifies is to verify a payments
// @params {identifier}: The unique id of the payment
//@returns: returns a verifypaymentResponse
func (p *PaymentService) Verify(identifier string) (*verifyPaymentResponse, *http.Response, error) {
	url := fmt.Sprintf(p.baseUrl + "%s/%s", VerifyEndpoint, identifier)
	if identifier == "" {
		return nil, nil, errors.New("please provide a reference or an identifier")
	}

	var verifyRes = new(verifyPaymentResponse)
	resp, err := newRequest(http.MethodGet, url, nil, &verifyRes)
	if err != nil {
		return nil, nil, err
	}

	return verifyRes, resp, nil
}
