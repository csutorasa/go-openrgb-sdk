package main

import (
	"time"

	"github.com/csutorasa/go-openrgb-sdk"
)

func main() {
	// Create a new client.
	c, err := openrgb.NewDefaultClient()
	if err != nil {
		panic(err)
	}
	defer c.Close()
	err = c.Initialize("Rainbow")
	if err != nil {
		panic(err)
	}
	// Update colors indefinitely.
	color := openrgb.ColorRed
	counter := int64(0)
	err = openrgb.Loop(c, 5*time.Millisecond, func(controllers openrgb.Controllers) error {
		for i, controller := range controllers {
			colors := openrgb.NewColors(len(controller.Colors), color)
			err = c.RGBControllerUpdateLeds(uint32(i), &openrgb.RGBControllerUpdateLedsRequest{
				LedColor: colors,
			})
			if err != nil {
				return err
			}
		}
		moveColor(&color, counter)
		counter++
		return nil
	})
	if err != nil {
		panic(err)
	}
}

func moveColor(color *openrgb.Color, counter int64) {
	state := (counter % (6 * 256))
	ratio := float64(state%256) / float64(256)
	if state < 256 {
		color.R = 255
		color.G = uint8(255 * ratio)
		color.B = 0
	} else if state < 2*256 {
		color.R = uint8(255 * (1 - ratio))
		color.G = 255
		color.B = 0
	} else if state < 3*256 {
		color.R = 0
		color.G = 255
		color.B = uint8(255 * ratio)
	} else if state < 4*256 {
		color.R = 0
		color.G = uint8(255 * (1 - ratio))
		color.B = 255
	} else if state < 5*256 {
		color.R = uint8(255 * ratio)
		color.G = 0
		color.B = 255
	} else {
		color.R = 255
		color.G = 0
		color.B = uint8(255 * (1 - ratio))
	}
}
