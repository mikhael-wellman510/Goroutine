package main

import (
	"fmt"
	"time"
)

func main() {

	// Init
	data := make(chan string, 2) // Alfredo , Alex , Aldo

	defer close(data)

	go func() {
		data <- "Mikhael"
		data <- "Deni"
		data <- "Alfredo"
		data <- "Alex"
		data <- "aldo"

		fmt.Println("Finish") // Ini tidak akan tereksekusi . karena buffer nya penuh
	}()

	go func() {
		// Menerima Data
		res1 := <-data
		res2 := <-data
		fmt.Println(res1)
		fmt.Println(res2)
	}()

	time.Sleep(5 * time.Second)

}
