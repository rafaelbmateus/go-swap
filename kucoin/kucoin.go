package kucoin

import (
	"context"
	"os"

	"github.com/Kucoin/kucoin-go-sdk"
	"github.com/rs/zerolog"
)

type Kucoin struct {
	Log     *zerolog.Logger
	Context *context.Context
	Api     *kucoin.ApiService
}

func NewKucoin(log *zerolog.Logger, ctx *context.Context) *Kucoin {
	return &Kucoin{
		Log:     log,
		Context: ctx,
		Api: kucoin.NewApiService(
			kucoin.ApiBaseURIOption(os.Getenv("KUCOIN_API_BASE_URI")),
			kucoin.ApiKeyOption(os.Getenv("KUCOIN_API_KEY")),
			kucoin.ApiSecretOption(os.Getenv("KUCOIN_API_SECRET")),
			kucoin.ApiPassPhraseOption(os.Getenv("KUCOIN_API_PASSPHRASE")),
			kucoin.ApiKeyVersionOption(kucoin.ApiKeyVersionV2),
		),
	}
}
