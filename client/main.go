package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	URL = "http://localhost:8080"
)

func main() {
	client := &http.Client{}

	resp, err := client.Get(URL)
	if err != nil {
		fmt.Print(err)
		return
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
