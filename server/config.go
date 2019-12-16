package server

// default configurations
const (
	DefaultPort = 8080
	DefaultHost = "0.0.0.0"
)

// Config holds the settings for the server
type Config struct {
	Port int
	Host string
}

func applyDefaults(config *Config) {
	if config.Port == 0 {
		config.Port = DefaultPort
	}
	if config.Host == "" {
		config.Host = DefaultHost
	}
}
