package util

import (
	"errors"
	m "github.com/gx/youtubeDownloader/models"
	"net/url"
	"regexp"
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
