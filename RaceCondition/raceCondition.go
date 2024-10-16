package main

import (
	"fmt"
	"time"
)

func main() {
	x := 0

	for i := 0; i < 10; i++ {

		go func() {
			for i := 0; i < 4; i++ {
				x = x + 1
			}
		}()
	}

	fmt.Println(x)
	time.Sleep(3 * time.Second)
}
