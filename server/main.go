package main

import (
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type greet struct {
	Name     string `json:"name,omitempty"`
	Location string `json:"location,omitempty"`
}

func main() {
	httpServer := &http.Server{
		Addr:      ":8080",
		TLSConfig: tlsConfig(),
	}

	http.HandleFunc("/greet", greetHandler)

	if err := httpServer.ListenAndServeTLS("", ""); err != nil {
		log.Fatal(err)
	}
}

func tlsConfig() *tls.Config {
	crt, err := ioutil.ReadFile("../cert/public.crt")
	if err != nil {
		log.Fatal(err)
	}

	key, err := ioutil.ReadFile("./cert/private.key")
	if err != nil {
		log.Fatal(err)
	}

	cert, err := tls.X509KeyPair(crt, key)
	if err != nil {
		log.Fatal(err)
	}

	return &tls.Config{
		Certificates: []tls.Certificate{cert},
		ServerName:   "localhost",
	}
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
	var t greet

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	fmt.Fprintf(w, "Hello %s, from %s\n", t.Name, t.Location)
}
