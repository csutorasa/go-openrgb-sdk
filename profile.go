package openrgb

import "context"

type RequestProfileListRequest struct {
}

// NetPacketId implements NetPacketCommand.
func (req *RequestProfileListRequest) NetPacketId() NetPacketId {
	return NetPacketIdRequestProfileList
}

// NetPacketId implements DataEncoder.
func (req *RequestProfileListRequest) Encode(v Version, b *NetPacketDataBuilder) {
}

type RequestProfileListResponse struct {
	// Profile name strings
	Names []string
}

// NetPacketId implements NetPacketCommand.
func (req *RequestProfileListResponse) NetPacketId() NetPacketId {
	return NetPacketIdRequestProfileList
}

// Decode implements DataDecoder.
func (res *RequestProfileListResponse) Decode(v Version, p *NetPacketDataParser) error {
	_, err := p.ReadUint32()
	if err != nil {
		return err
	}
	numProfiles, err := p.ReadUint16()
	if err != nil {
		return err
	}
	profiles := make([]string, numProfiles)
	for i := uint16(0); i < numProfiles; i++ {
		name, err := p.ReadString()
		if err != nil {
			return err
		}
		profiles[i] = name
	}
	res.Names = profiles
	return nil
}

// Gets the available profiles.
func (c *Client) RequestProfileList() (*RequestProfileListResponse, error) {
	err := c.AssertServerVersion(2)
	if err != nil {
		return nil, err
	}
	req := &RequestProfileListRequest{}
	res := new(RequestProfileListResponse)
	_, err = c.SendRequestAndExpectResponse(req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Gets the available profiles.
func (c *Client) RequestProfileListCtx(ctx context.Context) (*RequestProfileListResponse, error) {
	err := c.AssertServerVersion(2)
	if err != nil {
		return nil, err
	}
	req := &RequestProfileListRequest{}
	res := new(RequestProfileListResponse)
	_, err = c.SendRequestAndExpectResponseCtx(ctx, req, res)
	if err != nil {
		return nil, err
	}
	return res, nil
}

type RequestSaveProfileRequest struct {
	// Profile name
	ProfileName string
}

// NetPacketId implements NetPacketCommand.
func (req *RequestSaveProfileRequest) NetPacketId() NetPacketId {
	return NetPacketIdRequestSaveProfile
}

// NetPacketId implements DataEncoder.
func (req *RequestSaveProfileRequest) Encode(v Version, b *NetPacketDataBuilder) {
	b.WriteBytes([]byte(req.ProfileName))
}

// Saves the profile.
func (c *Client) RequestSaveProfile(req *RequestSaveProfileRequest) error {
	err := c.AssertServerVersion(2)
	if err != nil {
		return err
	}
	return c.SendRequest(req)
}

type RequestLoadProfileRequest struct {
	// Profile name
	ProfileName string
}

// NetPacketId implements NetPacketCommand.
func (req *RequestLoadProfileRequest) NetPacketId() NetPacketId {
	return NetPacketIdRequestLoadProfile
}

// NetPacketId implements DataEncoder.
func (req *RequestLoadProfileRequest) Encode(v Version, b *NetPacketDataBuilder) {
	b.WriteBytes([]byte(req.ProfileName))
}

// Loads the profile.
func (c *Client) RequestLoadProfile(req *RequestLoadProfileRequest) error {
	err := c.AssertServerVersion(2)
	if err != nil {
		return err
	}
	return c.SendRequest(req)
}

type RequestDeleteProfileRequest struct {
	// Profile name
	ProfileName string
}

// NetPacketId implements NetPacketCommand.
func (req *RequestDeleteProfileRequest) NetPacketId() NetPacketId {
	return NetPacketIdRequestDeleteProfile
}

// NetPacketId implements DataEncoder.
func (req *RequestDeleteProfileRequest) Encode(v Version, b *NetPacketDataBuilder) {
	b.WriteBytes([]byte(req.ProfileName))
}

// Deletes the profile.
func (c *Client) RequestDeleteProfile(req *RequestDeleteProfileRequest) error {
	err := c.AssertServerVersion(2)
	if err != nil {
		return err
	}
	return c.SendRequest(req)
}
