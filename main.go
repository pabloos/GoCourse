package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(2)

	channel := make(chan string)

	go sender(channel, &wg)
	go receiver(channel, &wg)

	wg.Wait()
}

func sender(channel chan<- string, wg *sync.WaitGroup) {
	channel <- "Hello, GDG Marbella!"

	wg.Done()
}

func receiver(channel <-chan string, wg *sync.WaitGroup) {
	message := <-channel

	fmt.Println(message)

	wg.Done()
}
