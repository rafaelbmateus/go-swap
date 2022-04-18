package main

import (
	"context"
	"os"

	"github.com/rafaelbmateus/go-swap/bot"
	"github.com/rs/zerolog"
)

func main() {
	ctx := context.Background()
	log := zerolog.New(os.Stdout)

	bot := bot.NewBot(&log, &ctx, nil, nil, nil)
	err := bot.Run()
	if err != nil {
		log.Err(err).Msgf("error on run %q", err)
	}
}
