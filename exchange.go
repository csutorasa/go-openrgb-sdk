package openrgb

import (
	"sync"
)

// Manages the request-responses for a connection.
// Create a new handler with NewExchangeHandler().
type ExchangeHandler struct {
	l        sync.Mutex
	requests map[NetPacketId][]chan *NetPacket
}

// Creates a new ExchangeHandler.
func NewExchangeHandler() *ExchangeHandler {
	return &ExchangeHandler{
		l:        sync.Mutex{},
		requests: map[NetPacketId][]chan *NetPacket{},
	}
}

// Adds a new handle for the command.
func (eh *ExchangeHandler) Create(commandId NetPacketId) chan *NetPacket {
	responseCh := make(chan *NetPacket)
	eh.l.Lock()
	defer eh.l.Unlock()
	requests, ok := eh.requests[commandId]
	if ok {
		eh.requests[commandId] = append(requests, responseCh)
	} else {
		eh.requests[commandId] = []chan *NetPacket{responseCh}
	}
	return responseCh
}

// Gets a handle for the command.
// Returns nil if no handler was found.
func (eh *ExchangeHandler) Pop(commandId NetPacketId) chan<- *NetPacket {
	eh.l.Lock()
	defer eh.l.Unlock()
	requests, ok := eh.requests[commandId]
	if !ok {
		return nil
	}
	if len(requests) == 0 {
		return nil
	}
	req := requests[0]
	eh.requests[commandId] = requests[1:]
	return req
}

// Deletes a handler.
func (eh *ExchangeHandler) Delete(commandId NetPacketId, responseCh chan *NetPacket) {
	eh.l.Lock()
	defer eh.l.Unlock()
	requests, ok := eh.requests[commandId]
	if !ok {
		return
	}
	if len(requests) == 0 {
		return
	}
	reqs := []chan *NetPacket{}
	for _, r := range requests {
		if r != responseCh {
			reqs = append(reqs, r)
		}
	}
	eh.requests[commandId] = reqs
}

// Sends a nil to all active handles.
func (eh *ExchangeHandler) Close() error {
	eh.l.Lock()
	defer eh.l.Unlock()
	for _, requests := range eh.requests {
		for _, req := range requests {
			req <- nil
		}
	}
	return nil
}
