package lazerpay

import (
	// "fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	// "github.com/joho/godotenv"
)

func TestInitailize(t *testing.T) {
	t.Run("initialize", func(t *testing.T) {

		client, _ := NewClient("pk_test_uuZfeK8QvsE69QnLIdujfAIlyl94pOmQ2UFXmk824U1m82VsPu", "sk_test_Ga2alceNHwWriPJWgJpEu2VYZlsLYnEZxjhwZH8N3kYLC5kioFsk_test_Ga2alceNHwWriPJWgJpEu2VYZlsLYnEZxjhwZH8N3kYLC5kioF")
		data := &PaymentOption{
			CustomerName: "omiete",
			CustomerEmail: "omiete@example.com",
			Coin: "BUSD",
			Currency: "USD",
			FiatAmount: 5,
		}

		resp, err := client.Payment.Initialize(data)
		assert.Nil(t, err)
		assert.NotNil(t, resp)
				
	})
}


// func TestVerify(t *testing.T) {

// 	tests := []struct {
// 		Name   string
// 		expected  *verifyPaymentResponse
// 		data    string
// 	}{
// 		{
// 			Name: "payment verification test",
// 			expected: new(verifyPaymentResponse),
// 			data: "1839fdce-5af7-45d1-bac8-157ba29019bb",
// 		},
// 		{
// 			Name: "payment verification test",
// 			expected: new(verifyPaymentResponse),
// 			data: "kwaH7iCAsdYp",
// 		},
// 	}

// 	for _, tt := range tests {
// 		t.Run(tt.Name, func(t *testing.T) {
// 			client, _ := NewClient("pk_test_uuZfeK8QvsE69QnLIdujfAIlyl94pOmQ2UFXmk824U1m82VsPu", "sk_test_Ga2alceNHwWriPJWgJpEu2VYZlsLYnEZxjhwZH8N3kYLC5kioF")
// 			resp, err := client.Payment.Verify(tt.data)
	
// 			assert.Nil(t, err)
// 			assert.NotNil(t, resp)
// 			// assert.Equal(t, tt.expected, resp)			
// 		})
// 	}
	
// }