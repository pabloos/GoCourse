package main

import (
	"log"
)

func main() {
	if err := newServer().ListenAndServeTLS("", ""); err != nil {
		log.Fatal(err)
	}
}
