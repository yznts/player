package main

import (
	"flag"
	"net/http"

	"github.com/yznts/player/pkg/server"
	"github.com/yznts/player/pkg/sync"
)

func main() {
	// Accept arguments
	HTTP := flag.String("http", ":8080", "HTTP service address")
	flag.Parse()

	// Create a new server
	_server := server.NewServer(server.Options{
		TemplateGlob: "web/template/*.go.html",
		Repository:   sync.NewJson("data.json"),
	})

	// Start the server
	println("Server started on", *HTTP)
	if err := http.ListenAndServe(*HTTP, _server); err != nil {
		panic(err)
	}
}
