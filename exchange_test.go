package openrgb_test

import (
	"testing"

	"github.com/csutorasa/go-openrgb-sdk"
)

func TestExchange(t *testing.T) {
	h := openrgb.NewExchangeHandler()
	c00 := h.Create(0)
	c01 := h.Create(0)
	c10 := h.Create(1)
	c02 := h.Create(0)
	p00 := h.Pop(0)
	if c00 != p00 {
		t.Fatal()
	}
	h.Delete(0, c01)
	p01 := h.Pop(0)
	if c02 != p01 {
		t.Fatal()
	}
	p02 := h.Pop(0)
	if p02 != nil {
		t.Fatal()
	}
	var closeError error
	go func() {
		closeError = h.Close()
	}()
	r := <-c10
	if r != nil {
		t.Fatal()
	}
	if closeError != nil {
		t.Fatal()
	}
}
