package main

import (
	"fmt"
	"runtime"
)

func main() {
	totalCpu := runtime.NumCPU()

	fmt.Println(totalCpu)
}
