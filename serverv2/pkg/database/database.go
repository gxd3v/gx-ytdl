package database

import (
	"github.com/gobuffalo/pop/v6"
	"github.com/gx/gx-ytdl/serverv2/config"
	"github.com/gx/gx-ytdl/serverv2/pkg/logger"
	"os"
	"path"
)

func Connect(cfg *config.Database) (*pop.Connection, error) {
	conn, err := pop.Connect(cfg.Conn)
	if err != nil {
		logger.Err(err).Msg("couldn't connect to database")
		return nil, err
	}

	cwd, _ := os.Getwd()

	migrator, err := pop.NewFileMigrator(path.Join(cwd, "pkg", "database", "migrations"), conn)
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
