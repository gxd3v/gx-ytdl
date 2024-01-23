package models

import "fmt"

type JSONBodyMessage struct {
	Code    string `json:"status"`
	Message string `json:"message"`
}

type WebsocketMessage struct {
	Code    string                 `json:"code"`
	Payload map[string]interface{} `json:"payload"`
}

type WebsocketServerResponse struct {
	Id      string          `json:"id"`
	Success bool            `json:"success"`
	Data    JSONBodyMessage `json:"data"`
}

func (wm WebsocketMessage) ToString() string {
	return fmt.Sprintf("%+v", wm)
}
