package main

import "fmt"

func giveMeResponse(val chan string) {
	val <- "Mikhael response you !"

}

func main() {

	// Init
	data := make(chan string)

	// Kirim channel ke parameter function
	// Parameter nya menjadi pass by reference , jdi tidak membuat duplikat
	go giveMeResponse(data)

	response := <-data

	fmt.Println(response)
}
