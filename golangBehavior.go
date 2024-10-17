package main

import (
	"sync"
	"time"
)

func sayHello() {
	// time.Sleep(3 * time.Second)
	var mutex sync.Mutex
	mutex.Lock()
	println("Hellooo ....!")
	mutex.Unlock()
}

func sayGoodbye() {
	var mutex sync.Mutex
	mutex.Lock()
	println("Good Byee .... !")
	mutex.Unlock()
}

func main() {

	// Ini akan membuat golang asyncronous

	go sayHello()

	go sayGoodbye()

	time.Sleep(1 * time.Second)

}
