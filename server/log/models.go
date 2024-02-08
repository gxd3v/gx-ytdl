package log

import (
	"bytes"
	"github.com/gx/youtubeDownloader/config"
	"github.com/rs/zerolog"
)

type Log struct {
	Config    *config.Config
	Buffer    *bytes.Buffer
	Inner     zerolog.Logger
	SessionID string
}
