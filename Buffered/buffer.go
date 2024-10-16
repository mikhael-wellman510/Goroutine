package main

import (
	"fmt"
	"time"
)

func main() {

	data := make(chan int, 3)

	go func() {

		for {
			i := <-data
			fmt.Println("result : ", i)
		}
	}()

	for i := 0; i < 2; i++ {
		// Dia akan kirim data sampai buffer penuh . lalu eksekusi function di atas
		fmt.Println("Send Data ", i)
		data <- i
	}

	fmt.Println("Finish")

	time.Sleep(5 * time.Second)

}
