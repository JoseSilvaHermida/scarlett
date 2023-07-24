// src/main.go
package main

import (
	"fmt"

	"josesilvahermida/scarlett/internal/config"
)

func main() {
	fmt.Printf("Application Port: %d\n", config.AppConfig.Port)
	fmt.Printf("Log Level: %s\n", config.AppConfig.LogLevel)
}
