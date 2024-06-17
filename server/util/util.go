package util

import (
	"errors"
	"github.com/gx/youtubeDownloader/internal/core/logger"
	m "github.com/gx/youtubeDownloader/internal/models"
	"net"
	"net/url"
	"regexp"
	"time"
)

func ResponseJSONBody(status, message string) m.JSONBodyMessage {
	return m.JSONBodyMessage{
		Code:    status,
		Message: message,
	}

}

func ParseURL(u string) (*url.URL, error) {
	regex, err := regexp.Compile("https?://(www\\.)?[-a-zA-Z0-9@:%._+~#=]{1,256}\\.[a-zA-Z0-9()]{1,6}\\b([-a-zA-Z0-9()!@:%_+.~#?&/=]*)")
	if err != nil {
		return nil, err
	}

	if regex.Match([]byte(u)) {
		parsed, _ := url.Parse(u)
		return parsed, nil
	}

	return nil, errors.New("url is not valid")
}

func CheckPortBusy(host, port string) bool {
	if conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), 3*time.Second); err != nil {
		logger.Warn().Msgf("port %s is busy for host %s", port, host)
		return true
	} else {
		_ = conn.Close()
		return false
	}
}
