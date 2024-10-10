package openrgb

type RGBControllerResizeZoneRequest struct {
	ZoneIdx int32
	NewSize int32
}

// NetPacketId implements NetPacketCommand.
func (req *RGBControllerResizeZoneRequest) NetPacketId() NetPacketId {
	return NetPacketIdRgbcontrollerResizezone
}

// NetPacketId implements DataEncoder.
func (req *RGBControllerResizeZoneRequest) Encode(v Version, b *NetPacketDataBuilder) {
	b.EnsureSize(8)
	b.WriteInt32(req.ZoneIdx)
	b.WriteInt32(req.NewSize)
}

// Resizes a zone.
func (c *Client) RGBControllerResizeZone(deviceId uint32, req *RGBControllerResizeZoneRequest) error {
	return c.SendRequestForDevice(deviceId, req)
}

type RGBControllerUpdateSingleLedRequest struct {
	LedIdx int32
	Color  Color
}

// NetPacketId implements NetPacketCommand.
func (req *RGBControllerUpdateSingleLedRequest) NetPacketId() NetPacketId {
	return NetPacketIdRgbcontrollerUpdatesingleled
}

// NetPacketId implements DataEncoder.
func (req *RGBControllerUpdateSingleLedRequest) Encode(v Version, b *NetPacketDataBuilder) {
	b.EnsureSize(8)
	b.WriteInt32(req.LedIdx)
	req.Color.Encode(v, b)
}

// Updates a single LED.
func (c *Client) RGBControllerUpdateSingleLed(deviceId uint32, req *RGBControllerUpdateSingleLedRequest) error {
	return c.SendRequestForDevice(deviceId, req)
}

type RGBControllerUpdateLedsRequest struct {
	LedColor []Color
}

// NetPacketId implements NetPacketCommand.
func (req *RGBControllerUpdateLedsRequest) NetPacketId() NetPacketId {
	return NetPacketIdRgbcontrollerUpdateleds
}

func (req *RGBControllerUpdateLedsRequest) Size(v Version) int {
	return 6 + 4*len(req.LedColor)
}

// NetPacketId implements DataEncoder.
func (req *RGBControllerUpdateLedsRequest) Encode(v Version, b *NetPacketDataBuilder) {
	dataSize := req.Size(v)
	b.EnsureSize(dataSize)
	b.WriteUint32(uint32(dataSize))
	b.WriteLen(req.LedColor)
	for _, color := range req.LedColor {
		color.Encode(v, b)
	}
}

// Updates LEDs.
func (c *Client) RGBControllerUpdateLeds(deviceId uint32, req *RGBControllerUpdateLedsRequest) error {
	return c.SendRequestForDevice(deviceId, req)
}

type RGBControllerUpdateZoneLedsRequest struct {
	ZoneIdx  uint32
	LedColor []Color
}

// NetPacketId implements NetPacketCommand.
func (req *RGBControllerUpdateZoneLedsRequest) NetPacketId() NetPacketId {
	return NetPacketIdRgbcontrollerUpdatezoneleds
}

func (req *RGBControllerUpdateZoneLedsRequest) Size(v Version) int {
	return 10 + 4*len(req.LedColor)
}

// NetPacketId implements DataEncoder.
func (req *RGBControllerUpdateZoneLedsRequest) Encode(v Version, b *NetPacketDataBuilder) {
	dataSize := req.Size(v)
	b.EnsureSize(dataSize)
	b.WriteUint32(uint32(dataSize))
	b.WriteUint32(req.ZoneIdx)
	b.WriteLen(req.LedColor)
	for _, color := range req.LedColor {
		color.Encode(v, b)
	}
}

// Updates zone LEDs.
func (c *Client) RGBControllerUpdateZoneLeds(deviceId uint32, req *RGBControllerUpdateZoneLedsRequest) error {
	return c.SendRequestForDevice(deviceId, req)
}

type RGBControllerSetCustomModeRequest struct {
}

// NetPacketId implements NetPacketCommand.
func (req *RGBControllerSetCustomModeRequest) NetPacketId() NetPacketId {
	return NetPacketIdRgbcontrollerSetcustommode
}

// NetPacketId implements DataEncoder.
func (req *RGBControllerSetCustomModeRequest) Encode(v Version, b *NetPacketDataBuilder) {
}

// Set custom mode.
func (c *Client) RGBControllerSetCustomMode(deviceId uint32) error {
	req := &RGBControllerSetCustomModeRequest{}
	return c.SendRequestForDevice(deviceId, req)
}

type RGBControllerUpdateModeRequest struct {
	ModeIdx int32
	Mode    *Mode
}

// NetPacketId implements NetPacketCommand.
func (req *RGBControllerUpdateModeRequest) NetPacketId() NetPacketId {
	return NetPacketIdRgbcontrollerUpdatemode
}

func (req *RGBControllerUpdateModeRequest) Size(v Version) int {
	return 8 + req.Mode.Size(v)
}

// NetPacketId implements DataEncoder.
func (req *RGBControllerUpdateModeRequest) Encode(v Version, b *NetPacketDataBuilder) {
	dataSize := req.Size(v)
	b.EnsureSize(dataSize)
	b.WriteUint32(uint32(dataSize))
	b.WriteInt32(req.ModeIdx)
	req.Mode.Encode(v, b)
}

// Updates the mode.
func (c *Client) RGBControllerUpdateMode(deviceId uint32, req *RGBControllerUpdateModeRequest) error {
	return c.SendRequestForDevice(deviceId, req)
}

type RGBControllerSaveModeRequest struct {
	ModeIdx int32
	Mode    *Mode
}

// NetPacketId implements NetPacketCommand.
func (req *RGBControllerSaveModeRequest) NetPacketId() NetPacketId {
	return NetPacketIdRgbcontrollerSavemode
}

func (req *RGBControllerSaveModeRequest) Size(v Version) int {
	return 8 + req.Mode.Size(v)
}

// NetPacketId implements DataEncoder.
func (req *RGBControllerSaveModeRequest) Encode(v Version, b *NetPacketDataBuilder) {
	dataSize := req.Size(v)
	b.EnsureSize(dataSize)
	b.WriteUint32(uint32(dataSize))
	b.WriteInt32(req.ModeIdx)
	req.Mode.Encode(v, b)
}

// Saves the mode.
func (c *Client) RGBControllerSaveMode(deviceId uint32, req *RGBControllerSaveModeRequest) error {
	return c.SendRequestForDevice(deviceId, req)
}
