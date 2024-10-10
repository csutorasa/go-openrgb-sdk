package openrgb_test

import (
	"bytes"
	"errors"
	"testing"

	"github.com/csutorasa/go-openrgb-sdk"
)

func TestNetPacketDecoder(t *testing.T) {
	buf := bytes.NewBuffer([]byte{
		// 79 82 71 66 = "ORGB"
		79, 82, 71, 66,
		// 44 * 1 + 1 * 256 = 300
		44, 1, 0, 0,
		// 232 * 1 + 3 * 256 = 1000
		232, 3, 0, 0,
		// 4 * 1 = 4
		4, 0, 0, 0,
		// 116 101 115 116 = "test"
		116, 101, 115, 116,
	})
	d := openrgb.NewNetPacketDecoder(buf)
	p, err := d.Decode()
	if err != nil {
		t.Fatalf("decoding should not fail %s", err.Error())
	}
	if p.Header.PktMagic != "ORGB" {
		t.Log("invalid magic value")
		t.Fail()
	}
	if p.Header.PktDevIdx != 300 {
		t.Log("invalid device id")
		t.Fail()
	}
	if p.Header.PktId != 1000 {
		t.Log("invalid command id")
		t.Fail()
	}
	if p.Header.PktSize != 4 {
		t.Log("invalid size")
		t.Fail()
	}
	if string(p.Data) != "test" {
		t.Log("invalid data")
		t.Fail()
	}
}

func TestNetPacketDataParserAssertSize(t *testing.T) {
	p := openrgb.NewNetPacket(0, 0, []byte{0, 0})
	parser := p.DataParser()
	err := parser.AssertAllRead()
	var x *openrgb.UnprocessedResponseError
	if err == nil {
		t.Log("should return error on unread data")
		t.Fail()
	} else if !errors.As(err, &x) {
		t.Log("should return correct error on unread data")
		t.Fail()
	}
	parser.ReadUint16()
	err = parser.AssertAllRead()
	if err != nil {
		t.Log("should not return error on read data")
		t.Fail()
	}
}

func TestNetPacketDataParserReadUint16(t *testing.T) {
	p := openrgb.NewNetPacket(0, 0, []byte{44, 1})
	parser := p.DataParser()
	i, err := parser.ReadUint16()
	if err != nil {
		t.Fatalf("decoding should not fail %s", err.Error())
	}
	if i != 300 {
		t.Log("invalid data")
		t.Fail()
	}
	err = parser.AssertAllRead()
	if err != nil {
		t.Log("data is not fully read")
		t.Fail()
	}
}

func TestNetPacketDataParserReadUint32(t *testing.T) {
	p := openrgb.NewNetPacket(0, 0, []byte{44, 1, 0, 0})
	parser := p.DataParser()
	i, err := parser.ReadUint32()
	if err != nil {
		t.Fatalf("decoding should not fail %s", err.Error())
	}
	if i != 300 {
		t.Log("invalid data")
		t.Fail()
	}
	err = parser.AssertAllRead()
	if err != nil {
		t.Log("data is not fully read")
		t.Fail()
	}
}

func TestNetPacketDataParserReadInt32(t *testing.T) {
	p := openrgb.NewNetPacket(0, 0, []byte{44, 1, 0, 0})
	parser := p.DataParser()
	i, err := parser.ReadInt32()
	if err != nil {
		t.Fatalf("decoding should not fail %s", err.Error())
	}
	if i != 300 {
		t.Log("invalid data")
		t.Fail()
	}
	err = parser.AssertAllRead()
	if err != nil {
		t.Log("data is not fully read")
		t.Fail()
	}
}

func TestNetPacketDataParserReadBytes(t *testing.T) {
	p := openrgb.NewNetPacket(0, 0, []byte{116, 101, 115, 116})
	parser := p.DataParser()
	b, err := parser.ReadBytes(4)
	if err != nil {
		t.Fatalf("decoding should not fail %s", err.Error())
	}
	if len(b) != 4 {
		t.Fatalf("decoded data should be 4 long but was %d", len(b))
	}
	if string(b) != "test" {
		t.Log("invalid data")
		t.Fail()
	}
	err = parser.AssertAllRead()
	if err != nil {
		t.Log("data is not fully read")
		t.Fail()
	}
}

func TestNetPacketDataParserReadString(t *testing.T) {
	p := openrgb.NewNetPacket(0, 0, []byte{4, 0, 116, 101, 115, 116})
	parser := p.DataParser()
	s, err := parser.ReadString()
	if err != nil {
		t.Fatalf("decoding should not fail %s", err.Error())
	}
	if len(s) != 4 {
		t.Fatalf("decoded data should be 4 long but was %d", len(s))
	}
	if s != "test" {
		t.Log("invalid data")
		t.Fail()
	}
	err = parser.AssertAllRead()
	if err != nil {
		t.Log("data is not fully read")
		t.Fail()
	}
}
