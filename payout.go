package lazerpay

import (
	"fmt"
	"net/http"
)

type PayoutService service


type CryptoPayoutOptions struct {
	Reference       string              `json:"reference"`
	Amount			int                  `json:"amount"`
	Recipient		string		         `json:"recipient"`
	Coin			string		         `json:"coin"`
	MetaData		map[string]any    `json:"metadata"`
	Blockchain		string		         `json:"blockchain"`
}

type payoutResponse struct {
	apiStatus
}

const (
	cryptoPayoutEndponit 	= "api/v1/crypto/payouts/initiate"
	bankPayoutEndponit 		= "api/v1/bank/payouts"
)

func (ps PayoutService) Crypto(options CryptoPayoutOptions) (*payoutResponse, error ) {
	url := fmt.Sprintf(ps.baseUrl + "%s", cryptoPayoutEndponit)

	var payout = new(payoutResponse)
	req, err := ps.Client.newRequest(http.MethodPost, url, options, ps.Client.secretKey)
	if err != nil {
		return nil, err
	}

	if err := do(req, &payout); err != nil {
		return nil, err
	}

	return payout, nil
}


type BankPayoutdata struct {
	BankName string 		         `json:"bankName"`
	AccountName string    `json:"accountName"`
	AccountType string  `json:"accountType"`
	RoutingNumber string  `json:"routingNumber"`
	Currency  string  `json:"currency"`
	Country  string  `json:"country"`
	State  string  `json:"state"`
	City  string  `json:"city"`
	Address  string  `json:"address"`
	PostalCode  string  `json:"postalCode"`
	Network  string  `json:"network"`
	Default  bool  `json:"default"`
	Deleted  bool  `json:"deleted"`
	Street  string  `json:"strret"`
	Id  string  `json:"id"`
	Status  string  `json:"status"`
	CreatedAt  string  `json:"createdAt"`
	UpdatedAt  string  `json:"updatedAt"`
}

type bankPayoutResponse struct {
	apiStatus
	Data  BankPayoutdata 	`json:"data"`
}

type BankOptions struct {
	BankName  		string    `json:"bank_name"`
	BankCode  		string    `json:"bank_code"`
	AccountName  	string    `json:"account_name"`
	AccountNumber   string    `json:"account_number"`
	Currency  		string    `json:"currency"`
	Country  		string    `json:"country"`
	Default  		bool      `json:"default"`
}

// Create: adds a new payout bank account
// @params {options}: required options to create bank payout
func (ps *PayoutService) Create(options *BankOptions) (*bankPayoutResponse, error ) {
	url := fmt.Sprintf(ps.baseUrl + "%s", bankPayoutEndponit)
	
	var bank = new(bankPayoutResponse)
	req, err := ps.Client.newRequest(http.MethodPost, url, options, ps.Client.secretKey)
	if err != nil {
		return nil, err
	}

	if err := do(req, &bank); err != nil {
		return nil, err
	}

	return bank, nil
}

type PayoutOptions struct {
	// The id of the bank to payout
	BankId        string             `json:"bank_payout_id"`
	// The coin to payout e.g "USDT"
	Coin            	string             `json:"coin"`
	// The amount to be payed out
	Amount            	int64              `json:"amount"`
}

// Initate: initates a new payout
// @params {options}: required options to initiate a bank payout
func (ps *PayoutService) Initiate(options *PayoutOptions) (*payoutResponse, error ) {
	url := fmt.Sprintf(ps.baseUrl + "%s/initiate", bankPayoutEndponit)
	
	var payout = new(payoutResponse)
	req, err := ps.Client.newRequest(http.MethodPost, url, options, ps.Client.secretKey)
	if err != nil {
		return nil, err
	}

	if err := do(req, &payout); err != nil {
		return nil, err
	}

	return payout, nil
}
// Update: updates an existing bank payout
// @params {options}: required options to create bank payout
func (ps *PayoutService) Update(identifier string, options *BankOptions) (*payoutResponse, error ) {
	url := fmt.Sprintf(ps.baseUrl + "%s/%s", bankPayoutEndponit, identifier)
	
	var payout = new(payoutResponse)
	req, err := ps.Client.newRequest(http.MethodPut, url, options, ps.Client.secretKey)
	if err != nil {
		return nil, err
	}

	if err := do(req, &payout); err != nil {
		return nil, err
	}

	return payout, nil
}


// Delete: deletes one or more existing bank payout using payout ids
// @params {identifiers}: An array of string containing the ids
func (ps *PayoutService) Delete(identifiers []string) (*payoutResponse, error ) {
	url := fmt.Sprintf(ps.baseUrl + "%s", bankPayoutEndponit)
	payload := map[string]any{
		"ids": identifiers,
	}
	var payout = new(payoutResponse)
	req, err := ps.Client.newRequest(http.MethodDelete, url, payload, ps.Client.secretKey)
	if err != nil {
		return nil, err
	}

	if err := do(req, &payout); err != nil {
		return nil, err
	}

	return payout, nil
}

// GetAll: retrieves all bank payouts
func (ps *PayoutService) GetAll() (*payoutResponse, error ) {
	url := fmt.Sprintf(ps.baseUrl + "%s", bankPayoutEndponit)
	
	var payout = new(payoutResponse)
	req, err := ps.Client.newRequest(http.MethodGet, url, nil, ps.Client.secretKey)
	if err != nil {
		return nil, err
	}

	if err := do(req, &payout); err != nil {
		return nil, err
	}

	return payout, nil
}

 