package main

import (
	"net/http"
	"time"
)

func newClient() *http.Client {
	return &http.Client{
		Timeout: 1 * time.Second,

		Transport: &http.Transport{
			TLSClientConfig: tlsConfig(),
		},
	}
}
