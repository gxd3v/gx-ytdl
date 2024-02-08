package config

import (
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/gx/youtubeDownloader/constants"
	"os"
)

func (config *Config) Get() (*Config, error) {
	if value, ok := os.LookupEnv(config.envString("SERVICE-ID")); !ok {
		return nil, errors.New("no service id found in env variables")
	} else {
		config.ServiceID = value
	}

	if value, ok := os.LookupEnv(config.envString("CONNECTION-ROUTE")); !ok {
		return nil, errors.New("no connection route found in env variables")
	} else {
		config.ConnectionRoute = value
	}

	if value, ok := os.LookupEnv(config.envString("DATABASE")); !ok {
		return nil, errors.New("no database address found in env variables")
	} else {
		decodeString, err := base64.StdEncoding.DecodeString(value)
		if err != nil {
			return nil, err
		}
		config.Database = string(decodeString)
	}

	if value, ok := os.LookupEnv(config.envString("PYTHON-BINARY")); !ok {
		return nil, errors.New("no python binary found in env variables")
	} else {
		config.PythonBinary = value
	}

	if value, ok := os.LookupEnv(config.envString("DOWNLOADER-PATH")); !ok {
		return nil, errors.New("no downloader path found in env variables")
	} else {
		config.DownloaderPath = value
	}

	if value, ok := os.LookupEnv(config.envString("OUTPUT-PATH")); !ok {
		return nil, errors.New("no output path found in env variables")
	} else {
		config.OutputPath = value
	}

	return config, nil
}

func (config *Config) envString(val string) string {
	config.ServicePrefix = constants.ServicePrefix
	return fmt.Sprintf("GX-%s-%s", config.ServicePrefix, val)
}
