package main

import (
	"fmt"
	"time"
)

func dataFast(data chan string) {

	time.Sleep(1 * time.Second)
	data <- "Fast"
}

func dataLow(data chan string) {
	time.Sleep(2 * time.Second)
	data <- "Low"
}

func main() {

	channel1 := make(chan string)
	channel2 := make(chan string)

	go dataFast(channel1)
	go dataLow(channel2)

	counter := 0
	res1 := ""
	res2 := ""
	for {
		select {
		case data1 := <-channel1:
			fmt.Println("Data Fast Succes : ", data1)
			res1 = data1
			counter++
		case data2 := <-channel2:
			fmt.Println("Data Low Succes : ", data2)
			res2 = data2
			counter++
		// Default akan terus berjalan selagi belum menerima hasil dari go routine nya
		default:
			fmt.Print("Menunggu Data")

		}

		if counter == 2 {
			break
		}
	}

	fmt.Println(res1)
	fmt.Println(res2)

}
