package main

import (
	"github.com/jkulton/platter/internal/auth"
	"github.com/jkulton/platter/internal/options"
	"log"
	"net/http"
)

func main() {
	options := options.ParseOptions()
	handler := http.FileServer(http.Dir(options.Directory))

	http.Handle("/", auth.AuthMiddleware(handler, options.Auth))

	log.Printf("Serving files from %s at http://%s:%s\n", options.Directory, options.Address, options.Port)
	log.Fatal(http.ListenAndServe(options.Address+":"+options.Port, nil))
}
