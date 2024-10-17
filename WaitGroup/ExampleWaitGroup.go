package main

import "sync"

func test1() {
	println("Tes 1")

}

func test2() {
	println("Tes 2")
}

func main() {

	var wg sync.WaitGroup

	wg.Add(1)
	go test1()
	wg.Done()
	wg.Add(1)
	go test2()
	wg.Done()
	wg.Wait()
	wg.Wait()

}
