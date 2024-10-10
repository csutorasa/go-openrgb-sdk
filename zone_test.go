package openrgb_test

import (
	"testing"

	"github.com/csutorasa/go-openrgb-sdk"
)

func TestZoneDecode(t *testing.T) {
	p := openrgb.NewNetPacket(0, 0, []byte{
		// 4
		4, 0,
		// "test"
		116, 101, 115, 116,
		// 5
		5, 0, 0, 0,
		// 1
		1, 0, 0, 0,
		// 100
		100, 0, 0, 0,
		// 50
		50, 0, 0, 0,
		// 0
		0, 0,
	})
	parser := p.DataParser()
	z := &openrgb.Zone{}
	err := z.Decode(0, parser)
	if err != nil {
		t.Fatalf("decoding should not fail %s", err.Error())
	}
	if z.ZoneName != "test" {
		t.Log("invalid data")
		t.Fail()
	}
	if z.ZoneType != 5 {
		t.Log("invalid data")
		t.Fail()
	}
	if z.ZoneLedsMin != 1 {
		t.Log("invalid data")
		t.Fail()
	}
	if z.ZoneLedsMax != 100 {
		t.Log("invalid data")
		t.Fail()
	}
	if z.ZoneLedsCount != 50 {
		t.Log("invalid data")
		t.Fail()
	}
	if z.ZoneMatrixData != nil {
		t.Log("invalid data")
		t.Fail()
	}
	err = parser.AssertAllRead()
	if err != nil {
		t.Log("data is not fully read")
		t.Fail()
	}
}

func TestZoneWithMatrixDecode(t *testing.T) {
	p := openrgb.NewNetPacket(0, 0, []byte{
		// 4
		4, 0,
		// "test"
		116, 101, 115, 116,
		// 5
		5, 0, 0, 0,
		// 1
		1, 0, 0, 0,
		// 100
		100, 0, 0, 0,
		// 50
		50, 0, 0, 0,
		// 6
		6, 0,
		// 2
		2, 0, 0, 0,
		// 3
		3, 0, 0, 0,
		// 10
		10, 0, 0, 0,
		// 11
		11, 0, 0, 0,
		// 12
		12, 0, 0, 0,
		// 13
		13, 0, 0, 0,
		// 14
		14, 0, 0, 0,
		// 15
		15, 0, 0, 0,
	})
	parser := p.DataParser()
	z := &openrgb.Zone{}
	err := z.Decode(0, parser)
	if err != nil {
		t.Fatalf("decoding should not fail %s", err.Error())
	}
	if z.ZoneName != "test" {
		t.Log("invalid data")
		t.Fail()
	}
	if z.ZoneType != 5 {
		t.Log("invalid data")
		t.Fail()
	}
	if z.ZoneLedsMin != 1 {
		t.Log("invalid data")
		t.Fail()
	}
	if z.ZoneLedsMax != 100 {
		t.Log("invalid data")
		t.Fail()
	}
	if z.ZoneLedsCount != 50 {
		t.Log("invalid data")
		t.Fail()
	}
	if z.ZoneMatrixData == nil {
		t.Log("invalid data")
		t.Fail()
	}
	if len(z.ZoneMatrixData) != 2 {
		t.Fatalf("matrix should have len 2 but was %d", len(z.ZoneMatrixData))
	}
	if len(z.ZoneMatrixData[0]) != 3 {
		t.Fatalf("matrix should have len 3 but was %d", len(z.ZoneMatrixData[0]))
	}
	if len(z.ZoneMatrixData[1]) != 3 {
		t.Fatalf("matrix should have len 3 but was %d", len(z.ZoneMatrixData[1]))
	}
	if z.ZoneMatrixData[0][0] != 10 || z.ZoneMatrixData[0][1] != 11 || z.ZoneMatrixData[0][2] != 12 || z.ZoneMatrixData[1][0] != 13 || z.ZoneMatrixData[1][1] != 14 || z.ZoneMatrixData[1][2] != 15 {
		t.Log("invalid data")
		t.Fail()
	}
	err = parser.AssertAllRead()
	if err != nil {
		t.Log("data is not fully read")
		t.Fail()
	}
}
