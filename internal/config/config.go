// config/config.go
package config

type tAppConfig struct {
	Port     int
	LogLevel string
}

var AppConfig *tAppConfig

func init() {
	AppConfig = &tAppConfig{
		Port:     8080,
		LogLevel: "info",
	}
}
