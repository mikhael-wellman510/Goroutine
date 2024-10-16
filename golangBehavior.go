package main

import "time"

func sayHello() {
	time.Sleep(3 * time.Second)
	println("Hellooo ....!")
}

func sayGoodbye() {
	println("Good Byee .... !")
}

func main() {

	// Ini akan membuat golang asyncronous
	go sayHello()
	go sayGoodbye()

	time.Sleep(5 * time.Second)

}
