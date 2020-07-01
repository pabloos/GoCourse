package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	sites := []string{
		"https://www.google.com",
		"https://drive.gogle.com", // there's a missing 'o' in google!!
		"https://maps.google.com",
		"https://hangouts.google.com",
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	wg.Add(len(sites))

	for _, site := range sites {
		go func(site string) {
			defer wg.Done()

			res, err := http.Get(site)
			if err != nil {
				cancel()
			}

			select {
			case <-ctx.Done():
				return
			default:
				io.WriteString(os.Stdout, res.Status+"\n")
			}
		}(site)

		// time.Sleep(time.Second)
	}

	wg.Wait()
}
