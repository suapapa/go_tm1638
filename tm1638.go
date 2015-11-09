// Copyright 2015, Homin Lee <homin.lee@suapapa.net>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tm1638

// TM1638 represent TM1638 base device
type TM1638 struct {
	TM16XX
}

// NewTM1638 retrives Pointer of a TM1638
func NewTM1638(data, clk, strobe int) (*TM1638, error) {
	activeDisplay := true
	intensity := byte(7)
	d, err := NewTM16XX(data, clk, strobe, activeDisplay, intensity)
	if err != nil {
		return nil, err
	}

	var r = TM1638{
		TM16XX: *d,
	}

	return &r, err
}
