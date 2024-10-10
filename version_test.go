package openrgb_test

import (
	"testing"

	"github.com/csutorasa/go-openrgb-sdk"
)

func TestRequestProtocolVersionRequestEncode(t *testing.T) {
	builder := &openrgb.NetPacketDataBuilder{}
	req := &openrgb.RequestProtocolVersionRequest{
		ClientVersion: 3,
	}
	if req.NetPacketId() != openrgb.NetPacketIdRequestProtocolVersion {
		t.Log("invalid net packet ID")
		t.Fail()
	}
	req.Encode(0, builder)
	b := builder.Bytes()
	if len(b) != 4 {
		t.Fatalf("encoded data should be 4 long but was %d", len(b))
	}
	if b[0] != 3 || b[1] != 0 || b[2] != 0 || b[3] != 0 {
		t.Log("invalid data")
		t.Fail()
	}
}

func TestRequestProtocolVersionResponseDecode(t *testing.T) {
	p := openrgb.NewNetPacket(0, 0, []byte{
		// 4
		4, 0, 0, 0,
	})
	parser := p.DataParser()
	res := new(openrgb.RequestProtocolVersionResponse)
	if res.NetPacketId() != openrgb.NetPacketIdRequestProtocolVersion {
		t.Log("invalid net packet ID")
		t.Fail()
	}
	err := res.Decode(0, parser)
	if err != nil {
		t.Fatalf("decoding should not fail %s", err.Error())
	}
	if res.ServerVersion != 4 {
		t.Log("invalid data")
		t.Fail()
	}
	err = parser.AssertAllRead()
	if err != nil {
		t.Log("data is not fully read")
		t.Fail()
	}
}

func TestSetClientNameRequestEncode(t *testing.T) {
	builder := &openrgb.NetPacketDataBuilder{}
	req := &openrgb.SetClientNameRequest{
		ClientName: "test",
	}
	if req.NetPacketId() != openrgb.NetPacketIdSetClientName {
		t.Log("invalid net packet ID")
		t.Fail()
	}
	req.Encode(0, builder)
	b := builder.Bytes()
	if len(b) != 4 {
		t.Fatalf("encoded data should be 4 long but was %d", len(b))
	}
	// 116 101 115 116 = "test"
	if b[0] != 116 || b[1] != 101 || b[2] != 115 || b[3] != 116 {
		t.Log("invalid data")
		t.Fail()
	}
}
