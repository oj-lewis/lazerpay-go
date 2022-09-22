package lazerpay

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
)

const (
	baseUrl string = "https://api.lazerpay.engineering/"
)

type Client struct {
	baseUrl      string
	secretKey    string

	Payment      *PaymentService
	PaymentLink  *PaymentLinkService
	Transfers    *PayoutService
	Swap         *Swapservice
}

type apiStatus struct {
	Status 		string		`json:"status"`
	StatusCode 	int			`json:"statusCode"`
	Message 	string		`json:"message"`
}

type service struct{
	*Client
}

//NewClient: Creates a new instance of lazerpay client
//@params {secretKey}: A unique secret token for interacting with the lazerpay api
func NewClient(secretKey string) (*Client, error) {
	c := new(Client)

	c.secretKey = secretKey
	c.serviceSetup()

	return c, c.validate()
}

// ServiceSetup sets up services provided by the API
func (c *Client) serviceSetup() {
	c.Payment       = &PaymentService{c}
	c.PaymentLink   = &PaymentLinkService{c}
	c.Transfers     = &PayoutService{c}
	c.Swap          = &Swapservice{c}
	c.setBaseUrl()
}


// Validate checks if an API key is provided,
// returns an error if non is provided
func (c *Client) validate() error {
	if len(strings.Trim(c.secretKey, " ")) == 0 {
		return errors.New("please provide your secret key")
	}

	return nil
}

// setBaseUrl sets the base url for the api
func (c *Client) setBaseUrl() {
	c.baseUrl = baseUrl
}


// NewRequest makes a new http request for an operation,
// @params {method} : The http request method,
// @params {_url} : The url of the request,
// @params {payload} : The body of the request,
// @params {v} : The body of the response is pointed to v
func newRequest(method string, _url string, payload, v interface{}) (*http.Response, error) {
	c := &Client{}
	if strings.HasPrefix(_url, "/") {
		return nil, errors.New("url should not start wth /")
	}

	var _body = bytes.NewBuffer(nil)
	if payload != nil {
		if err := json.NewEncoder(_body).Encode(payload); err != nil {
			return nil, err
		}
	}
	req, err := http.NewRequest(method, _url, _body)
	if err != nil {
		return nil, err
	}

	if payload != nil {
		req.Header.Add("Content-Type", "application/json")
	}
	req.Header.Add("Authorization", "Bearer" + c.secretKey)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	
	defer resp.Body.Close()
	
	if n := resp.StatusCode; n != http.StatusOK && n != http.StatusCreated {
		var d struct {
			Message string `json:"maessage"`
		}

		if err := json.NewDecoder(resp.Body).Decode(&d); err != nil {
			return nil, err
		}

		return nil, errors.New(d.Message)
	}

	if v != nil {
		if err := json.NewDecoder(resp.Body).Decode(v); err != nil {
			return nil, err
		}

	}

	return resp, nil
}

// RandomString generates a random string with a specific size s
func RandomString(s int) (string) {
	letters := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

	bytes := make([]byte, s)
	if _, err := rand.Read(bytes); err != nil {
		return string(err.Error())
	}
	for i, b := range bytes {
		bytes[i] = letters[b%byte(len(letters))]
	}
	return  string(bytes)	
}