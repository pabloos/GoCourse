package main

import (
	"log"
)

func main() {
	if err := newServer().ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
