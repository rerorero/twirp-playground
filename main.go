package main

import (
	"net/http"

	"github.com/rerorero/twrip-playground/internal/haberdasherserver"
	"github.com/rerorero/twrip-playground/rpc/haberdasher"
)

// Run the implementation in a local server
func main() {
	server := &haberdasherserver.Server{} // implements Haberdasher interface
	twirpHandler := haberdasher.NewHaberdasherServer(server, nil)

	http.ListenAndServe(":8080", twirpHandler)
}
