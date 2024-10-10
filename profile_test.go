package openrgb_test

import (
	"testing"

	"github.com/csutorasa/go-openrgb-sdk"
)

func TestRequestProfileListRequestEncode(t *testing.T) {
	builder := &openrgb.NetPacketDataBuilder{}
	req := &openrgb.RequestProfileListRequest{}
	if req.NetPacketId() != openrgb.NetPacketIdRequestProfileList {
		t.Log("invalid net packet ID")
		t.Fail()
	}
	req.Encode(0, builder)
	b := builder.Bytes()
	if len(b) != 0 {
		t.Fatalf("encoded data should be 0 long but was %d", len(b))
	}
}

func TestRequestProfileListResponseDecode(t *testing.T) {
	p := openrgb.NewNetPacket(0, 0, []byte{
		//
		17, 0, 0, 0,
		// 2
		2, 0,
		// 4
		4, 0,
		// test
		116, 101, 115, 116,
		// 5
		5, 0,
		// test1
		116, 101, 115, 116, 49,
	})
	res := new(openrgb.RequestProfileListResponse)
	if res.NetPacketId() != openrgb.NetPacketIdRequestProfileList {
		t.Log("invalid net packet ID")
		t.Fail()
	}
	parser := p.DataParser()
	err := res.Decode(0, parser)
	if err != nil {
		t.Fatalf("decoding should not fail %s", err.Error())
	}
	if len(res.Names) != 2 {
		t.Fatalf("names should be 2 long but was %d", len(res.Names))
	}
	if res.Names[0] != "test" {
		t.Log("invalid data")
		t.Fail()
	}
	if res.Names[1] != "test1" {
		t.Log("invalid data")
		t.Fail()
	}
	err = parser.AssertAllRead()
	if err != nil {
		t.Log("data is not fully read")
		t.Fail()
	}
}

func TestRequestSaveProfileRequestEncode(t *testing.T) {
	builder := &openrgb.NetPacketDataBuilder{}
	req := &openrgb.RequestSaveProfileRequest{
		ProfileName: "test",
	}
	if req.NetPacketId() != openrgb.NetPacketIdRequestSaveProfile {
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

func TestRequestLoadProfileRequestEncode(t *testing.T) {
	builder := &openrgb.NetPacketDataBuilder{}
	req := &openrgb.RequestLoadProfileRequest{
		ProfileName: "test",
	}
	if req.NetPacketId() != openrgb.NetPacketIdRequestLoadProfile {
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

func TestRequestDeleteProfileRequestEncode(t *testing.T) {
	builder := &openrgb.NetPacketDataBuilder{}
	req := &openrgb.RequestDeleteProfileRequest{
		ProfileName: "test",
	}
	if req.NetPacketId() != openrgb.NetPacketIdRequestDeleteProfile {
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
