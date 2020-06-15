package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type greet struct {
	Name     string `json:"name,omitempty"`
	Location string `json:"location,omitempty"`
}

func main() {
	httpServer := &http.Server{
		Addr: ":8080",
	}

	http.HandleFunc("/greet", greetHandler)

	if err := httpServer.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	var t greet

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	greet := fmt.Sprintf("Hello %s, from %s\n", t.Name, t.Location)

	w.Write([]byte(greet))
}
