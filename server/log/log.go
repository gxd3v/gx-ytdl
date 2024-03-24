package log

import (
	"github.com/pieterclaerhout/go-log"
	"os"
)

func Info(msg string, args ...string) {
	log.DebugMode = true
	log.PrintTimestamp = true
	log.PrintColors = true
	log.TimeFormat = "2006-01-02 15:04:05.000"

	if len(args) > 0 {
		log.Infof(msg, args)
	} else {
		log.Infof(msg)
	}
}

func Warn(msg string, args ...string) {
	log.DebugMode = true
	log.PrintTimestamp = true
	log.PrintColors = true
	log.TimeFormat = "2006-01-02 15:04:05.000"

	if len(args) > 0 {
		log.Warnf(msg, args)
	} else {
		log.Warnf(msg)
	}
}

func Error(err error, msg string, args ...string) {
	log.DebugMode = true
	log.PrintTimestamp = true
	log.PrintColors = true
	log.TimeFormat = "2006-01-02 15:04:05.000"

	if len(args) > 0 {
		log.Errorf(msg, args)
	} else {
		log.Error(msg)
	}
	log.StackTrace(err)
}

func Fatal(err error, msg string, args ...string) {
	log.DebugMode = true
	log.PrintTimestamp = true
	log.PrintColors = true
	log.TimeFormat = "2006-01-02 15:04:05.000"

	if len(args) > 0 {
		log.Errorf(msg, args)
	} else {
		log.Error(msg)
	}

	log.StackTrace(err)
	os.Exit(1)
}
