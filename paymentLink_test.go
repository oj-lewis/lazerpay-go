package lazerpay

// import (
// 	"testing"
// )

// func TestPaymentLinkCreate(t *testing.T) {

// 	tests := []struct {
// 		Name   string
// 		expected  *LinkResponse
// 		data    LinkOptions
// 	}{
// 		{
// 			Name: "create payment link",
// 			expected: new(LinkResponse),
// 			data: LinkOptions{

// 			},
// 		},
// 		// {
// 		// 	Name: "payment verification test",
// 		// 	expected: new(PaymentLinkResponse),
// 		// 	data: &PaymentLinkData{

// 		// 	},
// 		// },
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.Name, func(t *testing.T) {
// 			client, _ := NewClient("test", "test")
// 			resp, _, _ := client.PaymentLink.Create(tt.data)
// 			if resp != tt.expected {
// 				t.Errorf("want: %v, got: %v", tt.expected, resp)
// 			}
// 		})
// 	}
	
// }

// func TestPaymentLinkGetLinks(t *testing.T) {

// 	tests := []struct {
// 		Name   string
// 		expected  []*ListLinksResponse
// 		data    *PaymentLinkData
// 	}{
// 		{
// 			Name: "list all payment links",
// 			expected: make([]*ListLinksResponse, 4),
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.Name, func(t *testing.T) {
// 			client, _ := NewClient("test", "test")
// 			resp, _, _ := client.PaymentLink.GetAll()
// 			if resp != nil {
// 				t.Errorf("want: %v, got: %v", tt.expected, resp)
// 			}
// 		})
// 	}
	
// }