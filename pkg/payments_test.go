package lazerpay

import (
	"testing"
)

func TestInitailize(t *testing.T) {
	t.Run("initialize", func(t *testing.T) {

			var client, _ = NewClient("test")
			data := PaymentOption{
				CustomerName: "omiete",
				CustomerEmail: "omiete@example.com",
				Coin: "BUSD",
				Currency: "USD",
			}

			resp, _, err := client.Payment.Initialize(data)
			t.Errorf("want: %v, got: %v", resp, err)
			
	})
}


func TestVerify(t *testing.T) {

	tests := []struct {
		Name   string
		expected  *verifyPaymentResponse
		data    string
	}{
		{
			Name: "payment verification test",
			expected: new(verifyPaymentResponse),
			data: "eiunjigugugeiueneu",
		},
		{
			Name: "payment verification test",
			expected: new(verifyPaymentResponse),
			data: "eiunjigugugeiuenekhu",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			client, _ := NewClient("test")
			resp, _, _ := client.Payment.Verify(tt.data)
	
			t.Errorf("want: %v, got: %v", tt.expected, resp)
		})
	}
	
}