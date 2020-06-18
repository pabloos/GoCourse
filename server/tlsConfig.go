package main

import (
	"crypto/tls"
	"io/ioutil"
	"log"
)

const (
	publicKeyFile  = "../cert/public.crt"
	privateKeyFile = "./cert/private.key"
)

func tlsConfig() *tls.Config {
	crt, err := ioutil.ReadFile(publicKeyFile)
	if err != nil {
		log.Fatal(err)
	}

	key, err := ioutil.ReadFile(privateKeyFile)
	if err != nil {
		log.Fatal(err)
	}

	cert, err := tls.X509KeyPair(crt, key)
	if err != nil {
		log.Fatal(err)
	}

	return &tls.Config{
		Certificates: []tls.Certificate{cert},
		ServerName:   "localhost",
	}
}
