package bot_test

import (
	"context"
	"os"
	"testing"

	"github.com/rafaelbmateus/go-swap/bot"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

var (
	log = zerolog.New(os.Stdout)
	ctx = context.Background()
)

func TestRun(t *testing.T) {
	bot := bot.NewBot(&log, &ctx, nil, nil, nil)
	err := bot.Run()
	assert.NoError(t, err)
}
