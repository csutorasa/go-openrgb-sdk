package openrgb

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net"
)

// Type for handling requests coming from the server.
type DeviceListUpdatedHandler func(*DeviceListUpdatedResponse)

// Client for communicating with an OpenRGB server.
// Create a new client with NewClient().
// Calling RequestProtocolVersion() is recommended to sync versions of server and client.
type Client struct {
	conn                     io.ReadWriteCloser
	serverVersion            Version
	exchangeHandler          *ExchangeHandler
	decoder                  *NetPacketDecoder
	encoder                  *NetPacketEncoder
	deviceListUpdatedHandler DeviceListUpdatedHandler
}

// Creates a new client from a connection.
func NewClient(conn net.Conn) *Client {
	c := &Client{
		conn:                     conn,
		serverVersion:            0,
		exchangeHandler:          NewExchangeHandler(),
		decoder:                  NewNetPacketDecoder(conn),
		encoder:                  NewNetPacketEncoder(conn),
		deviceListUpdatedHandler: nil,
	}
	go c.readLoop()
	return c
}

// Creates a new client from a hostname and a port number.
func NewClientHostPort(host string, port int) (*Client, error) {
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", host, port))
	if err != nil {
		return nil, err
	}
	return NewClient(conn), nil
}

// Creates a new default client.
func NewDefaultClient() (*Client, error) {
	return NewClientHostPort("localhost", defaultPort)
}

// Creates a new client to locahost from a port number.
func NewLocalClient(port int) (*Client, error) {
	return NewClientHostPort("localhost", port)
}

// Closes all resources used by the client.
func (c *Client) Close() error {
	err := c.conn.Close()
	if err != nil {
		return err
	}
	return c.exchangeHandler.Close()
}

// Registers or overrides a handler to handle requests from the server.
func (c *Client) DeviceListUpdatedHandler(h DeviceListUpdatedHandler) {
	c.deviceListUpdatedHandler = h
}

// Sends a net packet to the server and expects no answer.
func (c *Client) SendPacket(req *NetPacket) error {
	return c.encoder.Encode(req)
}

// Sends a request to the server and expects no answer.
func (c *Client) SendRequest(req NetPacketRequest) error {
	return c.SendPacket(c.packetFromRequest(req))
}

// Sends a request to the server and expects no answer.
func (c *Client) SendRequestForDevice(deviceId uint32, req NetPacketRequest) error {
	return c.SendPacket(c.packetFromRequestForDevice(deviceId, req))
}

// Sends a packet to the server and expects an answer packet.
func (c *Client) SendPacketAndExpectPacket(req *NetPacket) (*NetPacket, error) {
	done, err := c.sendRequestAndReturnChan(req)
	if err != nil {
		return nil, err
	}
	return <-done, nil
}

// Sends a packet to the server and expects an answer packet.
func (c *Client) SendPacketAndExpectPacketCtx(ctx context.Context, req *NetPacket) (*NetPacket, error) {
	done, err := c.sendRequestAndReturnChan(req)
	if err != nil {
		return nil, err
	}
	select {
	case p := <-done:
		return p, nil
	case <-ctx.Done():
		return nil, &ResponseTimeoutError{
			cause: ctx.Err(),
		}
	}
}

// Sends a request to the server and expects an answer response.
func (c *Client) SendRequestAndExpectResponse(req NetPacketRequest, res NetPacketResponse) (*NetPacket, error) {
	p, err := c.SendPacketAndExpectPacket(c.packetFromRequest(req))
	if err != nil {
		return nil, err
	}
	return p, c.decodeResponse(p, res)
}

// Sends a request to the server and expects an answer response.
func (c *Client) SendRequestAndExpectResponseCtx(ctx context.Context, req NetPacketRequest, res NetPacketResponse) (*NetPacket, error) {
	p, err := c.SendPacketAndExpectPacketCtx(ctx, c.packetFromRequest(req))
	if err != nil {
		return nil, err
	}
	return p, c.decodeResponse(p, res)
}

// Sends a request for a device to the server and expects an answer response.
func (c *Client) SendRequestForDeviceAndExpectResponse(deviceId uint32, req NetPacketRequest, res NetPacketResponse) (*NetPacket, error) {
	p, err := c.SendPacketAndExpectPacket(c.packetFromRequestForDevice(deviceId, req))
	if err != nil {
		return nil, err
	}
	return p, c.decodeResponse(p, res)
}

// Sends a request for a device to the server and expects an answer response.
func (c *Client) SendRequestForDeviceAndExpectResponseCtx(ctx context.Context, deviceId uint32, req NetPacketRequest, res NetPacketResponse) (*NetPacket, error) {
	p, err := c.SendPacketAndExpectPacketCtx(ctx, c.packetFromRequestForDevice(deviceId, req))
	if err != nil {
		return nil, err
	}
	return p, c.decodeResponse(p, res)
}

func (c *Client) packetFromRequest(req NetPacketRequest) *NetPacket {
	return c.packetFromRequestForDevice(0, req)
}

func (c *Client) packetFromRequestForDevice(deviceId uint32, req NetPacketRequest) *NetPacket {
	v := c.CommonVersion()
	b := &NetPacketDataBuilder{}
	req.Encode(v, b)
	return NewNetPacket(req.NetPacketId(), deviceId, b.Bytes())
}

func (c *Client) decodeResponse(p *NetPacket, res NetPacketResponse) error {
	v := c.CommonVersion()
	if p.Header.PktId != res.NetPacketId() {
		return &InvalidResponseNetPacketIdError{
			expected: res.NetPacketId(),
			actual:   p.Header.PktId,
		}
	}
	parser := p.DataParser()
	err := res.Decode(v, parser)
	if err != nil {
		return err
	}
	return parser.AssertAllRead()
}

func (c *Client) sendRequestAndReturnChan(req *NetPacket) (<-chan *NetPacket, error) {
	done := c.exchangeHandler.Create(req.Header.PktId)
	err := c.SendPacket(req)
	if err != nil {
		c.exchangeHandler.Delete(req.Header.PktId, done)
		return nil, err
	}
	return done, nil
}

func (c *Client) deviceListUpdated(p *NetPacket) {
	res := new(DeviceListUpdatedResponse)
	parser := p.DataParser()
	err := res.Decode(c.CommonVersion(), parser)
	if err != nil {
		return
	}
	err = parser.AssertAllRead()
	if err != nil {
		return
	}
	if c.deviceListUpdatedHandler != nil {
		c.deviceListUpdatedHandler(res)
	}
}

func (c *Client) readLoop() {
	for {
		res, err := c.decoder.Decode()
		if errors.Is(err, net.ErrClosed) || errors.Is(err, io.EOF) {
			// Connection is closed
			return
		}
		if err != nil {
			// Ignore random errors
			continue
		}
		if res.Header.PktId == NetPacketIdDeviceListUpdated {
			c.deviceListUpdated(res)
			continue
		}
		done := c.exchangeHandler.Pop(res.Header.PktId)
		if done == nil {
			// Ignore if response is not expected
			continue
		}
		done <- res
		close(done)
	}
}
