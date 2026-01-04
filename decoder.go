package openrgb

import (
	"io"
)

// Decoder for OpenTGB net packets.
// Create a new decoder with NewNetPacketDecoder().
type NetPacketDecoder struct {
	r io.Reader
}

// Creates a new NetPacketDecoder.
func NewNetPacketDecoder(r io.Reader) *NetPacketDecoder {
	return &NetPacketDecoder{
		r: r,
	}
}

// Reads the header and the data from the reader.
func (d *NetPacketDecoder) Decode() (*NetPacket, error) {
	header, err := d.decodeHeader()
	if err != nil {
		return nil, err
	}
	data, err := d.decodeData(header)
	if err != nil {
		return nil, err
	}
	return NewNetPacket(header.PktId, header.PktDevIdx, data), nil
}

func (d *NetPacketDecoder) decodeHeader() (*NetPacketHeader, error) {
	buf := make([]byte, 16)
	_, err := io.ReadFull(d.r, buf)
	if err != nil {
		return nil, err
	}
	magic := string(buf[0:4])
	if magic != netPacketHeaderMagicValue {
		return nil, &InvalidPacketHeaderMagicValueError{v: magic}
	}
	deviceIndex := protocolByteOrder.Uint32(buf[4:8])
	packetId := protocolByteOrder.Uint32(buf[8:12])
	packetSize := protocolByteOrder.Uint32(buf[12:16])
	header := &NetPacketHeader{
		PktMagic:  magic,
		PktDevIdx: deviceIndex,
		PktId:     NetPacketId(packetId),
		PktSize:   packetSize,
	}
	return header, nil
}

func (d *NetPacketDecoder) decodeData(header *NetPacketHeader) (NetPacketData, error) {
	buf := make([]byte, header.PktSize)
	_, err := io.ReadFull(d.r, buf)
	if err != nil {
		return nil, err
	}
	return buf, nil
}

// Decodable net packet data.
type DataDecoder interface {
	// Reads the data from the NetPacketDataParser and fills the fields of this object.
	// Field values are changed only when the all reads were successful.
	// Errors can still happen after a successful decode.
	Decode(v Version, p *NetPacketDataParser) error
}

// Net packet response.
type NetPacketResponse interface {
	NetPacketCommand
	DataDecoder
}

// Creates a parser for the data.
func (p *NetPacket) DataParser() *NetPacketDataParser {
	return &NetPacketDataParser{
		content: p.Data,
		offset:  0,
	}
}

// Net packet data parser.
// Create a new response with NetPacket.DataParser()
type NetPacketDataParser struct {
	content []byte
	offset  int
}

func (r *NetPacketDataParser) nextBytes(size int) ([]byte, error) {
	if len(r.content)-r.offset < size {
		return nil, io.EOF
	}
	s := r.content[r.offset : r.offset+size]
	r.offset += size
	return s, nil
}

// Returns an error if there is unread data.
func (p *NetPacketDataParser) AssertAllRead() error {
	if len(p.content) != p.offset {
		return &UnprocessedResponseError{total: len(p.content), read: p.offset}
	}
	return nil
}

// Reads a uint16 from the content.
// Returns io.EOF error if the content runs out of data.
func (p *NetPacketDataParser) ReadUint16() (uint16, error) {
	b, err := p.nextBytes(2)
	if err != nil {
		return 0, err
	}
	return protocolByteOrder.Uint16(b), nil
}

// Reads a uint32 from the content.
// Returns io.EOF error if the content runs out of data.
func (p *NetPacketDataParser) ReadUint32() (uint32, error) {
	b, err := p.nextBytes(4)
	if err != nil {
		return 0, err
	}
	return protocolByteOrder.Uint32(b), nil
}

// Reads an int32 from the content.
// Returns io.EOF error if the content runs out of data.
func (p *NetPacketDataParser) ReadInt32() (int32, error) {
	i, err := p.ReadUint32()
	if err != nil {
		return 0, err
	}
	return int32(i), nil
}

// Reads the given number of bytes from the content.
// Returns io.EOF error if the content runs out of data.
func (p *NetPacketDataParser) ReadBytes(l int) ([]byte, error) {
	b, err := p.nextBytes(int(l))
	if err != nil {
		return nil, err
	}
	return b, nil
}

// Reads the length of the string, then reads the string itself from the content.
// Returns io.EOF error if the content runs out of data.
func (p *NetPacketDataParser) ReadString() (string, error) {
	l, err := p.ReadUint16()
	if err != nil {
		return "", err
	}
	b, err := p.nextBytes(int(l))
	if err != nil {
		return "", err
	}
	return string(b), nil
}
