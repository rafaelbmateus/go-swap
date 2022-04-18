package kucoin

import "github.com/Kucoin/kucoin-go-sdk"

// NewOrder create a new order in kucoin.
// https://docs.kucoin.com/#place-a-new-order
func (me *Kucoin) NewOrder(o *kucoin.CreateOrderModel,
) (*kucoin.ApiResponse, error) {
	res, err := me.Api.CreateOrder(o)
	if err != nil {
		return nil, err
	}

	return res, nil
}
