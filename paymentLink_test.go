package lazerpay

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPaymentLink_Create(t *testing.T) {

	tests := []struct {
		Name   string
		expected  *LinkResponse
		data    *LinkOptions
	}{
		{
			Name: "create payment link with charge",
			expected: new(LinkResponse),
			data: &LinkOptions{
				Title: "test payment link",
				Description: "payment link test",
				Options: options{
					CollectPhone: false,
					CollectAddress: false,
					AllowPromo: false,
				},
				Cart: cart{},
				Charge: charge{
					Amount: 3000,
					Currency: "USD",
				},

			},
		},

		{
			Name: "create payment link with cart",
			expected: new(LinkResponse),
			data: &LinkOptions{
				Title: "test payment link",
				Description: "payment link test",
				Options: options{
					CollectPhone: false,
					CollectAddress: false,
					AllowPromo: false,
				},
				Cart: cart{
					Quantity: 10,
					QuantityAdjustable: true,
				},
				Charge: charge{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			client, _ := NewClient(os.Getenv("TEST_PUBLIC_KEY"), os.Getenv("TEST_SECRET_KEY"))
			resp, err := client.PaymentLink.Create(tt.data)

			require.Nil(t, err)
			require.Equal(t, tt.expected, resp)
			
		})
	}
	
}

func TestPaymentLink_GetLinks(t *testing.T) {

	tests := struct {
		Name   string
		expected  []*ListLinksResponse
		data    *PaymentLinkData
	}{
		
		Name: "Get all payment links",
		expected: make([]*ListLinksResponse, 4),
	}

	t.Run(tests.Name, func(t *testing.T) {
		client, _ := NewClient(os.Getenv("TEST_PUBLIC_KEY"), os.Getenv("TEST_SECRET_KEY"))
		resp, err := client.PaymentLink.GetAll()

		require.Nil(t, err)
		require.Equal(t, tests.expected, resp)
	})
	
	
}

func TestPaymentLink_Get(t *testing.T) {
	t.Run("Get a payment link by id ", func(t *testing.T) {
		client, _ := NewClient(os.Getenv("TEST_PUBLIC_KEY"), os.Getenv("TEST_SECRET_KEY"))
		resp, err := client.PaymentLink.Get("id")

		require.Nil(t, err)
		require.NotNil(t, resp)
	})
}


func TestPaymentLink_Update(t *testing.T) {
	data := &LinkOptions{
		Title: "updated title",
		Description: "updated description",
	}

	t.Run("Get a payment link by id ", func(t *testing.T) {
		client, _ := NewClient(os.Getenv("TEST_PUBLIC_KEY"), os.Getenv("TEST_SECRET_KEY"))
		resp, err := client.PaymentLink.Update(data)

		require.Nil(t, err)
		require.NotNil(t, resp)
	})
}