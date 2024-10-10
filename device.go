package openrgb

import "context"

type RequestControllerCountRequest struct {
}

// NetPacketId implements NetPacketCommand.
func (req *RequestControllerCountRequest) NetPacketId() NetPacketId {
	return NetPacketIdRequestControllerCount
}

// NetPacketId implements DataEncoder.
func (req *RequestControllerCountRequest) Encode(v Version, b *NetPacketDataBuilder) {
}

type RequestControllerCountResponse struct {
	Count uint32
}

// NetPacketId implements NetPacketCommand.
func (req *RequestControllerCountResponse) NetPacketId() NetPacketId {
	return NetPacketIdRequestControllerCount
}

// Decode implements DataDecoder.
func (res *RequestControllerCountResponse) Decode(v Version, p *NetPacketDataParser) error {
	count, err := p.ReadUint32()
	if err != nil {
		return err
	}
	res.Count = count
	return nil
}

// Gets the number of available devices.
func (c *Client) RequestControllerCount() (*RequestControllerCountResponse, error) {
	req := &RequestControllerCountRequest{}
	res := new(RequestControllerCountResponse)
	_, err := c.SendRequestAndExpectResponse(req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Gets the number of available devices.
func (c *Client) RequestControllerCountCtx(ctx context.Context) (*RequestControllerCountResponse, error) {
	req := &RequestControllerCountRequest{}
	res := new(RequestControllerCountResponse)
	_, err := c.SendRequestAndExpectResponseCtx(ctx, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type RequestControllerDataRequest struct {
}

// NetPacketId implements NetPacketCommand.
func (req *RequestControllerDataRequest) NetPacketId() NetPacketId {
	return NetPacketIdRequestControllerData
}

// NetPacketId implements DataEncoder.
func (req *RequestControllerDataRequest) Encode(v Version, b *NetPacketDataBuilder) {
	if v < 1 {
		return
	}
	b.EnsureSize(4)
	b.WriteUint32(uint32(v))
}

type RequestControllerDataResponse struct {
	Controller *ControllerData
}

// NetPacketId implements NetPacketCommand.
func (req *RequestControllerDataResponse) NetPacketId() NetPacketId {
	return NetPacketIdRequestControllerData
}

// Decode implements DataDecoder.
func (res *RequestControllerDataResponse) Decode(v Version, p *NetPacketDataParser) error {
	c := &ControllerData{}
	err := c.Decode(v, p)
	if err != nil {
		return err
	}
	res.Controller = c
	return nil
}

// Gets the details of a device.
func (c *Client) RequestControllerData(deviceId uint32) (*RequestControllerDataResponse, error) {
	req := &RequestControllerDataRequest{}
	res := new(RequestControllerDataResponse)
	_, err := c.SendRequestForDeviceAndExpectResponse(deviceId, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Gets the details of a device.
func (c *Client) RequestControllerDataCtx(ctx context.Context, deviceId uint32) (*RequestControllerDataResponse, error) {
	req := &RequestControllerDataRequest{}
	res := new(RequestControllerDataResponse)
	_, err := c.SendRequestForDeviceAndExpectResponseCtx(ctx, deviceId, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type DeviceListUpdatedResponse struct {
}

// NetPacketId implements NetPacketCommand.
func (req *DeviceListUpdatedResponse) NetPacketId() NetPacketId {
	return NetPacketIdDeviceListUpdated
}

// Decode implements DataDecoder.
func (res *DeviceListUpdatedResponse) Decode(v Version, p *NetPacketDataParser) error {
	return nil
}
