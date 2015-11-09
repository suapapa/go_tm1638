package tm1638

// TM1638 represent TM1638 base device
type TM1638 struct {
	*TM16XX
}

// NewTM1638 retrives Pointer of a TM1638
func NewTM1638(data, clk, strobe int,
	display, activeDisplay, intensity byte) (*TM1638, error) {
	d, err := NewTM16XX(data, clk, strobe, display, activeDisplay, intensity)
	if err != nil {
		return nil, err
	}

	var r = TM1638{
		TM16XX: d,
	}

	return &r, err
}
