package main

import (
	"io"
	"net/http"
	"os"
)

func main() {
	sites := []string{
		"https://www.google.com",
		"https://drive.google.com",
		"https://maps.google.com",
		"https://hangouts.google.com",
	}

	for _, site := range sites {
		go func(site string) {
			res, err := http.Get(site)
			if err != nil {
			}

			io.WriteString(os.Stdout, res.Status+"\n")
		}(site)
	}
}
