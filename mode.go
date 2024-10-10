package openrgb

type Mode struct {
	// Mode name string value
	ModeName string
	// Mode value field value
	ModeValue int32
	// Mode flags field value
	ModeFlags uint32
	// Mode speed_min field value
	ModeSpeedMin uint32
	// Mode speed_max field value
	ModeSpeedMax uint32
	// Mode brightness_min field value
	ModeBrightnessMin uint32
	//Mode brightness_max field value
	ModeBrightnessMax uint32
	// Mode colors_min field value
	ModeColorsMin uint32
	// Mode colors_max field value
	ModeColorsMax uint32
	// Mode speed value
	ModeSpeed uint32
	// Mode brightness value
	ModeBrightness uint32
	// Mode direction value
	ModeDirection uint32
	// Mode color_mode value
	ModeColorMode uint32
	// Mode color values
	ModeColors []Color
}

func (m *Mode) Size(v Version) int {
	s := 40 + len(m.ModeName) + 4*len(m.ModeColors)
	if v >= 3 {
		s += 12
	}
	return s
}

// Encode implements DataEncoder.
func (m *Mode) Encode(v Version, b *NetPacketDataBuilder) {
	b.EnsureSize(m.Size(v))
	b.WrtieString(m.ModeName)
	b.WriteInt32(m.ModeValue)
	b.WriteUint32(m.ModeFlags)
	b.WriteUint32(m.ModeSpeedMin)
	b.WriteUint32(m.ModeSpeedMax)
	if v >= 3 {
		b.WriteUint32(m.ModeBrightnessMin)
		b.WriteUint32(m.ModeBrightnessMax)
	}
	b.WriteUint32(m.ModeColorsMin)
	b.WriteUint32(m.ModeColorsMax)
	b.WriteUint32(m.ModeSpeed)
	if v >= 3 {
		b.WriteUint32(m.ModeBrightness)
	}
	b.WriteUint32(m.ModeDirection)
	b.WriteUint32(m.ModeColorMode)
	b.WriteLen(m.ModeColors)
	for _, c := range m.ModeColors {
		c.Encode(v, b)
	}
}

// Decode implements DataDecoder.
func (m *Mode) Decode(v Version, p *NetPacketDataParser) error {
	modeName, err := p.ReadString()
	if err != nil {
		return err
	}
	modeValue, err := p.ReadInt32()
	if err != nil {
		return err
	}
	modeFlags, err := p.ReadUint32()
	if err != nil {
		return err
	}
	modeSpeedMin, err := p.ReadUint32()
	if err != nil {
		return err
	}
	modeSpeedMax, err := p.ReadUint32()
	if err != nil {
		return err
	}
	var modeBrightnessMin uint32
	if v >= 3 {
		modeBrightnessMin, err = p.ReadUint32()
		if err != nil {
			return err
		}
	}
	var modeBrightnessMax uint32
	if v >= 3 {
		modeBrightnessMax, err = p.ReadUint32()
		if err != nil {
			return err
		}
	}
	modeColorsMin, err := p.ReadUint32()
	if err != nil {
		return err
	}
	modeColorsMax, err := p.ReadUint32()
	if err != nil {
		return err
	}
	modeSpeed, err := p.ReadUint32()
	if err != nil {
		return err
	}
	var modeBrightness uint32
	if v >= 3 {
		modeBrightness, err = p.ReadUint32()
		if err != nil {
			return err
		}
	}
	modeDirection, err := p.ReadUint32()
	if err != nil {
		return err
	}
	modeColorMode, err := p.ReadUint32()
	if err != nil {
		return err
	}
	numModeColors, err := p.ReadUint16()
	if err != nil {
		return err
	}
	modeColors := make([]Color, numModeColors)
	for i := uint16(0); i < numModeColors; i++ {
		err := modeColors[i].Decode(v, p)
		if err != nil {
			return err
		}
	}
	m.ModeName = modeName
	m.ModeValue = modeValue
	m.ModeFlags = modeFlags
	m.ModeSpeedMin = modeSpeedMin
	m.ModeSpeedMax = modeSpeedMax
	m.ModeBrightnessMin = modeBrightnessMin
	m.ModeBrightnessMax = modeBrightnessMax
	m.ModeColorsMin = modeColorsMin
	m.ModeColorsMax = modeColorsMax
	m.ModeSpeed = modeSpeed
	m.ModeBrightness = modeBrightness
	m.ModeDirection = modeDirection
	m.ModeColorMode = modeColorMode
	m.ModeColors = modeColors
	return nil
}

type Modes []*Mode

func (m Modes) FindByName(n string) (uint32, *Mode) {
	for i, mode := range m {
		if mode.ModeName == n {
			return uint32(i), mode
		}
	}
	return 0, nil
}
