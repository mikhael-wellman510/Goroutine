package main

import "fmt"

func main() {
	orderSomeFood("pizza")
	orderSomeFood("mie")

}

func orderSomeFood(menu string) {

	defer fmt.Println("Defer")

	if menu == "pizza" {
		fmt.Println("Selamat menikmati ", menu)
		return
	}
}
