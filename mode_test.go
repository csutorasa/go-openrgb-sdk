package openrgb_test

import (
	"testing"

	"github.com/csutorasa/go-openrgb-sdk"
)

func TestModeDecode(t *testing.T) {
	p := openrgb.NewNetPacket(0, 0, []byte{
		// 4
		4, 0,
		// test
		116, 101, 115, 116,
		// 2
		2, 0, 0, 0,
		// 3
		3, 0, 0, 0,
		// 10
		10, 0, 0, 0,
		// 20
		20, 0, 0, 0,
		// 30
		50, 0, 0, 0,
		// 40
		60, 0, 0, 0,
		// 15
		15, 0, 0, 0,
		// 1
		1, 0, 0, 0,
		// 55
		55, 0, 0, 0,
		// 1
		1, 0,
		// 255,240,128
		255, 240, 128, 0,
	})
	parser := p.DataParser()
	m := &openrgb.Mode{}
	err := m.Decode(0, parser)
	if err != nil {
		t.Fatalf("decoding should not fail %s", err.Error())
	}
	if m.ModeName != "test" {
		t.Log("invalid data")
		t.Fail()
	}
	if m.ModeValue != 2 {
		t.Log("invalid data")
		t.Fail()
	}
	if m.ModeFlags != 3 {
		t.Log("invalid data")
		t.Fail()
	}
	if m.ModeSpeedMin != 10 {
		t.Log("invalid data")
		t.Fail()
	}
	if m.ModeSpeedMax != 20 {
		t.Log("invalid data")
		t.Fail()
	}
	if m.ModeBrightnessMin != 0 {
		t.Log("invalid data")
		t.Fail()
	}
	if m.ModeBrightnessMax != 0 {
		t.Log("invalid data")
		t.Fail()
	}
	if m.ModeColorsMin != 50 {
		t.Log("invalid data")
		t.Fail()
	}
	if m.ModeColorsMax != 60 {
		t.Log("invalid data")
		t.Fail()
	}
	if m.ModeSpeed != 15 {
		t.Log("invalid data")
		t.Fail()
	}
	if m.ModeBrightness != 0 {
		t.Log("invalid data")
		t.Fail()
	}
	if m.ModeDirection != 1 {
		t.Log("invalid data")
		t.Fail()
	}
	if m.ModeColorMode != 55 {
		t.Log("invalid data")
		t.Fail()
	}
	if len(m.ModeColors) != 1 {
		t.Fatalf("colors be 1 long but was %d", len(m.ModeColors))
	}
	if m.ModeColors[0].R != 255 || m.ModeColors[0].G != 240 || m.ModeColors[0].B != 128 {
		t.Log("invalid data")
		t.Fail()
	}
	err = parser.AssertAllRead()
	if err != nil {
		t.Log("data is not fully read")
		t.Fail()
	}
	p = openrgb.NewNetPacket(0, 0, []byte{
		// 4
		4, 0,
		// test
		116, 101, 115, 116,
		// 2
		2, 0, 0, 0,
		// 3
		3, 0, 0, 0,
		// 10
		10, 0, 0, 0,
		// 20
		20, 0, 0, 0,
		// 30
		30, 0, 0, 0,
		// 40
		40, 0, 0, 0,
		// 50
		50, 0, 0, 0,
		// 60
		60, 0, 0, 0,
		// 15
		15, 0, 0, 0,
		// 35
		35, 0, 0, 0,
		// 1
		1, 0, 0, 0,
		// 55
		55, 0, 0, 0,
		// 1
		1, 0,
		// 255,240,128
		255, 240, 128, 0,
	})
	parser = p.DataParser()
	err = m.Decode(3, parser)
	if err != nil {
		t.Fatalf("decoding should not fail %s", err.Error())
	}
	if m.ModeName != "test" {
		t.Log("invalid data")
		t.Fail()
	}
	if m.ModeValue != 2 {
		t.Log("invalid data")
		t.Fail()
	}
	if m.ModeFlags != 3 {
		t.Log("invalid data")
		t.Fail()
	}
	if m.ModeSpeedMin != 10 {
		t.Log("invalid data")
		t.Fail()
	}
	if m.ModeSpeedMax != 20 {
		t.Log("invalid data")
		t.Fail()
	}
	if m.ModeBrightnessMin != 30 {
		t.Log("invalid data")
		t.Fail()
	}
	if m.ModeBrightnessMax != 40 {
		t.Log("invalid data")
		t.Fail()
	}
	if m.ModeColorsMin != 50 {
		t.Log("invalid data")
		t.Fail()
	}
	if m.ModeColorsMax != 60 {
		t.Log("invalid data")
		t.Fail()
	}
	if m.ModeSpeed != 15 {
		t.Log("invalid data")
		t.Fail()
	}
	if m.ModeBrightness != 35 {
		t.Log("invalid data")
		t.Fail()
	}
	if m.ModeDirection != 1 {
		t.Log("invalid data")
		t.Fail()
	}
	if m.ModeColorMode != 55 {
		t.Log("invalid data")
		t.Fail()
	}
	if len(m.ModeColors) != 1 {
		t.Fatalf("colors be 1 long but was %d", len(m.ModeColors))
	}
	if m.ModeColors[0].R != 255 || m.ModeColors[0].G != 240 || m.ModeColors[0].B != 128 {
		t.Log("invalid data")
		t.Fail()
	}
	err = parser.AssertAllRead()
	if err != nil {
		t.Log("data is not fully read")
		t.Fail()
	}
}

