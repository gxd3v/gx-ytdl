package helpers

import (
	"context"
	"errors"
	"github.com/gx/youtubeDownloader/internal/logger"
	"os"
	"strings"
)

func CheckPathExists(path string) (bool, error) {
	_, err := os.Stat(path)

	if err != nil && os.IsExist(err) {
		return true, nil
	}

	if err != nil && !os.IsNotExist(err) {
		logger.Err(err).Msgf("%s failed to check if target path exists", path)
		return false, err
	}

	return false, err
}

func GetContextSession(ctx context.Context) (string, error) {
	sessionId, ok := ctx.Value("session").(string)
	if !ok {
		logger.Warn().Msg("session cast failed")
		return "", errors.New("no session")
	}

	if len(strings.TrimSpace(sessionId)) <= 0 {
		err := errors.New("no session")
		logger.Err(err).Msg("session id was not in context")
		return "", err
	}

	return sessionId, nil
}
