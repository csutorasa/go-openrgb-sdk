package openrgb_test

import (
	"testing"

	"github.com/csutorasa/go-openrgb-sdk"
)

func TestColorDecode(t *testing.T) {
	p := openrgb.NewNetPacket(0, 0, []byte{255, 240, 128, 0})
	parser := p.DataParser()
	var c openrgb.Color
	err := c.Decode(0, parser)
	if err != nil {
		t.Fatalf("decoding should not fail %s", err.Error())
	}
	if c.R != 255 || c.G != 240 || c.B != 128 {
		t.Log("invalid data")
		t.Fail()
	}
	err = parser.AssertAllRead()
	if err != nil {
		t.Log("data is not fully read")
		t.Fail()
	}
}

func TestColorEncode(t *testing.T) {
	builder := &openrgb.NetPacketDataBuilder{}
	c := openrgb.Color{R: 255, G: 240, B: 128}
	c.Encode(0, builder)
	b := builder.Bytes()
	if len(b) != 4 {
		t.Fatalf("encoded data should be 4 long but was %d", len(b))
	}
	if b[0] != 255 || b[1] != 240 || b[2] != 128 || b[3] != 0 {
		t.Log("invalid data")
		t.Fail()
	}
}
