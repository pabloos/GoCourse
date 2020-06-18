package main

import (
	"log"
	"net/http"
	"os"
)

func newServer() *http.Server {
	return &http.Server{
		Addr:      ":8080",
		Handler:   newMux(),
		TLSConfig: tlsConfig(),
		ErrorLog:  log.New(os.Stderr, "HTTP Server says: ", log.Llongfile),
	}
}
