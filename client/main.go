package main

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

const (
	URL = "https://localhost:8080"
)

type greet struct {
	Name     string `json:"name,omitempty"`
	Location string `json:"location,omitempty"`
}

func main() {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: tlsConfig(),
		},
	}

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

func tlsConfig() *tls.Config {
	crt, err := ioutil.ReadFile("../cert/public.crt")
	if err != nil {
		log.Fatal(err)
	}

	rootCAs := x509.NewCertPool()
	rootCAs.AppendCertsFromPEM(crt)

	return &tls.Config{
		RootCAs:            rootCAs,
		InsecureSkipVerify: false,
		ServerName:         "localhost",
	}
}
