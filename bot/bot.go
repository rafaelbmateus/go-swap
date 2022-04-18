package bot

import (
	"context"

	"github.com/rafaelbmateus/go-swap/binance"
	"github.com/rafaelbmateus/go-swap/kucoin"
	"github.com/rafaelbmateus/go-swap/notify"
	"github.com/rs/zerolog"
)

// Bot is a representation of this package.
type Bot struct {
	Log     *zerolog.Logger
	Context *context.Context
	Binance *binance.Binance
	Kucoin  *kucoin.Kucoin
	Notify  *notify.SlackNotify
}

// NewBot create a new bot instance.
func NewBot(log *zerolog.Logger, ctx *context.Context, binance *binance.Binance,
	kucoin *kucoin.Kucoin, notify *notify.SlackNotify) *Bot {
	return &Bot{
		Log:     log,
		Context: ctx,
		Binance: binance,
		Kucoin:  kucoin,
		Notify:  notify,
	}
}

// Run bot.
func (me *Bot) Run() error {
	me.Log.Info().Msg("bot is running")
	return nil
}
