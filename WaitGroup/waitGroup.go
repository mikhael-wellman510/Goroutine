package main

import (
	"fmt"
	"sync"
)

func doPrint(wg *sync.WaitGroup, message string) {

	fmt.Println(message)

	wg.Done()
}

func main() {

	// Biasa nya kita menggunakan sleep untuk menunggu goRoutine selesai
	// Sekarang kita tidak perlu . tinggal menggunakan WaitGROUP

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		data := fmt.Sprintf("Data %d ", i)
		wg.Add(1)
		go doPrint(&wg, data)
	}

	wg.Wait()

}
