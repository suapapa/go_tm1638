// Copyright 2015, Homin Lee <homin.lee@suapapa.net>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tm1638

// TM1638Color... are color of leds
const (
	TM1638ColorNone = iota
	TM1638ColorRed
	TM1638ColorGreen
)

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

// DisplayDecNumberAt displays dec numbers at startPos on displays
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

// DisplayDecNumber displays dec numbers on display
func (d *TM1638) DisplayDecNumber(num uint64, dots byte, leadingZeros bool) {
	d.DisplayDecNumberAt(num, dots, 0, leadingZeros)
}

// DisplaySignedDecNumber displays signed dec numbers on display
func (d *TM1638) DisplaySignedDecNumber(num int64, dots byte, leadingZeros bool) {
	if num >= 0 {
		d.DisplayDecNumber(uint64(num), dots, leadingZeros)
		return
	}
	if -num > 99999999 {
		d.DisplayError()
		return
	}
	d.DisplayDecNumberAt(uint64(-num), dots, 1, leadingZeros)
	d.sendChar(0, fontDefault[13], (dots&0x80) != 0)
}

// DisplayBinNumber displays binary number on display
func (d *TM1638) DisplayBinNumber(num byte, dots byte) {
	for i := 0; i < d.displays; i++ {
		var v byte
		if num&(1<<byte(i)) != 0 {
			v = 1
		}
		d.DisplayDigit(v, d.displays-i-1, (dots&(1<<byte(i))) != 0)
	}
}

// SetLED sets led in pos to given color
func (d *TM1638) SetLED(color byte, pos byte) {
	d.sendData(pos<<1+1, color)
}

// SetLEDs sets leds
func (d *TM1638) SetLEDs(leds uint16) {
	for i := 0; i < d.displays; i++ {
		var color byte
		if (leds & uint16(1<<uint16(i))) != 0 {
			color |= TM1638ColorRed
		}
		if (leds & uint16(1<<uint16(i+8))) != 0 {
			color |= TM1638ColorGreen
		}

		d.SetLED(color, byte(i))

	}
}

func (d *TM1638) GetButton() byte {
	//TODO: fill here
	return 0
}
