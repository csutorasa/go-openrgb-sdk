package openrgb

import "fmt"

// Occurs, when an invalid magic header was found.
type InvalidPacketHeaderMagicValueError struct {
	v string
}

// Error implements error
func (e *InvalidPacketHeaderMagicValueError) Error() string {
	return fmt.Sprintf("invalid packet header magic value: %s", e.v)
}

// Occurs, when an invalid size response is returned.
type UnprocessedResponseError struct {
	total int
	read  int
}

// Error implements error
func (e *UnprocessedResponseError) Error() string {
	return fmt.Sprintf("unprocessed data in response read %d but have %d", e.read, e.total)
}

// Occurs, when an invalid net packet ID is returned.
type InvalidResponseNetPacketIdError struct {
	expected NetPacketId
	actual   NetPacketId
}

// Error implements error
func (e *InvalidResponseNetPacketIdError) Error() string {
	return fmt.Sprintf("expected net packet id %d but got %d", e.expected, e.actual)
}

// Occurs, when a timeout happens while waiting for a response.
type ResponseTimeoutError struct {
	cause error
}

// Error implements error.
func (e *ResponseTimeoutError) Error() string {
	return "response timeout"
}

func (e *ResponseTimeoutError) Unwrap() error {
	return e.cause
}

// Occurs, when the server version is lower than expected.
type InvalidServerVersionError struct {
	expected Version
	actual   Version
}

// Error implements error
func (e *InvalidServerVersionError) Error() string {
	return fmt.Sprintf("expected server version %d but got %d", e.expected, e.actual)
}
