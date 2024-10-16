package main

import (
	"errors"
	"fmt"
)

func main() {

	data := func(val string) (string, error) {

		if val == "" {
			return "Fail", errors.New("Data Not found")
		}

		return "Hello " + val, nil
	}

	if val, err := data("Deni"); err != nil {
		fmt.Println(val, err)
	} else {
		fmt.Println(val)
	}

}
