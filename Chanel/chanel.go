package main

import (
	"fmt"
	"time"
)

func main() {

	// Channel jika tidak di kirim  . maka akan deadlock.
	// Dan jika tidak diterima , maka akan sleep dan blocking

	name := make(chan string)

	sayHello := func(nama string) {
		hello := fmt.Sprintf("Hi ,.. %s ", nama)
		time.Sleep(2 * time.Second)
		fmt.Println("Kirim data ke Chanel ...")
		name <- hello // jika ini tidak ada . maka dia akan blocking ,dan tidak menjalankan baris program yang di bawah nya
		fmt.Println("Selesai Kirim Data ke Chanel")
	}

	go sayHello("Mike")

	data := <-name

	fmt.Println("hasil channel ... ", data)

	time.Sleep(4 * time.Second)
}
