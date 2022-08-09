package lazerpay



type Payments struct {
	LazerPay
}

type initializePayment interface {
	initialize(payload PaymentData) (map[string]interface{}, error)
}

type VerifyPayment interface {
	verify(reference string) (map[string]interface{}, error)
}

type PaymentData struct {
	Reference string    `json:"reference"`
	CustomerName string  `json:"customer_name"`
	CustomerEmail string  `json:"customer_email"`
	Coin string            `json:"coin"`
	Currency  string       `json:"currency"`
	Amount     float32     `json:"amount"`
	AcceptPartialPayment bool `json:"accept_partial_payment"`
}

func ( p Payments) initialize(payload PaymentData) (map[string]interface{}, error) {
	data := ToJSON(payload)
	url := p.GetBaseurl() + p.Endpoint("payment", "initialize")

	resp, err := PostRequest(data, url, p.SecretKey)
	if err != nil {
		return  nil, err
	}

	return resp, nil
}

func (p Payments) verify(reference string) (map[string]interface{}, error) {
	url := p.GetBaseurl() + p.Endpoint("payment", "verify") + reference

	resp, err := GetRequest(url, p.SecretKey)
	if err != nil {
		return nil, err
	}

	return resp, nil
	
}

