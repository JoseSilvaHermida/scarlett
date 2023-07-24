package handlers

import (
	"net/http"
)

// HelloHandler handles the /hello endpoint.
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	// Write the response "Hello, World!" to the client.
	w.Write([]byte("Hello, World!"))
}
