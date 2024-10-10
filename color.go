package openrgb

import "image/color"

type Color struct {
	R uint8
	G uint8
	B uint8
}

func NewRGBColor(colors color.Color) Color {
	r, g, b, _ := colors.RGBA()
	return Color{
		R: uint8(r / 256),
		G: uint8(g / 256),
		B: uint8(b / 256),
	}
}

// Encode implements DataEncoder.
func (c Color) Encode(v Version, b *NetPacketDataBuilder) []byte {
	i := uint32(c.R)
	i |= uint32(c.G) << 8
	i |= uint32(c.B) << 16
	b.WriteUint32(i)
	return b.Bytes()
}

// Decode implements DataDecoder.
func (c *Color) Decode(v Version, p *NetPacketDataParser) error {
	i, err := p.ReadUint32()
	if err != nil {
		return err
	}
	c.R = uint8(i)
	c.G = uint8(i >> 8)
	c.B = uint8(i >> 16)
	return nil
}

// Helper for multiple led states.
type Colors []Color

var ColorBlack = Color{R: 0, G: 0, B: 0}
var ColorRed = Color{R: 255, G: 0, B: 0}
var ColorGreen = Color{R: 0, G: 255, B: 0}
var ColorBlue = Color{R: 0, G: 0, B: 255}
var ColorYellow = Color{R: 255, G: 255, B: 0}
var ColorCyan = Color{R: 0, G: 255, B: 255}
var ColorMagenta = Color{R: 255, G: 0, B: 255}
var ColorWhite = Color{R: 255, G: 255, B: 255}

// Create LED states with a predefined color.
func NewColors(numLeds int, c Color) Colors {
	colors := Colors(make([]Color, numLeds))
	for i := range colors {
		colors[i] = c
	}
	return colors
}

// Create LED states off state.
func NewBlackColors(numLeds int) Colors {
	return NewColors(numLeds, ColorBlack)
}

// Moves the light to the right.
func (c Colors) ShiftRight(i int) Colors {
	if i > len(c) {
		return NewBlackColors(len(c))
	}
	n := NewBlackColors(i)
	return append(n, c[:len(c)-i]...)
}

// Moves the light to the left.
func (c Colors) ShiftLeft(i int) Colors {
	if i > len(c) {
		return NewBlackColors(len(c))
	}
	n := NewBlackColors(i)
	return append(c[i:], n...)
}
