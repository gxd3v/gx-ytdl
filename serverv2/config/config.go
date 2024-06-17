package config

import (
	"encoding/json"
	"github.com/gx/gx-ytdl/serverv2/pkg/logger"
	"os"
)

func GetAPI(file string) (*API, error) {
	cfg := new(API)

	data, err := os.ReadFile(file)
	if err != nil {
		logger.Fatal().Err(err).Msgf("failed to read %s", file)
		return nil, err
	}

	if err = json.Unmarshal(data, &cfg); err != nil {
		logger.Fatal().Err(err).Msg("failed to parse file into config model")
		return nil, err
	}

	return cfg, nil
}

func GetDatabase(file string) (*Database, error) {
	cfg := new(Database)

	data, err := os.ReadFile(file)
	if err != nil {
		logger.Fatal().Err(err).Msgf("failed to read %s", file)
		return nil, err
	}

	if err = json.Unmarshal(data, &cfg); err != nil {
		logger.Fatal().Err(err).Msg("failed to parse file into config model")
		return nil, err
	}

	return cfg, nil
}
