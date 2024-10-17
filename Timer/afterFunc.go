package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	// Set Time out
	group := sync.WaitGroup{}

	group.Add(1)

	time.AfterFunc(5*time.Second, func() {
		fmt.Println("Jos")
		group.Done()
	})

	group.Wait()
}
