package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// spaceX api GET endpoint to get the last launch
const lstLaunchesURL = "https://api.spacexdata.com/v3/launches/latest"

func main() {
	resp, err := http.Get(lstLaunchesURL)
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()

	io.Copy(os.Stdout, resp.Body)
}
