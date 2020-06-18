package main

import (
	"net/http"
	"time"
)

func newMux() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("/", Debug(index))
	mux.HandleFunc("/greet", Delay(2*time.Second, POST(greetHandler)))
	// mux.HandleFunc("/greet", POST(greetHandler))

	return mux
}
