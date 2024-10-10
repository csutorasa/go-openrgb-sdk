package openrgb

type Zone struct {
	// Zone name string value
	ZoneName string
	// Zone type value
	ZoneType int32
	// Zone leds_min value
	ZoneLedsMin uint32
	// Zone leds_max value
	ZoneLedsMax uint32
	// Zone leds_count value
	ZoneLedsCount uint32
	// Zone matrix_map data (*only if matrix_map exists)
	ZoneMatrixData [][]uint32
}

// Decode implements DataDecoder.
func (z *Zone) Decode(v Version, p *NetPacketDataParser) error {
	zoneName, err := p.ReadString()
	if err != nil {
		return err
	}
	zoneType, err := p.ReadInt32()
	if err != nil {
		return err
	}
	zoneLedsMin, err := p.ReadUint32()
	if err != nil {
		return err
	}
	zoneLedsMax, err := p.ReadUint32()
	if err != nil {
		return err
	}
	zoneLedsCount, err := p.ReadUint32()
	if err != nil {
		return err
	}
	zoneMatrixLen, err := p.ReadUint16()
	if err != nil {
		return err
	}
	var zoneMatrixData [][]uint32
	if zoneMatrixLen > 0 {
		zoneMatrixHeight, err := p.ReadUint32()
		if err != nil {
			return err
		}
		zoneMatrixWidth, err := p.ReadUint32()
		if err != nil {
			return err
		}
		zoneMatrixData = make([][]uint32, zoneMatrixHeight)
		for i := uint32(0); i < zoneMatrixHeight; i++ {
			zoneMatrixData[i] = make([]uint32, zoneMatrixWidth)
			for j := uint32(0); j < zoneMatrixWidth; j++ {
				matrixValue, err := p.ReadUint32()
				if err != nil {
					return err
				}
				zoneMatrixData[i][j] = matrixValue
			}
		}
	}
	z.ZoneName = zoneName
	z.ZoneType = zoneType
	z.ZoneLedsMin = zoneLedsMin
	z.ZoneLedsMax = zoneLedsMax
	z.ZoneLedsCount = zoneLedsCount
	z.ZoneMatrixData = zoneMatrixData
	return nil
}

type Zones []*Zone

func (z Zones) FindByName(n string) (uint32, *Zone) {
	for i, zone := range z {
		if zone.ZoneName == n {
			return uint32(i), zone
		}
	}
	return 0, nil
}
