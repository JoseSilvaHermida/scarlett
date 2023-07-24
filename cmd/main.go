// src/main.go
package main

import (
	"context"
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
	// Register the /hello endpoint handler.
	http.HandleFunc("/hello", handlers.HelloHandler)

	// Create a custom HTTP server with Unix domain socket and TCP listener.
	server := http.Server{}

	// Create listeners
	unixListener := createUnixListener(&server)
	defer unixListener.Close()

	httpListener := createHttpListener(&server)
	defer httpListener.Close()

	// Handle graceful shutdown on SIGINT and SIGTERM signals.
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan

	// Shutdown both servers gracefully on receiving the signal.
	err := server.Shutdown(context.TODO())
	if err != nil {
		log.Printf("Error while shutting down servers: %v", err)
	}

	// Remove the socket file on shutdown.
	os.Remove(config.AppConfig.SocketPath)

	log.Println("Servers gracefully shutdown")
}

func createUnixListener(server *http.Server) net.Listener {
	// Remove the socket file if it already exists.
	os.Remove(config.AppConfig.SocketPath)

	// Create a listener for the Unix domain socket.
	unixListener, err := net.Listen("unix", config.AppConfig.SocketPath)
	if err != nil {
		log.Fatalf("Failed to create Unix socket listener: %v", err)
	}

	// Start the HTTP server on the Unix domain socket in a separate goroutine.
	go func() {
		log.Printf("Starting server on %s", config.AppConfig.SocketPath)
		err := server.Serve(unixListener)
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start Unix socket server: %v", err)
		}
	}()

	return unixListener
}

func createHttpListener(server *http.Server) net.Listener {
	// Create a listener for the HTTP server on localhost:8080.
	httpListener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", config.AppConfig.Port))
	if err != nil {
		log.Fatalf("Failed to create HTTP listener: %v", err)
	}

	// Start the HTTP server on localhost:8080 in a separate goroutine.
	go func() {
		log.Println("Starting server on localhost:8080")
		err := server.Serve(httpListener)
		if err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start HTTP server: %v", err)
		}
	}()

	return httpListener
}
