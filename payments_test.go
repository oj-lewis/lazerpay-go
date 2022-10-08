package lazerpay

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/joho/godotenv"
)

var _ = godotenv.Load()

func TestPayment_Initailize(t *testing.T) {
	t.Run("initialize", func(t *testing.T) {

		client, _ := NewClient(os.Getenv("TEST_PUBLIC_KEY"), os.Getenv("TEST_SECRET_KEY"))
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
			client, _ := NewClient(os.Getenv("TEST_PUBLIC_KEY"), os.Getenv("TEST_SECRET_KEY"))
			resp, err := client.Payment.Verify(tt.data)
	
			require.Nil(t, err)
			require.NotNil(t, resp)		
		})
	}
	
}