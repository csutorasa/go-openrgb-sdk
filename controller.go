package openrgb

import "sync"

type ControllerData struct {
	// RGBController type field value
	Type int32
	// RGBController name field string value
	Name string
	// RGBController vendor field string value
	Vendor string
	// RGBController description field string value
	Description string
	// RGBController version field string value
	Version string
	// RGBController serial field string value
	Serial string
	// RGBController location field string value
	Location string
	// RGBController active_mode field value
	ActiveMode int32
	Modes      Modes
	Zones      Zones
	Leds       []*Led
	Colors     []Color
}

// Decode implements DataDecoder.
func (c *ControllerData) Decode(v Version, p *NetPacketDataParser) error {
	_, err := p.ReadUint32()
	if err != nil {
		return err
	}
	t, err := p.ReadInt32()
	if err != nil {
		return err
	}
	name, err := p.ReadString()
	if err != nil {
		return err
	}
	var vendor string
	if v >= 1 {
		vendor, err = p.ReadString()
		if err != nil {
			return err
		}
	}
	description, err := p.ReadString()
	if err != nil {
		return err
	}
	version, err := p.ReadString()
	if err != nil {
		return err
	}
	serial, err := p.ReadString()
	if err != nil {
		return err
	}
	location, err := p.ReadString()
	if err != nil {
		return err
	}
	numModes, err := p.ReadUint16()
	if err != nil {
		return err
	}
	activeMode, err := p.ReadInt32()
	if err != nil {
		return err
	}
	modes := make([]*Mode, numModes)
	for i := uint16(0); i < numModes; i++ {
		modes[i] = &Mode{}
		err := modes[i].Decode(v, p)
		if err != nil {
			return err
		}
	}
	numZones, err := p.ReadUint16()
	if err != nil {
		return err
	}
	zones := make([]*Zone, numZones)
	for i := uint16(0); i < numZones; i++ {
		zones[i] = &Zone{}
		err := zones[i].Decode(v, p)
		if err != nil {
			return err
		}
	}
	numLeds, err := p.ReadUint16()
	if err != nil {
		return err
	}
	leds := make([]*Led, numLeds)
	for i := uint16(0); i < numLeds; i++ {
		leds[i] = &Led{}
		err := leds[i].Decode(v, p)
		if err != nil {
			return err
		}
	}
	numColors, err := p.ReadUint16()
	if err != nil {
		return err
	}
	colors := make([]Color, numColors)
	for i := uint16(0); i < numColors; i++ {
		err := colors[i].Decode(v, p)
		if err != nil {
			return err
		}
	}
	c.Type = t
	c.Name = name
	c.Vendor = vendor
	c.Description = description
	c.Version = version
	c.Serial = serial
	c.Location = location
	c.ActiveMode = activeMode
	c.Modes = modes
	c.Zones = zones
	c.Leds = leds
	c.Colors = colors
	return nil
}

type Controllers []*ControllerData

// Finds a controller by a condition.
func (c Controllers) Find(f func(*ControllerData) bool) (uint32, *ControllerData) {
	for i, controller := range c {
		if f(controller) {
			return uint32(i), controller
		}
	}
	return 0, nil
}

// Finds a controller by name.
func (c Controllers) FindByName(n string) (uint32, *ControllerData) {
	return c.Find(func(d *ControllerData) bool {
		return d.Name == n
	})
}

// Cache to use for controller data.
// Create a new cache with NewControllerCache().
type ControllerCache struct {
	l           sync.Mutex
	client      *Client
	controllers []*ControllerData
}

// Creates a new ControllerCache.
func NewControllerCache(c *Client) *ControllerCache {
	return &ControllerCache{
		client:      c,
		controllers: nil,
	}
}

// Gets the controllers.
func (c *ControllerCache) Controllers() (Controllers, error) {
	c.l.Lock()
	defer c.l.Unlock()
	if c.controllers == nil {
		err := c.requestControllers()
		if err != nil {
			return nil, err
		}
	}
	return c.controllers, nil
}

// Invalidates the current state.
// This should be called when a DeviceListUpdated packet is received.
func (c *ControllerCache) Invalidate() {
	c.l.Lock()
	defer c.l.Unlock()
	c.controllers = nil
}

func (c *ControllerCache) requestControllers() error {
	l, err := c.client.RequestControllerCount()
	if err != nil {
		return err
	}
	c.controllers = make([]*ControllerData, l.Count)
	for i := uint32(0); i < l.Count; i++ {
		controller, err := c.client.RequestControllerData(i)
		if err != nil {
			return err
		}
		c.controllers[i] = controller.Controller
	}
	return nil
}
