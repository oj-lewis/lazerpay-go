package lazerpay

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)


func TestPayout_Crypto(t *testing.T) {
	data := &CryptoPayoutOptions{
		Amount: 100,
		Recipient: "",
		Coin: "USDT",
		Blockchain: "binance smart chain",
	}

	t.Run("test for crypto payout", func(t *testing.T) {
		client, _ := NewClient(os.Getenv("TEST_PUBLIC_KEY"), os.Getenv("TEST_SECRET_KEY"))
		resp, err := client.Payout.Crypto(data)

		require.Nil(t, err)
		require.NotNil(t, resp)
	})
}

func TestPayout_Create(t *testing.T) {
	data := &BankOptions{
		BankName: "",
		BankCode: "",
		AccountName: "",
		AccountNumber: "",
		Currency: "",
		Country: "",
		Default: true,
	}

	t.Run("test for crypto payout", func(t *testing.T) {
		client, _ := NewClient(os.Getenv("TEST_PUBLIC_KEY"), os.Getenv("TEST_SECRET_KEY"))

		resp, err := client.Payout.Create(data)

		require.Nil(t, err)
		require.NotNil(t, resp)
	})
}

func TestPayout_Initiate(t *testing.T) {
	data := &PayoutOptions{
		BankId: "",
		Coin: "",
		Amount: 12,
	}

	t.Run("test for crypto payout", func(t *testing.T) {
		client, _ := NewClient(os.Getenv("TEST_PUBLIC_KEY"), os.Getenv("TEST_SECRET_KEY"))

		resp, err := client.Payout.Initiate(data)

		require.Nil(t, err)
		require.NotNil(t, resp)
	})
}

func TestPayout_Update(t *testing.T) {
	data := &BankOptions{
		BankName: "",
		BankCode: "",
		AccountName: "",
		AccountNumber: "",
		Currency: "",
		Country: "",
		Default: false,
	}

	t.Run("test for crypto payout", func(t *testing.T) {
		client, _ := NewClient(os.Getenv("TEST_PUBLIC_KEY"), os.Getenv("TEST_SECRET_KEY"))

		resp, err := client.Payout.Update("", data)

		require.Nil(t, err)
		require.NotNil(t, resp)
	})
}

func TestPayout_Delete(t *testing.T) {
	data := []string{}

	t.Run("test for crypto payout", func(t *testing.T) {
		client, _ := NewClient(os.Getenv("TEST_PUBLIC_KEY"), os.Getenv("TEST_SECRET_KEY"))

		resp, err := client.Payout.Delete(data)

		require.Nil(t, err)
		require.NotNil(t, resp)
	})
}
func TestPayout_GetAall(t *testing.T) {
	
	t.Run("test for crypto payout", func(t *testing.T) {
		client, _ := NewClient(os.Getenv("TEST_PUBLIC_KEY"), os.Getenv("TEST_SECRET_KEY"))

		resp, err := client.Payout.GetAll()

		require.Nil(t, err)
		require.NotNil(t, resp)
	})
}