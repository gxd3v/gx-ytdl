package util

import m "github.com/gx/youtubeDownloader/models"

func ResponseJSONBody(status, message string) m.JSONBodyMessage {
	return m.JSONBodyMessage{
		Code:    status,
		Message: message,
	}

}
