package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func foo(c chan int, someValue int) {
	defer wg.Done()
	c <- someValue * 5
}

func main() {
	fooval := make(chan int, 10)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go foo(fooval, i)
	}

	wg.Wait()
	close(fooval)

	for item := range fooval {
		fmt.Println(item)
	}

	time.Sleep(time.Second * 2)
}
