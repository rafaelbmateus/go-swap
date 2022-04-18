package kucoin_test

import (
	"testing"

	"github.com/rafaelbmateus/go-swap/kucoin"
	"github.com/stretchr/testify/assert"
)

func TestNewKucoin(t *testing.T) {
	assert.NotNil(t, kucoin.NewKucoin(nil, nil))
}
