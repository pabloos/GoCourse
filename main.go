package main

import (
	"fmt"
	"sync"
)

type msgMap struct {
	values map[int]string

	sync.Mutex
}

func main() {
	var wg sync.WaitGroup

	msgsIndexed := new(msgMap)

	msgsIndexed.values = make(map[int]string)

	messages := []string{
		"Hello, GDG Marbella!",
		"Hello, Gophers!",
		"Hello, World!",
	}

	for i, msg := range messages {
		wg.Add(1)

		go func(msg string, index int) {
			defer wg.Done()

			msgsIndexed.Lock()
			msgsIndexed.values[index] = msg
			msgsIndexed.Unlock()
		}(msg, i)
	}

	wg.Wait()

	for i := range messages {
		fmt.Println(msgsIndexed.values[i])
	}
}
