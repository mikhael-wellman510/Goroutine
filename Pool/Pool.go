package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	pool := sync.Pool{}

	pool.Put("Mikhael")
	pool.Put("Wellman")
	pool.Put("Sitorus")

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			data := pool.Get()
			fmt.Println(data)

			pool.Put(data)
			wg.Done()
		}()
	}

	fmt.Println("Selesai")
	wg.Wait()
}
