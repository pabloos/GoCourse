package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	URL = "http://localhost:8080"
)

type greet struct {
	Name     string `json:"name,omitempty"`
	Location string `json:"location,omitempty"`
}

func main() {
	client := &http.Client{}

	greetReq := greet{
		Name:     "John Doe",
		Location: "USA",
	}

	buf := &bytes.Buffer{} // b := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(greetReq)

	greetURL := fmt.Sprintf("%s/%s", URL, "greet")

	resp, err := client.Post(greetURL, "application/json; charset=utf-8", buf)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
