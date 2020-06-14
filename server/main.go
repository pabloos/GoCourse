package main

import (
	"log"
	"net/http"
)

func main() {
	httpServer := &http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/", rootHandler)

	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World!\n"))
}
