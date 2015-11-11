// Copyright 2015, Homin Lee <homin.lee@suapapa.net>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package tm1638

import (
	"sync"

	"github.com/davecheney/gpio"
)

// TM16XX represent a TM16XX module
type TM16XX struct {
	sync.Mutex

	data, clk, strobe gpio.Pin
	displays          int
}

// NewTM16XX returns point of TM16XX which is initialized given gpio numbers
func NewTM16XX(data, clk, strobe int,
	activeDisplay bool, intensity byte) (*TM16XX, error) {

	var d = TM16XX{}
	var err error

	d.data, err = gpio.OpenPin(data, gpio.ModeOutput)
	if err != nil {
		return nil, err
	}
	d.clk, err = gpio.OpenPin(clk, gpio.ModeOutput)
	if err != nil {
		return nil, err
	}
	d.strobe, err = gpio.OpenPin(strobe, gpio.ModeOutput)
	if err != nil {
		return nil, err
	}

	d.data.SetMode(gpio.ModeOutput)
	d.clk.SetMode(gpio.ModeOutput)
	d.strobe.SetMode(gpio.ModeOutput)

	d.strobe.Set()
	d.clk.Set()

	d.sendCmd(0x40)
	v := min(7, intensity)
	if activeDisplay {
		v |= 8
	}
	d.sendCmd(0x80 | v)

	d.strobe.Clear()
	d.sendCmd(0xC0)
	for i := 0; i < 16; i++ {
		d.sendCmd(0x00)
	}

	d.strobe.Set()
	return &d, nil
}

// SetupDisplay initialized the display
func (d *TM16XX) SetupDisplay(active bool, intensity byte) {
	d.Lock()
	v := min(7, intensity)
	if active {
		v |= 8
	}
	d.sendCmd(0x80 | v)

	d.strobe.Clear()
	d.clk.Clear()
	d.clk.Set()
	d.strobe.Set()
	d.Unlock()
}

// DisplayDigit displays a digit
func (d *TM16XX) DisplayDigit(digit byte, pos int, dot bool) {
	d.sendChar(byte(pos), fontNumber[digit&0x0F], dot)
}

// DisplayError display Error
func (d *TM16XX) DisplayError() {
	d.setDisplay(fontErrorData)
}

// ClearDigit clear digit in given position
func (d *TM16XX) ClearDigit(pos int, dot bool) {
	d.sendChar(byte(pos), 0, dot)
}

func (d *TM16XX) setDisplay(val []byte) {
	for i, c := range val {
		d.sendChar(byte(i), c, false)
	}
}

func (d *TM16XX) sendCmd(cmd byte) {
	d.strobe.Clear()
	d.send(cmd)
	d.strobe.Set()
}

func (d *TM16XX) sendData(addr, data byte) {
	d.sendCmd(0x44)
	d.strobe.Clear()
	d.send(0xC0 | addr)
	d.send(data)
	d.strobe.Set()
}

func (d *TM16XX) send(data byte) {
	d.Lock()
	for i := 0; i < 8; i++ {
		d.clk.Clear()
		if data&1 == 0 {
			d.data.Clear()
		} else {
			d.data.Set()
		}
		data >>= 1
		d.clk.Set()
	}
	d.Unlock()
}

func (d *TM16XX) receive() (temp byte) {
	d.Lock()
	d.data.SetMode(gpio.ModeInput)
	d.data.Set() // TODO: is this makes data pin pull up?

	for i := 0; i < 8; i++ {
		temp >>= 1
		d.clk.Clear()
		if d.data.Get() {
			temp |= 0x80
		}
		d.clk.Set()
	}

	d.data.SetMode(gpio.ModeOutput)
	d.data.Clear()
	d.Unlock()

	return
}

func (d *TM16XX) sendChar(pos byte, data byte, dot bool) {
	if dot {
		data |= 0x80
	}
	d.sendData(pos<<1, data)
}
