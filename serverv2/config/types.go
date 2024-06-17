package config

type API struct {
	Host string `json:"host"`
	Port string `json:"port"`
}

type Database struct {
	Conn string       `json:"connection"`
	Type string       `json:"type"`
	Soda string       `json:"soda"`
	Pool DatabasePool `json:"pool"`
}

type DatabasePool struct {
	MaxConnections  int `json:"maxConnections"`
	IdleConnections int `json:"idleConnections"`
	TTL             int `json:"TTL"`
}
