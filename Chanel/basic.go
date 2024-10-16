package main

import "fmt"

func main() {

	// Inisialisasi menggunakan make
	message := make(chan string)

	// Kirim menggunakan function
	go func() {
		message <- "Mikhael"
	}()

	result := <-message
	fmt.Println(result)
}
