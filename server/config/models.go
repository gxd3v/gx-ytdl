package config

type Config struct {
	Database     Database     `json:"database"`
	Transporters Transporters `json:"transporters"`
}

type Database struct {
	Conn string `json:"connection"`
	Type string `json:"type"`
	Soda string `json:"soda"`
	Pool struct {
		MaxConnections  int `json:"maxConnections"`
		IdleConnections int `json:"idleConnections"`
		TTL             int `json:"TTL"`
	} `json:"pool"`
}

type Transporters struct {
	HTTP HTTP      `json:"HTTP"`
	RPC  RPC       `json:"RPC"`
	WS   Websocket `json:"WS"`
}

type HTTP struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type RPC struct {
	Host    string `json:"host"`
	Port    string `json:"port"`
	Reflect bool   `json:"reflect"`
}

type Websocket struct {
	Host    string `json:"host"`
	Port    string `json:"port"`
	MaxSize int    `json:"maxSize"`
}
