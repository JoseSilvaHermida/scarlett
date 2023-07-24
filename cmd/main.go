// src/main.go
package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"josesilvahermida/scarlett/internal/config"
	"josesilvahermida/scarlett/internal/handlers"
)

func main() {
	fmt.Printf("Application Port: %d\n", config.AppConfig.Port)
	fmt.Printf("Socket Path: %s\n", config.AppConfig.SocketPath)

	// Create a Unix domain socket path.
	socketPath := config.AppConfig.SocketPath

	// Remove the socket file if it already exists.
	os.Remove(socketPath)

	// Register the /hello endpoint handler.
	http.HandleFunc("/hello", handlers.HelloHandler)

	// Create a custom HTTP server with Unix domain socket and TCP listener.
	server := http.Server{}

	// Create a listener for the Unix domain socket.
	unixListener, err := net.Listen("unix", socketPath)
	if err != nil {
		log.Fatalf("Failed to create Unix socket listener: %v", err)
	}
	defer unixListener.Close()

	// Start the HTTP server on the Unix domain socket in a separate goroutine.
	go func() {
		log.Printf("Starting server on %s", socketPath)
		err := server.Serve(unixListener)
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start Unix socket server: %v", err)
		}
	}()

	// Create a listener for the HTTP server on localhost:8080.
	httpListener, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", config.AppConfig.Port))
	if err != nil {
		log.Fatalf("Failed to create HTTP listener: %v", err)
	}
	defer httpListener.Close()

	// Start the HTTP server on localhost:8080 in a separate goroutine.
	go func() {
		log.Println("Starting server on localhost:8080")
		err := server.Serve(httpListener)
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start HTTP server: %v", err)
		}
	}()

	// Handle graceful shutdown on SIGINT and SIGTERM signals.
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan

	// Shutdown both servers gracefully on receiving the signal.
	err = server.Shutdown(nil)
	if err != nil {
		log.Printf("Error while shutting down servers: %v", err)
	}

	// Remove the socket file on shutdown.
	os.Remove(socketPath)

	log.Println("Servers gracefully shutdown")
}
