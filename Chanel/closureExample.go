package main

import (
	"fmt"
	"time"
)

func main() {

	// Buat Channel
	data := make(chan string)

	sayHello := func(name string) {
		time.Sleep(2 * time.Second)

		say := fmt.Sprintf("Hi ... %s", name)
		// Kirim Channel
		data <- say

		fmt.Println("Ambil Chanel Succes")

	}

	// Gunakan Go Routine untuk mengirim/Menggunakan Channel
	go sayHello("Mikhael")
	go sayHello("Deni")
	go sayHello("Aldi")

	message1 := <-data
	message2 := <-data
	message3 := <-data

	fmt.Println(message1)
	fmt.Println(message2)
	fmt.Println(message3)

	time.Sleep(4 * time.Second)

}
