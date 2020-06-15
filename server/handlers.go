package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/pabloos/http/greet"
	"github.com/pkg/errors"
)

func greetHandler(w http.ResponseWriter, r *http.Request) {
	var t greet.Greet

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, errors.Wrap(err, "Error while decoding the JSON payload:\n").Error(), 400)
		return
	}

	fmt.Fprintf(w, "Hello %s, from %s\n", t.Name, t.Location)
}
