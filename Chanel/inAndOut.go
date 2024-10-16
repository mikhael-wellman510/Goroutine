package main

import (
	"fmt"
	"time"
)

// Khusus Hanya mengirim data ke channel
func OnlyIn(data chan<- string, name string) {
	data <- name
}

// Khusus hanya menerima data ke chanel
func OnlyOut(val <-chan string) {
	result := <-val

	fmt.Println("Get Data : ", result)
}

func main() {

	// Membuat chanel khusus untuk mengirim dan menerima data saja
	data := make(chan string)

	go OnlyIn(data, "mikhael")
	go OnlyOut(data)

	fmt.Println("Finnisih")
	time.Sleep(4 * time.Second)
}
