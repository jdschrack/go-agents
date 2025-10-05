package main

import (
	"context"
	"database/sql"

	"github.com/jdschrack/go-agents/internal/config"
	"github.com/jdschrack/go-agents/internal/data"
	"github.com/jdschrack/go-agents/internal/log"
)

type setup struct {
	config *config.Config
	db     *sql.DB
}

var appSetup setup

func main() {
	ctx := initConfig(context.Background())

	logger := log.FromContext(ctx)
	logger.Info().Any("config", appSetup.config).Msg("starting agent...")

	defer func() {
		if err := appSetup.db.Close(); err != nil {
			logger.Error().Err(err).Msg("failed to close database")
		}
	}()
}

func initConfig(parentCtx context.Context) context.Context {
	appConfig := config.LoadConfig(parentCtx)
	appSetup = setup{config: appConfig}

	logger, err := log.NewLogger("go-agents", "0.0.1")
	if err != nil {
		panic("could not set up logger")
	}

	ctx := log.WithLogger(parentCtx, logger)

	db, err := data.GetConnection(ctx, appConfig.DatabasePath, false)
	if err != nil {
		logger.Err(err).Msg("could not connect to database")
	}

	return data.WithDatabase(ctx, db)
}
