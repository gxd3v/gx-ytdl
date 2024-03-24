package config

type Config struct {
	ServicePrefix   string `json:"servicePrefix,omitempty"`
	ServiceID       string `json:"serviceID,omitempty"`
	ConnectionRoute string `json:"routeConnect,omitempty"`

	DatabaseEnv string `json:"database,omitempty"`

	PythonBinary   string `json:"pythonBinary,omitempty"`
	DownloaderPath string `json:"downloaderPath,omitempty"`
	OutputPath     string `json:"outputPath,omitempty"`

	SodaPath string `json:"sodaPath,omitempty"`
}
