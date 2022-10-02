package lazerpay

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSwap_SwapCrypto(t *testing.T) {
	data := &SwapOptions{
		Amount: 100,
		FromCoin: "USDT",
		ToCoin: "BUSD",
		Blockchain: "binance smart chain",
	}

	t.Run("swap crypto", func(t *testing.T) {
		client, _ := NewClient("test", "test")

		resp, err := client.Swap.SwapCrypto(data)

		require.Nil(t, err)
		require.NotNil(t, resp)
	})
}

func TestSwap_GetSwapAmount(t *testing.T) {
	data := &SwapOptions{
		Amount: 100,
		FromCoin: "USDT",
		ToCoin: "BUSD",
		Blockchain: "binance smart chain",
	}

	t.Run("swap crypto", func(t *testing.T) {
		client, _ := NewClient("test", "test")

		resp, err := client.Swap.SwapCrypto(data)

		require.Nil(t, err)
		require.NotNil(t, resp)
	})
}