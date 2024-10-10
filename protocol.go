package openrgb

import (
	"encoding/binary"
)

// Default TCP server port. This is "ORGB" on a telephone keypad.
const defaultPort = 6742

// Magic string for packet headers.
const netPacketHeaderMagicValue = "ORGB"

// Byte order for the protocol.
var protocolByteOrder binary.ByteOrder = binary.LittleEndian

// Packet command ID
type NetPacketId uint32

const (
	// Request RGBController device count from server
	NetPacketIdRequestControllerCount NetPacketId = 0
	// Request RGBController data block
	NetPacketIdRequestControllerData NetPacketId = 1
	// Request OpenRGB SDK protocol version from server
	NetPacketIdRequestProtocolVersion NetPacketId = 40
	// Send client name string to server
	NetPacketIdSetClientName NetPacketId = 50
	// Indicate to clients that device list has updated
	NetPacketIdDeviceListUpdated NetPacketId = 100
	// Request profile list
	NetPacketIdRequestProfileList NetPacketId = 150
	// Save current configuration in a new profile
	NetPacketIdRequestSaveProfile NetPacketId = 151
	// Load a given profile
	NetPacketIdRequestLoadProfile NetPacketId = 152
	// Delete a given profile
	NetPacketIdRequestDeleteProfile NetPacketId = 153
	//RGBController::ResizeZone()
	NetPacketIdRgbcontrollerResizezone NetPacketId = 1000
	// RGBController::UpdateLEDs()
	NetPacketIdRgbcontrollerUpdateleds NetPacketId = 1050
	// RGBController::UpdateZoneLEDs()
	NetPacketIdRgbcontrollerUpdatezoneleds NetPacketId = 1051
	// RGBController::UpdateSingleLED()
	NetPacketIdRgbcontrollerUpdatesingleled NetPacketId = 1052
	// RGBController::SetCustomMode()
	NetPacketIdRgbcontrollerSetcustommode NetPacketId = 1100
	// RGBController::UpdateMode()
	NetPacketIdRgbcontrollerUpdatemode NetPacketId = 1101
	// RGBController::SaveMode()
	NetPacketIdRgbcontrollerSavemode NetPacketId = 1102
)

// Packet for requests and responses.
type NetPacket struct {
	Header *NetPacketHeader
	Data   NetPacketData
}

// Packet header for requests and responses.
type NetPacketHeader struct {
	// Magic value, "ORGB"
	PktMagic string
	// Device Index
	PktDevIdx uint32
	// Packet ID
	PktId NetPacketId
	// Packet Size
	PktSize uint32
}

type NetPacketData []byte

func NewNetPacket(commandId NetPacketId, deviceId uint32, data []byte) *NetPacket {
	return &NetPacket{
		Header: &NetPacketHeader{
			PktMagic:  netPacketHeaderMagicValue,
			PktDevIdx: deviceId,
			PktId:     commandId,
			PktSize:   uint32(len(data)),
		},
		Data: data,
	}
}

// Net packet command.
type NetPacketCommand interface {
	// Gets the packet ID for the command.
	NetPacketId() NetPacketId
}
