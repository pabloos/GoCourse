package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
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

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Print(err)
		return
	}

	fmt.Print(string(body))
}
