package main

import (
	"fmt"

	"./tm1638"
)

func main() {
	fmt.Println("Hello world")
	d, err := tm1638.NewTM1638(18, 23, 24)
	if err != nil {
		panic(nil)
	}

	d.DisplayError()
	d.DisplayDigit(4, 1, true)
}
