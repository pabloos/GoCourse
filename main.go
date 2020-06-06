package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	msgMap := make(map[int]string)

	messages := []string{
		"Hello, GDG Marbella!",
		"Hello, Gophers!",
		"Hello, World!",
	}

	for i, msg := range messages {
		wg.Add(1)

		go func() {
			msgMap[i] = msg

			wg.Done()
		}()
	}

	wg.Wait()

	for i := range messages {
		fmt.Println(msgMap[i])
	}
}
