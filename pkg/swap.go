package lazerpay

import (
	"fmt"
	"net/http"
)

type Swapservice service


type SwapOptions struct {
	// A unique reference
	Reference           string          `json:"reference"`
	// Amount to swap
	Amount      		int				`json:"amount"`
	// The crypto to swap from
	FromCoin				string			`json:"from_coin"`
	// The crypto to swap to
	ToCoin					string			`json:"to_coin"`
	// The blockchain you are sending to
	Blockchain			string			`json:"blockchain"`
}

type SwapResponse struct {
	apiStatus
}
type SwapAmountResponse struct {
	apiStatus
	Data   struct{
		FromCoin  	string  `json:"fromCoin"`
		ToCoin  	string  `json:"toCoin"`
	} 						`json:"data"`
}

const (
	swapCryptoEndpoint = "api/v1/swap"
)

// SwapCrypto: swaps one coin to another
// @params {options}: swap options
func (s Swapservice)  SwapCrypto(options SwapOptions) (*SwapResponse, *http.Response, error ) {
	url := fmt.Sprintf(s.baseUrl + "%s/crypto", swapCryptoEndpoint)
	if options.Reference == "" {
		options.Reference = RandomString(12)
	}
	var swapRes = new(SwapResponse)
	resp, err := newRequest(http.MethodPost, url, options, &swapRes)
	if err != nil {
		return nil, nil, err
	}

	return swapRes, resp, nil
}

// GetSwapAmout: retuns the amount to be recieved on swap
// @params {options}: swap options
func (s Swapservice) GetSwapAmount(options SwapOptions) (*SwapAmountResponse, *http.Response, error) {
	url := fmt.Sprintf(s.baseUrl + "%s/crypto/amount-out", swapCryptoEndpoint)
	
	var swapRes = new(SwapAmountResponse)
	resp, err := newRequest(http.MethodPost, url, options, &swapRes)
	if err != nil {
		return nil, nil, err
	}

	return swapRes, resp, nil
}
