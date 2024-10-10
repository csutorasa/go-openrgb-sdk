package openrgb

type Led struct {
	// LED name string value
	LedName string
	// LED value field value
	LedValue uint32
}

// Decode implements DataDecoder.
func (l *Led) Decode(v Version, p *NetPacketDataParser) error {
	ledName, err := p.ReadString()
	if err != nil {
		return err
	}
	ledValue, err := p.ReadUint32()
	if err != nil {
		return err
	}
	l.LedName = ledName
	l.LedValue = ledValue
	return nil
}
