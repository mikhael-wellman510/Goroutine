package main

import (
	"fmt"
	"time"
)

func RunHelloWorld() {
	fmt.Println("Hello World")
}

func main() {

	// Menggunakan keyword go untuk menggunakan goroutine
	go RunHelloWorld()
	fmt.Println("Ups...")

	time.Sleep(5 * time.Second) // 5 detik
}
