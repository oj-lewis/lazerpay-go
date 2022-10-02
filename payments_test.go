package lazerpay

import (
	"testing"

	"github.com/stretchr/testify/require"
	// "github.com/joho/godotenv"
)

func TestPayment_Initailize(t *testing.T) {
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
		require.Nil(t, err)
		require.NotNil(t, resp)
				
	})
}


func TestPayments_Verify(t *testing.T) {

	tests := []struct {
		Name   string
		expected  *verifyPaymentResponse
		data    string
	}{
		{
			Name: "payment verification test",
			expected: new(verifyPaymentResponse),
			data: "1839fdce-5af7-45d1-bac8-157ba29019bb",
		},
		{
			Name: "payment verification test",
			expected: new(verifyPaymentResponse),
			data: "kwaH7iCAsdYp",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			client, _ := NewClient("pk_test_uuZfeK8QvsE69QnLIdujfAIlyl94pOmQ2UFXmk824U1m82VsPu", "sk_test_Ga2alceNHwWriPJWgJpEu2VYZlsLYnEZxjhwZH8N3kYLC5kioF")
			resp, err := client.Payment.Verify(tt.data)
	
			require.Nil(t, err)
			require.NotNil(t, resp)		
		})
	}
	
}