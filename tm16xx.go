package tm1638

import "github.com/davecheney/gpio"

// TM16XX represent a TM16XX module
type TM16XX struct {
	data, clk, strobe gpio.Pin
}

// NewTM16XX returns point of TM16XX which is initialized given gpio numbers
func NewTM16XX(data, clk, strobe int,
	display, activeDisplay, intensity byte) (d *TM16XX, err error) {
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
	if activeDisplay != 0 {
		v |= 8
	}
	d.sendCmd(0x80 | v)

	d.strobe.Clear()
	d.sendCmd(0xC0)
	for i := 0; i < 16; i++ {
		d.sendCmd(0x00)
	}

	d.strobe.Set()
	return
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
	for i := 0; i < 8; i++ {
		d.clk.Clear()
		v := data & 1
		if v&1 == 0 {
			d.data.Clear()
		} else {
			d.data.Set()
		}
		data >>= 1
		d.clk.Set()
	}
}
