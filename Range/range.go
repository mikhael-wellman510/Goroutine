package main

import (
	"fmt"
	"strconv"
)

func loopData(val chan string) {

	for i := 0; i < 10; i++ {
		val <- "Loop ke  : " + strconv.Itoa(i)
	}

	close(val) // menutup data . karena saat looping di range nya , dia akan menunggu data terus menerus
}

func main() {

	channel := make(chan string)

	go loopData(channel)

	// Jika tidak di close() , dia akan menunggu data terus
	for data := range channel {
		fmt.Println(data)
	}

}
