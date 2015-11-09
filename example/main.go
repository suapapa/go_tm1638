package main

import "github.com/suapapa/go_tm1638"

func main() {
	d, err := tm1638.NewTM1638(18, 23, 24)
	if err != nil {
		panic(nil)
	}

	d.DisplayError()
}