func TestModeEncode(t *testing.T) {
	builder := &openrgb.NetPacketDataBuilder{}
	m := &openrgb.Mode{
		ModeName:          "test",
		ModeValue:         2,
		ModeFlags:         3,
		ModeSpeedMin:      10,
		ModeSpeedMax:      20,
		ModeBrightnessMin: 30,
		ModeBrightnessMax: 40,
		ModeColorsMin:     50,
		ModeColorsMax:     60,
		ModeSpeed:         15,
		ModeBrightness:    35,
		ModeDirection:     1,
		ModeColorMode:     55,
		ModeColors: []openrgb.Color{
			{R: 255, G: 240, B: 128},
		},
	}
	m.Encode(0, builder)
	b := builder.Bytes()
	if len(b) != 48 {
		t.Fatalf("encoded data should be 48 long but was %d", len(b))
	}
	// 4
	if b[0] != 4 || b[1] != 0 {
		t.Log("invalid data")
		t.Fail()
	}
	// 116 101 115 116 = "test"
	if b[2] != 116 || b[3] != 101 || b[4] != 115 || b[5] != 116 {
		t.Log("invalid data")
		t.Fail()
	}
	// 2
	if b[6] != 2 || b[7] != 0 || b[8] != 0 || b[9] != 0 {
		t.Log("invalid data")
		t.Fail()
	}
	// 3
	if b[10] != 3 || b[11] != 0 || b[12] != 0 || b[13] != 0 {
		t.Log("invalid data")
		t.Fail()
	}
	// 10
	if b[14] != 10 || b[15] != 0 || b[16] != 0 || b[17] != 0 {
		t.Log("invalid data")
		t.Fail()
	}
	// 20
	if b[18] != 20 || b[19] != 0 || b[20] != 0 || b[21] != 0 {
		t.Log("invalid data")
		t.Fail()
	}
	// 50
	if b[22] != 50 || b[23] != 0 || b[24] != 0 || b[25] != 0 {
		t.Log("invalid data")
		t.Fail()
	}
	// 60
	if b[26] != 60 || b[27] != 0 || b[28] != 0 || b[29] != 0 {
		t.Log("invalid data")
		t.Fail()
	}
	// 15
	if b[30] != 15 || b[31] != 0 || b[32] != 0 || b[33] != 0 {
		t.Log("invalid data")
		t.Fail()
	}
	// 1
	if b[34] != 1 || b[35] != 0 || b[36] != 0 || b[37] != 0 {
		t.Log("invalid data")
		t.Fail()
	}
	// 55
	if b[38] != 55 || b[39] != 0 || b[40] != 0 || b[41] != 0 {
		t.Log("invalid data")
		t.Fail()
	}
	// 1
	if b[42] != 1 || b[43] != 0 {
		t.Log("invalid data")
		t.Fail()
	}
	// 255,240,128
	if b[44] != 255 || b[45] != 240 || b[46] != 128 || b[47] != 0 {
		t.Log("invalid data")
		t.Fail()
	}
	builder = &openrgb.NetPacketDataBuilder{}
	m.Encode(3, builder)
	b = builder.Bytes()
	if len(b) != 60 {
		t.Fatalf("encoded data should be 60 long but was %d", len(b))
	}
	// 4
	if b[0] != 4 || b[1] != 0 {
		t.Log("invalid data")
		t.Fail()
	}
	// 116 101 115 116 = "test"
	if b[2] != 116 || b[3] != 101 || b[4] != 115 || b[5] != 116 {
		t.Log("invalid data")
		t.Fail()
	}
	// 2
	if b[6] != 2 || b[7] != 0 || b[8] != 0 || b[9] != 0 {
		t.Log("invalid data")
		t.Fail()
	}
	// 3
	if b[10] != 3 || b[11] != 0 || b[12] != 0 || b[13] != 0 {
		t.Log("invalid data")
		t.Fail()
	}
	// 10
	if b[14] != 10 || b[15] != 0 || b[16] != 0 || b[17] != 0 {
		t.Log("invalid data")
		t.Fail()
	}
	// 20
	if b[18] != 20 || b[19] != 0 || b[20] != 0 || b[21] != 0 {
		t.Log("invalid data")
		t.Fail()
	}
	// 30
	if b[22] != 30 || b[23] != 0 || b[24] != 0 || b[25] != 0 {
		t.Log("invalid data")
		t.Fail()
	}
	// 40
	if b[26] != 40 || b[27] != 0 || b[28] != 0 || b[29] != 0 {
		t.Log("invalid data")
		t.Fail()
	}
	// 50
	if b[30] != 50 || b[31] != 0 || b[32] != 0 || b[33] != 0 {
		t.Log("invalid data")
		t.Fail()
	}
	// 60
	if b[34] != 60 || b[35] != 0 || b[36] != 0 || b[37] != 0 {
		t.Log("invalid data")
		t.Fail()
	}
	// 15
	if b[38] != 15 || b[39] != 0 || b[40] != 0 || b[41] != 0 {
		t.Log("invalid data")
		t.Fail()
	}
	// 35
	if b[42] != 35 || b[43] != 0 || b[44] != 0 || b[45] != 0 {
		t.Log("invalid data")
		t.Fail()
	}
	// 1
	if b[46] != 1 || b[47] != 0 || b[48] != 0 || b[49] != 0 {
		t.Log("invalid data")
		t.Fail()
	}
	// 55
	if b[50] != 55 || b[51] != 0 || b[52] != 0 || b[53] != 0 {
		t.Log("invalid data")
		t.Fail()
	}
	// 1
	if b[54] != 1 || b[55] != 0 {
		t.Log("invalid data")
		t.Fail()
	}
	// 255,240,128
	if b[56] != 255 || b[57] != 240 || b[58] != 128 || b[59] != 0 {
		t.Log("invalid data")
		t.Fail()
	}
}
