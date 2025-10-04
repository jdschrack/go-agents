package main

import (
	"context"

	"github.com/jdschrack/go-agents/internal/config"
	"github.com/jdschrack/go-agents/internal/log"
)

type setup struct {
	config *config.Config
}

var appSetup setup

func main() {
	ctx := initConfig(context.Background())

	logger := log.FromContext(ctx)
	logger.Info().Any("config", appSetup.config).Msg("Starting agent...")
}

func initConfig(parentCtx context.Context) context.Context {
	// configure suggared logging using zap
	appConfig := config.LoadConfig(parentCtx)
	appSetup = setup{config: appConfig}

	logger, err := log.NewLogger("go-agents", "0.0.1")
	if err != nil {
		panic("could not set up logger")
	}

	return log.WithLogger(parentCtx, logger)
}
