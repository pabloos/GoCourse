package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/pabloos/http/greet"
	"github.com/pkg/errors"
)

const (
	URL = "http://localhost:8080/greet"
)

func main() {
	client := &http.Client{}

	greetReq := greet.Greet{
		Name:     "John Doe",
		Location: "USA",
	}

	buf := &bytes.Buffer{} // b := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(greetReq)

	resp, err := client.Post(URL, "application/json; charset=utf-8", buf)
	if err != nil {
		fmt.Println(errors.Wrap(err, "Error on POSTing"))
	}
	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
