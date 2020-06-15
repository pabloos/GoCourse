package main

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"
)

const publicCertFile = "../cert/public.crt"

func tlsConfig() *tls.Config {
	crt, err := ioutil.ReadFile(publicCertFile)
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
