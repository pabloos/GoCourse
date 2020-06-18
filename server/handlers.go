package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/pabloos/http/greet"
)

func index(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "You are on the index page\n")
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	var t greet.Greet

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	fmt.Fprintf(w, "Hello %s, from %s\n", t.Name, t.Location)
}
