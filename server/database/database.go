package database

import (
	"github.com/gobuffalo/pop/v6"
	"github.com/gx/youtubeDownloader/config"
	"github.com/gx/youtubeDownloader/internal/core/logger"
)

func Connect(cfg config.Database) (*pop.Connection, error) {
	conn, err := pop.Connect(cfg.Conn)
	if err != nil {
		logger.Err(err).Msg("couldn't connect to database")
		return nil, err
	}

	migrator, err := pop.NewFileMigrator("./migrations", conn)
	if err != nil {
		logger.Err(err).Msg("failed to read migration files")
		return nil, err
	}

	if err := migrator.Up(); err != nil {
		logger.Err(err).Msg("failed to run migrations")
		return nil, err
	}

	return conn, nil
}
