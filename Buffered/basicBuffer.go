package main

import (
	"fmt"
	"time"
)

func main() {

	// jika buffer penuh . maka buffer akan memblocking semua baris kode

	data := make(chan string, 2)

	go func() {
		data <- "Mikhael"
		data <- "Deni"
		data <- "Alfredo"

		fmt.Println("Running ...") // Block ini tidak akan terbaca sebelum data di terima
	}()

	res1 := <-data // Menerima data . otomatis isi buffer akan berkurang
	fmt.Println(res1)
	time.Sleep(4 * time.Second)

}
