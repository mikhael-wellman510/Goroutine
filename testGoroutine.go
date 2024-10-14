package main

import (
	"fmt"
	"time"
)

func DisplayGoroutine(val int) {
	fmt.Println("Display : ", val)
}

func main() {

	// Menggunakan goroutine dia tidak akan berurutan
	for i := 0; i < 100000; i++ {
		go DisplayGoroutine(i)
	}

	time.Sleep(8 * time.Second)

}
