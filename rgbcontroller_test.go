package openrgb_test

import (
	"testing"

	"github.com/csutorasa/go-openrgb-sdk"
)

func TestRGBControllerResizeZoneRequestEncode(t *testing.T) {
	builder := &openrgb.NetPacketDataBuilder{}
	req := &openrgb.RGBControllerResizeZoneRequest{
		ZoneIdx: 1,
		NewSize: 10,
	}
	if req.NetPacketId() != openrgb.NetPacketIdRgbcontrollerResizezone {
		t.Log("invalid net packet ID")
		t.Fail()
	}
	req.Encode(0, builder)
	b := builder.Bytes()
	if len(b) != 8 {
		t.Fatalf("encoded data should be 8 long but was %d", len(b))
	}
	// 1
	if b[0] != 1 || b[1] != 0 || b[2] != 0 || b[3] != 0 {
		t.Log("invalid data")
		t.Fail()
	}
	// 1
	if b[4] != 10 || b[5] != 0 || b[6] != 0 || b[7] != 0 {
		t.Log("invalid data")
		t.Fail()
	}
}

func TestRGBControllerUpdateSingleLedRequestEncode(t *testing.T) {
	builder := &openrgb.NetPacketDataBuilder{}
	req := &openrgb.RGBControllerUpdateSingleLedRequest{
		LedIdx: 1,
		Color: openrgb.Color{
			R: 255,
			G: 240,
			B: 128,
		},
	}
	if req.NetPacketId() != openrgb.NetPacketIdRgbcontrollerUpdatesingleled {
		t.Log("invalid net packet ID")
		t.Fail()
	}
	req.Encode(0, builder)
	b := builder.Bytes()
	if len(b) != 8 {
		t.Fatalf("encoded data should be 8 long but was %d", len(b))
	}
	// 1
	if b[0] != 1 || b[1] != 0 || b[2] != 0 || b[3] != 0 {
		t.Log("invalid data")
		t.Fail()
	}
	// 255,240,128
	if b[4] != 255 || b[5] != 240 || b[6] != 128 || b[7] != 0 {
		t.Log("invalid data")
		t.Fail()
	}
}

func TestRGBControllerUpdateLedsRequestEncode(t *testing.T) {
	builder := &openrgb.NetPacketDataBuilder{}
	req := &openrgb.RGBControllerUpdateLedsRequest{
		LedColor: []openrgb.Color{
			{
				R: 255,
				G: 240,
				B: 128,
			},
			{
				R: 128,
				G: 255,
				B: 240,
			},
		},
	}
	if req.NetPacketId() != openrgb.NetPacketIdRgbcontrollerUpdateleds {
		t.Log("invalid net packet ID")
		t.Fail()
	}
	req.Encode(0, builder)
	b := builder.Bytes()
	if len(b) != 14 {
		t.Fatalf("encoded data should be 14 long but was %d", len(b))
	}
	// 14
	if b[0] != 14 || b[1] != 0 || b[2] != 0 || b[3] != 0 {
		t.Log("invalid data")
		t.Fail()
	}
	// 2
	if b[4] != 2 || b[5] != 0 {
		t.Log("invalid data")
		t.Fail()
	}
	// 255,240,128
	if b[6] != 255 || b[7] != 240 || b[8] != 128 || b[9] != 0 {
		t.Log("invalid data")
		t.Fail()
	}
	// 128,255,240
	if b[10] != 128 || b[11] != 255 || b[12] != 240 || b[13] != 0 {
		t.Log("invalid data")
		t.Fail()
	}
}

func TestRGBControllerUpdateZoneLedsRequestEncode(t *testing.T) {
	builder := &openrgb.NetPacketDataBuilder{}
	req := &openrgb.RGBControllerUpdateZoneLedsRequest{
		ZoneIdx: 1,
		LedColor: []openrgb.Color{
			{
				R: 255,
				G: 240,
				B: 128,
			},
			{
				R: 128,
				G: 255,
				B: 240,
			},
		},
	}
	if req.NetPacketId() != openrgb.NetPacketIdRgbcontrollerUpdatezoneleds {
		t.Log("invalid net packet ID")
		t.Fail()
	}
	req.Encode(0, builder)
	b := builder.Bytes()
	if len(b) != 18 {
		t.Fatalf("encoded data should be 18 long but was %d", len(b))
	}
	// 18
	if b[0] != 18 || b[1] != 0 || b[2] != 0 || b[3] != 0 {
		t.Log("invalid data")
		t.Fail()
	}
	// 1
	if b[4] != 1 || b[5] != 0 || b[6] != 0 || b[7] != 0 {
		t.Log("invalid data")
		t.Fail()
	}
	// 2
	if b[8] != 2 || b[9] != 0 {
		t.Log("invalid data")
		t.Fail()
	}
	// 255,240,128
	if b[10] != 255 || b[11] != 240 || b[12] != 128 || b[13] != 0 {
		t.Log("invalid data")
		t.Fail()
	}
	// 128,255,240
	if b[14] != 128 || b[15] != 255 || b[16] != 240 || b[17] != 0 {
		t.Log("invalid data")
		t.Fail()
	}
}

func TestRGBControllerSetCustomModeRequestEncode(t *testing.T) {
	builder := &openrgb.NetPacketDataBuilder{}
	req := &openrgb.RGBControllerSetCustomModeRequest{}
	if req.NetPacketId() != openrgb.NetPacketIdRgbcontrollerSetcustommode {
		t.Log("invalid net packet ID")
		t.Fail()
	}
	req.Encode(0, builder)
	b := builder.Bytes()
	if len(b) != 0 {
		t.Fatalf("encoded data should be 0 long but was %d", len(b))
	}
}

func TestRGBControllerUpdateModeRequestEncode(t *testing.T) {
	builder := &openrgb.NetPacketDataBuilder{}
	req := &openrgb.RGBControllerUpdateModeRequest{
		ModeIdx: 1,
		Mode: &openrgb.Mode{
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
		},
	}
	if req.NetPacketId() != openrgb.NetPacketIdRgbcontrollerUpdatemode {
		t.Log("invalid net packet ID")
		t.Fail()
	}
	req.Encode(0, builder)
	b := builder.Bytes()
	if len(b) != 56 {
		t.Fatalf("encoded data should be 56 long but was %d", len(b))
	}
	// 18
	if b[0] != 56 || b[1] != 0 || b[2] != 0 || b[3] != 0 {
		t.Log("invalid data")
		t.Fail()
	}
	// 1
	if b[4] != 1 || b[5] != 0 || b[6] != 0 || b[7] != 0 {
		t.Log("invalid data")
		t.Fail()
	}
}

func TestRGBControllerSaveModeRequestEncode(t *testing.T) {
	builder := &openrgb.NetPacketDataBuilder{}
	req := &openrgb.RGBControllerSaveModeRequest{
		ModeIdx: 1,
		Mode: &openrgb.Mode{
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
		},
	}
	if req.NetPacketId() != openrgb.NetPacketIdRgbcontrollerSavemode {
		t.Log("invalid net packet ID")
		t.Fail()
	}
	req.Encode(0, builder)
	b := builder.Bytes()
	if len(b) != 56 {
		t.Fatalf("encoded data should be 56 long but was %d", len(b))
	}
	// 18
	if b[0] != 56 || b[1] != 0 || b[2] != 0 || b[3] != 0 {
		t.Log("invalid data")
		t.Fail()
	}
	// 1
	if b[4] != 1 || b[5] != 0 || b[6] != 0 || b[7] != 0 {
		t.Log("invalid data")
		t.Fail()
	}
}
