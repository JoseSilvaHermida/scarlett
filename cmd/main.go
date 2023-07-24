// src/main.go
package main

import (
	"fmt"
	"log"
	"net/http"

	"josesilvahermida/scarlett/internal/config"
	"josesilvahermida/scarlett/internal/handlers"
)

func main() {
	fmt.Printf("Application Port: %d\n", config.AppConfig.Port)
	fmt.Printf("Log Level: %s\n", config.AppConfig.LogLevel)

	// Register the /hello endpoint handler.
	http.HandleFunc("/hello", handlers.HelloHandler)

	// Start the HTTP server on port 8080.
	log.Println("Starting server on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
