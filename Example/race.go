package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	// Dia akan lock dulu sampai selesai proses . lalu menjalankan goroutine yang lain

	var num int = 0
	var mutex sync.Mutex
	go func() {
		for i := 0; i < 500000; i++ {
			mutex.Lock()
			num++
			mutex.Unlock()
		}
	}()

	go func() {
		for i := 0; i < 500000; i++ {
			mutex.Lock()
			num++
			mutex.Unlock()
		}
	}()

	fmt.Println(num)
	time.Sleep(4 * time.Second)
	fmt.Println(num)
}
