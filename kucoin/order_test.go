package kucoin_test

import (
	"testing"

	sdk "github.com/Kucoin/kucoin-go-sdk"
	"github.com/google/uuid"
	"github.com/rafaelbmateus/go-swap/kucoin"
)

func TestNewOrder(t *testing.T) {
	tests := []struct {
		name   string
		symbol string
		typ    string
		side   string
		size   string
		price  string
	}{
		{
			name:   "basic test",
			symbol: "ETH-BTC",
			side:   "buy",
			size:   "10",
			price:  "0.1",
		},
	}
	k := kucoin.NewKucoin(nil, nil)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k.NewOrder(&sdk.CreateOrderModel{
				ClientOid: uuid.New().String(),
				Symbol:    tt.symbol,
				Side:      tt.side,
				Size:      tt.size,
				Price:     tt.price,
			})
		})
	}
}
