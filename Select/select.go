package main

import (
	"fmt"
	"time"
)

// Mencari rata2
func getAverage(channel1 chan float64, data []int) {

	total := 0
	for _, num := range data {
		total = total + num
	}
	res := float64(total) / float64(len(data))

	channel1 <- res
}

// Mencari Max
func getMax(channel2 chan int, data []int) {

	max := data[0]

	for _, data := range data {
		if max < data {
			max = data
		}
	}

	channel2 <- max
}

func main() {
	// fungsi select yaitu untuk mendapatkan data yang selesai lebih dulu
	// Jika berbarengan . dia akan random mengeluarkan data nya

	// * Kenapa harus pakai Select ?
	// Karena dia bisa mengeluarkan yang selesai dahulu , tanpa menunggu atau memblocking function lain
	numbers := []int{10, 20, 40, 60, 5, 2}

	channel1 := make(chan float64)

	channel2 := make(chan int)

	go getAverage(channel1, numbers)
	go getMax(channel2, numbers)

	for i := 0; i < 2; i++ {

		// Ini bisa random yang mana aja selesai
		select {
		case data1 := <-channel1:
			fmt.Println(data1, "Average Succes")
		case data2 := <-channel2:
			fmt.Println(data2, "Max Succes")

		}

	}

	time.Sleep(6 * time.Second)
}
