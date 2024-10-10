package openrgb_test

import (
	"bytes"
	"testing"

	"github.com/csutorasa/go-openrgb-sdk"
)

func TestNetPacketEncoder(t *testing.T) {
	buf := bytes.NewBuffer([]byte{})
	e := openrgb.NewNetPacketEncoder(buf)
	p := openrgb.NewNetPacket(openrgb.NetPacketId(1000), 300, []byte(string("test")))
	err := e.Encode(p)
	if err != nil {
		t.Fatalf("ecoding should not fail %s", err.Error())
	}
	b := buf.Bytes()
	if len(b) != 20 {
		t.Fatalf("encoded data should be 20 long but was %d", len(b))
	}
	// 79 82 71 66 = "ORGB"
	if b[0] != 79 || b[1] != 82 || b[2] != 71 || b[3] != 66 {
		t.Log("invalid magic value")
		t.Fail()
	}
	// 44 * 1 + 1 * 256 = 300
	if b[4] != 44 || b[5] != 1 || b[6] != 0 || b[7] != 0 {
		t.Log("invalid device id")
		t.Fail()
	}
	// 232 * 1 + 3 * 256 = 1000
	if b[8] != 232 || b[9] != 3 || b[10] != 0 || b[11] != 0 {
		t.Log("invalid command id")
		t.Fail()
	}
	// 4 * 1 = 4
	if b[12] != 4 || b[13] != 0 || b[14] != 0 || b[15] != 0 {
		t.Log("invalid size")
		t.Fail()
	}
	// 116 101 115 116 = "test"
	if b[16] != 116 || b[17] != 101 || b[18] != 115 || b[19] != 116 {
		t.Log("invalid data")
		t.Fail()
	}
}

func TestNetPacketDataBuilderWriteUint16(t *testing.T) {
	builder := &openrgb.NetPacketDataBuilder{}
	builder.WriteUint16(300)
	b := builder.Bytes()
	if len(b) != 2 {
		t.Fatalf("encoded data should be 2 long but was %d", len(b))
	}
	// 44 * 1 + 1 * 256 = 300
	if b[0] != 44 || b[1] != 1 {
		t.Log("invalid data")
		t.Fail()
	}
}

func TestNetPacketDataBuilderWriteUint32(t *testing.T) {
	builder := &openrgb.NetPacketDataBuilder{}
	builder.WriteUint32(300)
	b := builder.Bytes()
	if len(b) != 4 {
		t.Fatalf("encoded data should be 4 long but was %d", len(b))
	}
	// 44 * 1 + 1 * 256 = 300
	if b[0] != 44 || b[1] != 1 || b[2] != 0 || b[3] != 0 {
		t.Log("invalid data")
		t.Fail()
	}
}

func TestNetPacketDataBuilderWriteInt32(t *testing.T) {
	builder := &openrgb.NetPacketDataBuilder{}
	builder.WriteInt32(300)
	b := builder.Bytes()
	if len(b) != 4 {
		t.Fatalf("encoded data should be 4 long but was %d", len(b))
	}
	// 44 * 1 + 1 * 256 = 300
	if b[0] != 44 || b[1] != 1 || b[2] != 0 || b[3] != 0 {
		t.Log("invalid data")
		t.Fail()
	}
}

func TestNetPacketDataBuilderWriteLen(t *testing.T) {
	builder := &openrgb.NetPacketDataBuilder{}
	builder.WriteLen("test")
	builder.WriteLen([]byte("test"))
	b := builder.Bytes()
	if len(b) != 4 {
		t.Fatalf("encoded data should be 4 long but was %d", len(b))
	}
	// 4 * 1 = 4
	if b[0] != 4 || b[1] != 0 {
		t.Log("invalid data")
		t.Fail()
	}
	// 4 * 1 = 4
	if b[2] != 4 || b[3] != 0 {
		t.Log("invalid data")
		t.Fail()
	}
}

func TestNetPacketDataBuilderWriteBytes(t *testing.T) {
	builder := &openrgb.NetPacketDataBuilder{}
	builder.WriteBytes([]byte("test"))
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

func TestNetPacketDataBuilderWriteString(t *testing.T) {
	builder := &openrgb.NetPacketDataBuilder{}
	builder.WrtieString("test")
	b := builder.Bytes()
	if len(b) != 6 {
		t.Fatalf("encoded data should be 6 long but was %d", len(b))
	}
	// 4 * 1 = 4
	if b[0] != 4 || b[1] != 0 {
		t.Log("invalid len")
		t.Fail()
	}
	// 116 101 115 116 = "test"
	if b[2] != 116 || b[3] != 101 || b[4] != 115 || b[5] != 116 {
		t.Log("invalid data")
		t.Fail()
	}
}
