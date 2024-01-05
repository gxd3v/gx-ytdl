package log

import (
	"bytes"
	"github.com/rs/zerolog"
)

type Log struct {
	Buffer    *bytes.Buffer
	Inner     zerolog.Logger
	SessionID string
}
