package main

import (
	"net/http"
)

func newMux() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/greet", greetHandler)

	return mux
}
