package lazerpay

import (
	"fmt"
	"net/http"
)

type Swapservice service


type SwapOptions struct {
	// A unique reference
	Reference   string          `json:"reference"`
	// Amount to swap
	Amount      int				`json:"amount"`
	// The crypto to swap from
	FromCoin	string			`json:"from_coin"`
	// The crypto to swap to
	ToCoin		string			`json:"to_coin"`
	// The blockchain you are sending to
	Blockchain	string			`json:"blockchain"`
}

type SwapResponse struct {
	apiStatus
}
type SwapAmountResponse struct {
	apiStatus
	Data	swapData	`json:"data"`
}

type swapData struct {
	FromCoin  	string  `json:"fromCoin"`
	ToCoin  	string  `json:"toCoin"`
}

const (
	swapCryptoEndpoint = "api/v1/swap"
)

// SwapCrypto: swaps one coin to another
// @params {options}: swap options
func (s Swapservice)  SwapCrypto(options *SwapOptions) (*SwapResponse, error ) {
	url := fmt.Sprintf(s.baseUrl + "%s/crypto", swapCryptoEndpoint)
	if options.Reference == "" {
		options.Reference = randomString(12)
	}
	var swap = new(SwapResponse)
	
	req, err := s.Client.newRequest(http.MethodPost, url, options, s.Client.secretKey)
	if err != nil {
		return nil, err
	}

	if err := do(req, &swap); err != nil {
		return nil, err
	}
	 
	return swap, nil
}

// GetSwapAmout: retuns the amount to be recieved on swap
// @params {options}: swap options
func (s Swapservice) GetSwapAmount(options *SwapOptions) (*SwapAmountResponse, error) {
	url := fmt.Sprintf(s.baseUrl + "%s/crypto/amount-out", swapCryptoEndpoint)
	
	var swap = new(SwapAmountResponse)
	
	req, err := s.Client.newRequest(http.MethodPost, url, options, s.Client.secretKey)
	if err != nil {
		return nil, err
	}

	if err := do(req, &swap); err != nil {
		return nil, err
	}
	 
	return swap, nil
}
