package openrgb

import (
	"io"
	"reflect"
)

// Encoder for OpenRGB net packets.
// Create a new encoder with NewNetPacketEncoder().
type NetPacketEncoder struct {
	w io.Writer
}

// Creates a new NetPacketEncoder.
func NewNetPacketEncoder(w io.Writer) *NetPacketEncoder {
	return &NetPacketEncoder{
		w: w,
	}
}

// Writes the header and the data to the writer.
func (e *NetPacketEncoder) Encode(p *NetPacket) error {
	err := e.writeHeader(p.Header)
	if err != nil {
		return err
	}
	return e.writeData(p.Data)
}

func (e *NetPacketEncoder) writeHeader(h *NetPacketHeader) error {
	buf := make([]byte, 16)
	copy(buf[0:4], netPacketHeaderMagicValue)
	protocolByteOrder.PutUint32(buf[4:8], h.PktDevIdx)
	protocolByteOrder.PutUint32(buf[8:12], uint32(h.PktId))
	protocolByteOrder.PutUint32(buf[12:16], uint32(h.PktSize))
	_, err := e.w.Write(buf)
	return err
}

func (e *NetPacketEncoder) writeData(data NetPacketData) error {
	_, err := e.w.Write(data)
	return err
}

// Encodable net packet data.
type DataEncoder interface {
	// Gets the body for the request with the given version.
	Encode(v Version, b *NetPacketDataBuilder)
}

// Net packet request.
type NetPacketRequest interface {
	NetPacketCommand
	DataEncoder
}

// Net packet data builder.
// The zero value for the builder is an empty data ready to use.
// All methods increase the size, if needed.
type NetPacketDataBuilder struct {
	content []byte
	offset  int
}

func (r *NetPacketDataBuilder) nextBytes(size int) []byte {
	r.EnsureSize(size)
	s := r.content[r.offset : r.offset+size]
	r.offset += size
	return s
}

// Makes sure that there are at least size bytes available.
func (r *NetPacketDataBuilder) EnsureSize(size int) {
	if r.content == nil {
		r.content = make([]byte, size)
	} else if len(r.content)-r.offset < size {
		r.content = append(r.content, make([]byte, size+r.offset-len(r.content))...)
	}
}

// Writes the data with a uint16.
func (r *NetPacketDataBuilder) WriteUint16(i uint16) {
	protocolByteOrder.PutUint16(r.nextBytes(2), i)
}

// Writes the data with a uint32.
func (r *NetPacketDataBuilder) WriteUint32(i uint32) {
	protocolByteOrder.PutUint32(r.nextBytes(4), i)
}

// Writes the data with a int32.
func (r *NetPacketDataBuilder) WriteInt32(i int32) {
	protocolByteOrder.PutUint32(r.nextBytes(4), uint32(i))
}

// Writes the data with a uint16 length.
// It panics if v's Kind is not [Array], [Chan], [Map], [Slice], [String], or pointer to [Array].
func (r *NetPacketDataBuilder) WriteLen(v any) {
	l := reflect.ValueOf(v).Len()
	r.WriteUint16(uint16(l))
}

// Writes the data with a slice of bytes.
func (r *NetPacketDataBuilder) WriteBytes(b []byte) {
	copy(r.nextBytes(len(b)), b)
}

// Writes the data with the length of a string followed by the string data.
func (r *NetPacketDataBuilder) WrtieString(b string) {
	r.WriteLen(b)
	copy(r.nextBytes(len(b)), b)
}

// Gets the data of the request.
func (r *NetPacketDataBuilder) Bytes() []byte {
	r.EnsureSize(0)
	return r.content
}
