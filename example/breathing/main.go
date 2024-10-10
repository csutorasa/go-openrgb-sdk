package main

import (
	"math"
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
	err = c.Initialize("Breathing")
	if err != nil {
		panic(err)
	}
	// Update colors indefinitely.
	baseColor := openrgb.ColorRed
	color := baseColor
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
		moveColor(&color, baseColor, counter)
		counter++
		return nil
	})
	if err != nil {
		panic(err)
	}
}

func moveColor(color *openrgb.Color, baseColor openrgb.Color, counter int64) {
	state := (counter % 512)
	ratio := float64(state) / 256
	if state >= 256 {
		ratio = 2 - ratio
	}
	ratio = math.Min(math.Max(0, ratio), 1)
	color.R = uint8(float64(baseColor.R) * ratio)
	color.G = uint8(float64(baseColor.G) * ratio)
	color.B = uint8(float64(baseColor.B) * ratio)
}
