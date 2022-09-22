package lazerpay

import (
	"fmt"
	"net/http"
)

type PaymentLinkService service

type PaymentLinkData struct {
	Charge	   		map[string]any 		`json:"charge"`
	Cart 			map[string]any		`json:"cart"`
	Type  			string  		  	`json:"type"`
	IsActive		bool 				`json:"isActive"`
	Network 		string 				`json:"network"`
	Slug 			string 				`json:"slug"`
	Id  			string 				`json:"id"`
	Title 			string 				`json:"title"`
	Description  	string 				`json:"description"`
	BusinessId		string 				`json:"businessId"`
	Options 		map[string]any 		`json:"options"`
	Action 			map[string]any 		`json:"action"`
	Url 			string 				`json:"url"`
	CreatedAt 		string 				`json:"createdAt"`
	UpdatedAt 		string 				`json:"updatedAt"`
}

type LinkOptions struct{
	Title 				string 			`json:"title"`
	Description  		string 			`json:"description"`
	Options  			struct{
		CollectPhone 	bool 			`json:"collect_phone"`
		AllowPromo 		bool 			`json:"allow_promo"`
		CollectAddress  bool 			`json:"collect_address"`
	} 									`json:"options"`

	Cart     			struct{
		PriceId  					string 	`json:"price_id"`
		Quantity   int   `json:"quantity"`
		QuantityAdjustable   bool   `json:"quantity_adjustable"`
	}									`json:"cart"`

	Charge    	   		struct{
		Amount  		int64 			`json:"amount"`
		Currency  		string   		`json:"currency"`
	} 									`json:"charge"`
}

type LinkResponse struct {
	apiStatus
	Data 	PaymentLinkData	`json:"data"`
}

const (
	paymentLinkEndpoint = "api/v1/payment_links"
)

// Create: create a new standrad payment link
// @params {options}: the required options to create a payment link
func (p *PaymentLinkService) Create(options LinkOptions) (*LinkResponse, *http.Response, error ) {
	url := fmt.Sprintf(p.baseUrl + "%s/standard", paymentLinkEndpoint)

	var linkRes = new(LinkResponse)
	resp, err := newRequest(http.MethodPost, url, options, &linkRes)
	if err != nil {
		return nil, nil, err
	}
	 
	return linkRes, resp, nil
}

// Update: updates a standrad payment link
// @params {options}: the required options to be updated
func (p *PaymentLinkService) Update(options LinkOptions) (*LinkResponse, *http.Response, error) {
	url := fmt.Sprintf(p.baseUrl + "%s/standard", paymentLinkEndpoint)
	
	
	var linkres = new(LinkResponse)
	resp, err := newRequest(http.MethodPut, url, options, &linkres)
	if err != nil {
		return nil, nil, err
	}

	return linkres, resp, nil	
}

type ListLinksResponse struct {
	apiStatus

	Links 			[]struct{
		Id 			string 		`json:"id"`
		Title 		string 		`json:"title"`
		Decription  string 		`json:"decription"`
		Type        string  	`json:"type"`
		Network 	string   	`json:"network"`
		Slug 		string 		`json:"slug"`
		IsActive  	bool  		`json:"isActive"`
		WebsiteUrl 	string  	`json:"webSiteUrl"`
		CreatedAt 	string   	`json:"createdAt"`
		UpdatedAt 	string   	`json:"updatedAt"`
	}							`json:"links"`

	Count 			int 		`json:"count"`
	CurrentPage 	int 		`json:"currentPage"`
	NextPage 		int     	`json:"nextPage"`
	PrevPage 		int  		`json:"prevPage"`
	LastPage 		int  		`json:"lastPage"`
}

// GetAll: gets all available payment links
func (p *PaymentLinkService) GetAll() (*ListLinksResponse, *http.Response, error ) {
	url := fmt.Sprintf(p.baseUrl + "%s", paymentLinkEndpoint)

	var linksRes *ListLinksResponse
	resp, err := newRequest(http.MethodGet, url, nil, &linksRes)
	if err != nil {
		return nil, nil, err
	}

	return linksRes, resp, nil
}

// Get: gets a payment link either by id or by slug
func (p *PaymentLinkService) Get(identifier string) (*LinkResponse, *http.Response, error) {
	url := fmt.Sprintf(p.baseUrl + "%s/pay/%s", paymentLinkEndpoint, identifier)

	var linkRes = new(LinkResponse)
	resp, err := newRequest(http.MethodGet, url, nil, &linkRes)
	if err != nil {
		return nil, nil, err
	}

	return linkRes, resp, nil
}


