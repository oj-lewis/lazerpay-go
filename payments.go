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
	Network            		string  `json:"network"`
	AcceptPartialPayment 	bool   	`json:"accept_partial_payment"`
}

type PaymentOption struct {
	// A unique reference
	Reference            	string 	`json:"reference,omitempty"`
	// The name of the customer
	CustomerName         	string 	`json:"customer_name"`
	// The email of the customer
	CustomerEmail        	string 	`json:"customer_email"`
	// The to pay with e.g "USDT"
	Coin                 	string 	`json:"coin"`
	// The currency   e.g "USD"
	Currency             	string 	`json:"currency"`
	// The qmount to pay
	FiatAmount            		float64 `json:"amount"`
	// To accept partial payment defaults to false
	AcceptPartialPayment 	bool   	`json:"accept_partial_payment,omitempty"`
	MetaDate				map[string]any `json:"metadata,omitempty"`
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
func (p *PaymentService) Initialize(options *PaymentOption) (*paymentResponse, error ) {
	url := fmt.Sprintf( "%s%s", p.baseUrl, InitializeEndpoint)
	if options.Reference == "" {
		options.Reference = randomString(12)
	}

	var pay = new(paymentResponse)
	req, err := p.Client.newRequest(http.MethodPost, url, options, p.Client.publicKey)
	if err != nil {
		return nil, err
	}
	if err := do(req, &pay); err != nil {
		return nil, err
	}
	

	return pay, nil
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
	Customer 		  		map[string]any 	`json:"customer"`
	PaymentLink				map[string]any 		`json:"paymentLink"`
	PaymentButton  			map[string]any 		`json:"paymentButton"`
	FeeInCrypto 			float64 	`json:"feeInCrypto"`
}

type verifyPaymentResponse struct {
	apiStatus
	Data verifyPaymentData   `json:"data"`
}

// Verify: This verifies is to verify a payments
// @params {identifier}: The unique id of the payment
//@returns: returns a verifypaymentResponse
func (p *PaymentService) Verify(identifier string) (*verifyPaymentResponse, error) {
	url := fmt.Sprintf("%s%s/%s", p.Client.baseUrl, VerifyEndpoint, identifier)
	if identifier == "" {
		return nil, errors.New("please provide a reference or an identifier")
	}

	var verify = new(verifyPaymentResponse)
	req, err := p.Client.newRequest(http.MethodGet, url, nil, p.Client.publicKey)
	if err != nil {
		return nil, err
	}

	if err := do(req, &verify); err != nil {
		return nil, err
	}
	
	return verify, nil
}
