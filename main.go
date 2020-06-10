package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const lstLaunchesURL = "https://api.spacexdata.com/v3/launches/latest"

func main() {
	resp, err := http.Get(lstLaunchesURL)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(body))
}
