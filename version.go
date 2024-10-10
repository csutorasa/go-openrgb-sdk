package openrgb

import (
	"context"
	"time"
)

type Version uint32

// Highest version supported by the SDK
const ClientVersion Version = 3

type RequestProtocolVersionRequest struct {
	ClientVersion Version
}

// NetPacketId implements NetPacketCommand.
func (req *RequestProtocolVersionRequest) NetPacketId() NetPacketId {
	return NetPacketIdRequestProtocolVersion
}

// NetPacketId implements DataEncoder.
func (req *RequestProtocolVersionRequest) Encode(v Version, b *NetPacketDataBuilder) {
	b.EnsureSize(4)
	b.WriteUint32(uint32(req.ClientVersion))
}

type RequestProtocolVersionResponse struct {
	ServerVersion Version
}

// NetPacketId implements NetPacketCommand.
func (req *RequestProtocolVersionResponse) NetPacketId() NetPacketId {
	return NetPacketIdRequestProtocolVersion
}

// Decode implements DataDecoder.
func (res *RequestProtocolVersionResponse) Decode(v Version, p *NetPacketDataParser) error {
	version, err := p.ReadUint32()
	if err != nil {
		return err
	}
	res.ServerVersion = Version(version)
	return nil
}

// Gets the server protocol version.
func (c *Client) RequestProtocolVersion() error {
	req := &RequestProtocolVersionRequest{
		ClientVersion: ClientVersion,
	}
	res := new(RequestProtocolVersionResponse)
	_, err := c.SendRequestAndExpectResponse(req, res)
	if err != nil {
		return err
	}
	c.serverVersion = res.ServerVersion
	return nil
}

// Gets the server protocol version.
func (c *Client) RequestProtocolVersionCtx(ctx context.Context) error {
	req := &RequestProtocolVersionRequest{
		ClientVersion: ClientVersion,
	}
	res := new(RequestProtocolVersionResponse)
	_, err := c.SendRequestAndExpectResponseCtx(ctx, req, res)
	if err != nil {
		return err
	}
	c.serverVersion = res.ServerVersion
	return nil
}

type SetClientNameRequest struct {
	// Client name
	ClientName string
}

// NetPacketId implements NetPacketCommand.
func (req *SetClientNameRequest) NetPacketId() NetPacketId {
	return NetPacketIdSetClientName
}

// NetPacketId implements DataEncoder.
func (req *SetClientNameRequest) Encode(v Version, b *NetPacketDataBuilder) {
	b.WriteBytes([]byte(req.ClientName))
}

// Sets the client name.
func (c *Client) SetClientName(req *SetClientNameRequest) error {
	err := c.AssertServerVersion(1)
	if err != nil {
		return err
	}
	return c.SendRequest(req)
}

// Initializes the client.
func (c *Client) Initialize(name string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	err := c.RequestProtocolVersionCtx(ctx)
	if err != nil {
		return err
	}
	if c.CommonVersion() < 1 {
		return nil
	}
	err = c.SetClientName(&SetClientNameRequest{
		ClientName: name,
	})
	if err != nil {
		return err
	}
	return nil
}

// Finds the lowest common version between the client and the server.
// RequestProtocolVersion() needs to be called before this function.
func (c *Client) CommonVersion() Version {
	if c.serverVersion < ClientVersion {
		return c.serverVersion
	}
	return ClientVersion
}

// Returns an error if server version is lower that the given version.
// RequestProtocolVersion() needs to be called before this function.
func (c *Client) AssertServerVersion(v Version) error {
	if v > c.serverVersion {
		return &InvalidServerVersionError{
			expected: v,
			actual:   c.serverVersion,
		}
	}
	return nil
}
