package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/pabloos/http/greet"
)

const (
	URL = "https://localhost:8080"
)

func main() {
	client := newClient()

	resp1, err := client.Get(URL)
	if err != nil {
		fmt.Println(err)
	}
	defer resp1.Body.Close()

	io.Copy(os.Stdout, resp1.Body)

	greetReq := greet.Greet{
		Name:     "John Doe",
		Location: "USA",
	}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(greetReq)

	greetURL := fmt.Sprintf("%s/%s", URL, "greet")

	resp, err := client.Post(greetURL, "application/json; charset=utf-8", buf)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)

	fmt.Fprintf(os.Stdout, "Status code: %d\n", resp.StatusCode)
}
