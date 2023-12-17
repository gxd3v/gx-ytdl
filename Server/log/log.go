package log

import (
	"bytes"
	"fmt"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

type Log struct {
	Buffer *bytes.Buffer
	Inner  zerolog.Logger
}

func (l *Log) callerLocation(file string, line int) string {
	fileName := filepath.Base(file)
	parentFolder := filepath.Base(filepath.Dir(file))
	location := fmt.Sprintf("%s/%s:%d", parentFolder, fileName, line)

	return location
}

func (l *Log) Setup() {
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})
	var buffer bytes.Buffer
	writer := zerolog.ConsoleWriter{Out: &buffer, TimeFormat: time.RFC3339}
	l.Inner = log.Output(writer)
	l.Buffer = &buffer
}

func (l *Log) Info(message string, args ...interface{}) {
	defer l.StoreLog()
	_, file, line, _ := runtime.Caller(1)
	location := l.callerLocation(file, line)
	log.Info().Interface("arguments", args).Str("caller", location).Msg(message)
	l.Inner.Info().Interface("arguments", args).Str("caller", location).Msg(message)
}

func (l *Log) Error(message string, args ...interface{}) {
	defer l.StoreLog()
	_, file, line, _ := runtime.Caller(1)
	location := l.callerLocation(file, line)
	log.Error().Interface("arguments", args).Str("caller", location).Msg(message)
	l.Inner.Error().Interface("arguments", args).Str("caller", location).Msg(message)
}

func (l *Log) Warning(message string, args ...interface{}) {
	defer l.StoreLog()
	_, file, line, _ := runtime.Caller(1)
	location := l.callerLocation(file, line)
	log.Warn().Interface("arguments", args).Str("caller", location).Msg(message)
	l.Inner.Warn().Interface("arguments", args).Str("caller", location).Msg(message)
}

func (l *Log) Debug(message string, args ...interface{}) {
	defer l.StoreLog()
	_, file, line, _ := runtime.Caller(1)
	location := l.callerLocation(file, line)
	log.Debug().Interface("arguments", args).Str("caller", location).Msg(message)
	l.Inner.Debug().Interface("arguments", args).Str("caller", location).Msg(message)
}

func (l *Log) StoreLog() {
	l.Buffer.Reset()
}
