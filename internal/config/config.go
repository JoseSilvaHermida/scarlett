// config/config.go
package config

type AppConfigType struct {
	Port       int
	SocketPath string
}

var AppConfig *AppConfigType

func init() {
	AppConfig = &AppConfigType{
		Port:       8080,
		SocketPath: "/var/run/scarlett.sock",
	}
}
