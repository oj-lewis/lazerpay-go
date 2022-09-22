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

func (ps PayoutService) Crypto(options CryptoPayoutOptions) (*payoutResponse, *http.Response, error ) {
	url := fmt.Sprintf(ps.baseUrl + "%s", cryptoPayoutEndponit)

	var payoutRes = new(payoutResponse)
	resp, err := newRequest(http.MethodPost, url, options, &payoutRes)
	if err != nil {
		return nil, resp, err
	}

	return payoutRes, resp, nil
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

type BankPayoutOptions struct {
	BankName  		string    `json:"bank_name"`
	BankCode  		string    `json:"bank_code"`
	AccountName  	string    `json:"account_name"`
	AccountNumber   string    `json:"account_number"`
	Currency  		string    `json:"currency"`
	Country  		string    `json:"country"`
	Default  		bool      `json:"default"`
}

// Create: creates a new payout bank account
// @params {options}: required options to create bank payout
func (p *PayoutService) Create(options BankPayoutOptions) (*bankPayoutResponse, *http.Response, error ) {
	url := fmt.Sprintf(p.baseUrl + "%s", bankPayoutEndponit)
	
	var bank = new(bankPayoutResponse)
	resp, err := newRequest(http.MethodPost, url, options, &bank)
	if err != nil {
		return nil, nil, err
	}

	return bank, resp, nil
}

type PayoutOptions struct {
	// The id of the bank to payout
	BankPayoutId        string             `json:"bank_payout_id"`
	// The coin to payout e.g "USDT"
	Coin            	string             `json:"coin"`
	// The amount to be payed out
	Amount            	int64              `json:"amount"`
}

// Initate: initates a new payout
// @params {options}: required options to create bank payout
func (p *PayoutService) Initiate(options PayoutOptions) (*payoutResponse, *http.Response, error ) {
	url := fmt.Sprintf(p.baseUrl + "%s/initiate", bankPayoutEndponit)
	
	var bank = new(payoutResponse)
	resp, err := newRequest(http.MethodPost, url, options, &bank)
	if err != nil {
		return nil, nil, err
	}

	return bank, resp, nil
}


// Update: updates an existing bank payout
// @params {options}: required options to create bank payout
func (p *PayoutService) Update(identifier string, options BankPayoutOptions) (*payoutResponse, *http.Response, error ) {
	url := fmt.Sprintf(p.baseUrl + "%s/%s", bankPayoutEndponit, identifier)
	
	var bank = new(payoutResponse)
	resp, err := newRequest(http.MethodPatch, url, options, &bank)
	if err != nil {
		return nil, nil, err
	}

	return bank, resp, nil
}


// Delete: deletes one or more existing bank payout using payout ids
// @params {identifiers}: An array of string containing the ids
func (p *PayoutService) Delete(identifiers []string) (*payoutResponse, *http.Response, error ) {
	url := fmt.Sprintf(p.baseUrl + "%s", bankPayoutEndponit)
	payload := map[string]any{
		"ids": identifiers,
	}
	var bank = new(payoutResponse)
	resp, err := newRequest(http.MethodDelete, url, payload, &bank)
	if err != nil {
		return nil, nil, err
	}

	return bank, resp, nil
}

// GetAll: retrieves all bank payouts
// @params {identifiers}: An array of string containing the ids
func (p *PayoutService) GetAll() (*payoutResponse, *http.Response, error ) {
	url := fmt.Sprintf(p.baseUrl + "%s", bankPayoutEndponit)
	
	var bank = new(payoutResponse)
	resp, err := newRequest(http.MethodGet, url, nil, &bank)
	if err != nil {
		return nil, nil, err
	}

	return bank, resp, nil
}

 