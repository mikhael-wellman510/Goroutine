package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var num int = 0
	var mutex sync.Mutex

	func() {
		for i := 0; i < 1000; i++ {

			go func() {
				for i := 0; i < 100; i++ {
					mutex.Lock()
					num++
					mutex.Unlock()
				}
			}()
		}

	}()

	time.Sleep(4 * time.Second)
	fmt.Println(num)
}
