package config

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gx/youtubeDownloader/internal/logger"
	"os"
)

const configFile = "config.json"

func New(_ context.Context) *Config {
	cfg := new(Config)

	data, err := os.ReadFile(configFile)
	if err != nil {
		logger.Fatal().Err(err).Msgf("failed to read %s", configFile)
		os.Exit(-1)
	}

	err = json.Unmarshal(data, &cfg)
	if err != nil {
		logger.Fatal().Err(err).Msg("failed to parse file into config structure")
		os.Exit(-1)
	}

	if value, ok := os.LookupEnv("DB_PWD"); !ok {
		logger.Fatal().Err(errors.New("no env variable found for database connection")).Msg("database connection not found")
		os.Exit(-1)
	} else {
		conn, err := base64.StdEncoding.DecodeString(value)
		if err != nil {
			logger.Fatal().Err(err).Msg("database connection string invalid")
			os.Exit(-1)
		}

		cfg.Database.Conn = fmt.Sprintf(cfg.Database.Conn, string(conn))
	}

	return cfg
}
