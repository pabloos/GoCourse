package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	channel := make(chan string)

	defer close(channel)

	messages := []string{
		"Hello, GDG Marbella!",
		"Hello, Gophers!",
		"Hello, World!",
	}

	for _, msg := range messages {
		wg.Add(1)

		go sender(msg, channel, &wg)
	}

	go receiver(channel, &wg)

	wg.Wait()
}

func sender(msg string, channel chan<- string, wg *sync.WaitGroup) {
	channel <- msg

	wg.Done()
}

func receiver(channel <-chan string, wg *sync.WaitGroup) {
	for msg := range channel {
		fmt.Println(msg)
	}
}
