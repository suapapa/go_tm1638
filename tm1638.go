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

	d.displays = 8

	var r = TM1638{
		TM16XX: *d,
	}

	return &r, err
}

// DisplayHexNumber displays hex numbers on displays
func (d *TM1638) DisplayHexNumber(num uint64, dots byte, leadingZeros bool) {
	for i := 0; i < d.displays; i++ {
		if !leadingZeros && num == 0 {
			d.ClearDigit(d.displays-i-1, dots&(1<<uint8(i)) != 0)
		} else {
			d.DisplayDigit(byte(num)&0xF, d.displays-i-1, dots&(1<<uint8(i)) != 0)
			num >>= 4
		}
	}
}

// DisplayDecNumberAt display dec numbers at startPos on displays
func (d *TM1638) DisplayDecNumberAt(num uint64, dots byte, startPos int, leadingZeros bool) {
	if num > 99999999 {
		d.DisplayError()
		return
	}
	for i := 0; i < d.displays-startPos; i++ {
		if num != 0 {
			d.DisplayDigit(byte(num%10), d.displays-i-1, dots&(1<<uint8(i)) != 0)
			num /= 10
		} else {
			if leadingZeros {
				d.DisplayDigit(0, d.displays-i-1, dots&(1<<uint8(i)) != 0)
			} else {
				d.ClearDigit(d.displays-i-1, dots&(1<<uint8(i)) != 0)
			}
		}
	}
}

// DisplayDecNumber display dec numbers on display
func (d *TM1638) DisplayDecNumber(num uint64, dots byte, leadingZeros bool) {
	d.DisplayDecNumberAt(num, dots, 0, leadingZeros)
}
