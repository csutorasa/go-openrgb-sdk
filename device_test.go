package openrgb_test

import (
	"testing"

	"github.com/csutorasa/go-openrgb-sdk"
)

func TestRequestControllerCountRequestEncode(t *testing.T) {
	builder := &openrgb.NetPacketDataBuilder{}
	req := &openrgb.RequestControllerCountRequest{}
	if req.NetPacketId() != openrgb.NetPacketIdRequestControllerCount {
		t.Log("invalid net packet ID")
		t.Fail()
	}
	req.Encode(0, builder)
	b := builder.Bytes()
	if len(b) != 0 {
		t.Fatalf("encoded data should be 0 long but was %d", len(b))
	}
}

func TestRequestControllerCountResponseDecode(t *testing.T) {
	p := openrgb.NewNetPacket(0, 0, []byte{
		// 4
		4, 0, 0, 0,
	})
	parser := p.DataParser()
	res := new(openrgb.RequestControllerCountResponse)
	if res.NetPacketId() != openrgb.NetPacketIdRequestControllerCount {
		t.Log("invalid net packet ID")
		t.Fail()
	}
	err := res.Decode(0, parser)
	if err != nil {
		t.Fatalf("deoding should not fail %s", err.Error())
	}
	if res.Count != 4 {
		t.Log("invalid data")
		t.Fail()
	}
	err = parser.AssertAllRead()
	if err != nil {
		t.Log("data is not fully read")
		t.Fail()
	}
}

func TestRequestControllerDataRequestEncode(t *testing.T) {
	builder := &openrgb.NetPacketDataBuilder{}
	req := &openrgb.RequestControllerDataRequest{}
	if req.NetPacketId() != openrgb.NetPacketIdRequestControllerData {
		t.Log("invalid net packet ID")
		t.Fail()
	}
	req.Encode(0, builder)
	b := builder.Bytes()
	if len(b) != 0 {
		t.Fatalf("encoded data should be 0 long but was %d", len(b))
	}
	req.Encode(1, builder)
	b = builder.Bytes()
	if len(b) != 4 {
		t.Fatalf("encoded data should be 4 long but was %d", len(b))
	}
	if b[0] != 1 || b[1] != 0 || b[2] != 0 || b[3] != 0 {
		t.Log("invalid data")
		t.Fail()
	}
}

func TestRequestControllerDataResponseDecode(t *testing.T) {
	res := new(openrgb.RequestControllerDataResponse)
	if res.NetPacketId() != openrgb.NetPacketIdRequestControllerData {
		t.Log("invalid net packet ID")
		t.Fail()
	}
	// TODO test decode
}

func TestDeviceListUpdatedResponseDecode(t *testing.T) {
	p := openrgb.NewNetPacket(0, 0, []byte{})
	parser := p.DataParser()
	res := new(openrgb.DeviceListUpdatedResponse)
	if res.NetPacketId() != openrgb.NetPacketIdDeviceListUpdated {
		t.Log("invalid net packet ID")
		t.Fail()
	}
	err := res.Decode(0, parser)
	if err != nil {
		t.Fatalf("deoding should not fail %s", err.Error())
	}
	err = parser.AssertAllRead()
	if err != nil {
		t.Log("data is not fully read")
		t.Fail()
	}
}
