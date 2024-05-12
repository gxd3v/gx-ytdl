package downloader

import (
	"context"
	"github.com/google/uuid"
	"github.com/gx/youtubeDownloader/constants"
	"github.com/gx/youtubeDownloader/internal/downloader/helpers"
	"github.com/gx/youtubeDownloader/internal/logger"
	"os"
	"path"
	"strings"
	"time"
)

func CreateOrGetSession(ctx context.Context, sessionId string) (string, error) {
	log := logger.Fields(map[string]any{
		"where":   "downloader",
		"session": sessionId,
	})

	if len(strings.TrimSpace(sessionId)) < 0 {
		sessionId = uuid.NewString()
	}

	fullPath := path.Join(constants.Temp, sessionId)

	_, err := os.Stat(fullPath)
	if err != nil && os.IsExist(err) {
		log.Info().Msg("session folder is already created")
		return fullPath, nil
	}

	if err != nil && !os.IsNotExist(err) {
		log.Err(err).Msg("failed to check for session folder")
		return "", err
	}

	if err := os.Mkdir(fullPath, os.ModeAppend); err != nil {
		logger.Err(err).Msg("failed to create session folder")
		return "", err
	}

	t := time.Now()

	data := &Session{
		Id:        uuid.NewString(),
		CreatedAt: t,
		UpdatedAt: t,
		DeletedAt: nil,
		CreatedBy: "admin",
		Session:   sessionId,
		LastLogin: t,
	}

	//TODO save session on database

	_ = data

	ctx = context.WithValue(ctx, "session", sessionId)

	return fullPath, nil
}

func DeleteSession(ctx context.Context) (bool, error) {
	sessionId, err := helpers.GetContextSession(ctx)
	if err != nil {
		logger.Err(err).Msg("failed to delete session")
		return false, err
	}

	t := time.Now()

	//TODO get sesion from DB
	session := &Session{}

	session.DeletedAt = &t
	session.UpdatedAt = t

	//TODO update session value on DB

	fullPath := path.Join(constants.Temp, sessionId)

	exists, err := helpers.CheckPathExists(fullPath)
	if err != nil {
		logger.Err(err).Msg("failed to delete session")
		return false, err
	}

	if exists {
		if err := os.Remove(fullPath); err != nil {
			return false, err
		}
	}

	return true, nil
}
