package openrgb_test

import (
	"testing"

	"github.com/csutorasa/go-openrgb-sdk"
)

func TestLedDecode(t *testing.T) {
	p := openrgb.NewNetPacket(0, 0, []byte{
		// 4
		4, 0,
		// "test"
		116, 101, 115, 116,
		// value
		0, 0, 0, 0,
	})
	parser := p.DataParser()
	l := &openrgb.Led{}
	err := l.Decode(0, parser)
	if err != nil {
		t.Fatalf("decoding should not fail %s", err.Error())
	}
	if l.LedName != "test" {
		t.Log("invalid data")
		t.Fail()
	}
	if l.LedValue != 0 {
		t.Log("invalid data")
		t.Fail()
	}
	err = parser.AssertAllRead()
	if err != nil {
		t.Log("data is not fully read")
		t.Fail()
	}
}
