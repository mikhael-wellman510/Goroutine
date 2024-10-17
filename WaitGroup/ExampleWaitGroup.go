package main

import "sync"

func test1(wg *sync.WaitGroup) {
	defer wg.Done()
	println("Tes 1")

}

func test2(wg *sync.WaitGroup) {
	defer wg.Done()
	println("Tes 2")

}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go test1(&wg)

	wg.Add(1)
	go test2(&wg)

	wg.Wait()

}
