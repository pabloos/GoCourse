package main

import "net/http"

func newServer() *http.Server {
	return &http.Server{
		Addr:    ":8080",
		Handler: newMux(),
	}
}
