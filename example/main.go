// Copyright 2015, Homin Lee <homin.lee@suapapa.net>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"time"

	"github.com/suapapa/go_tm1638"
)

func main() {
	d, err := tm1638.NewTM1638(18, 23, 24)
	if err != nil {
		panic(err)
	}

	d.DisplayError()
	time.Sleep(3 * time.Second)
	d.DisplayDecNumber(12345678, 0, false)
	time.Sleep(3 * time.Second)
}
