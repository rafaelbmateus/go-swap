package binance

import (
	"github.com/adshao/go-binance/v2"
)

// CreateOrder create new order in binance.
func (me *Binance) CreateOrder(o *binance.CreateOrderService,
) (*binance.CreateOrderResponse, error) {
	order, err := o.Do(*me.Context)
	if err != nil {
		return nil, err
	}

	return order, nil
}
