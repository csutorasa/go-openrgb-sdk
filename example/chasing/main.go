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
	err = c.Initialize("Chasing")
	if err != nil {
		panic(err)
	}
	// Update colors indefinitely.
	baseColor := openrgb.ColorRed
	counter := int64(0)
	err = openrgb.Loop(c, 50*time.Millisecond, func(controllers openrgb.Controllers) error {
		for i, controller := range controllers {
			if len(controller.Leds) < 10 {
				continue
			}
			colors := openrgb.NewBlackColors(len(controller.Colors))
			setColors(colors, baseColor, counter)
			err = c.RGBControllerUpdateLeds(uint32(i), &openrgb.RGBControllerUpdateLedsRequest{
				LedColor: colors,
			})
			if err != nil {
				return err
			}
		}
		counter++
		return nil
	})
	if err != nil {
		panic(err)
	}
}

func setColors(colors openrgb.Colors, baseColor openrgb.Color, counter int64) {
	state := int(counter % int64(len(colors)*2))
	center := state
	if state >= len(colors) {
		center = 2*len(colors) - state
	}
	for i := range colors {
		distance := intAbs(i - center)
		if distance > 4 {
			continue
		}
		r := 0.2 * float64(5-distance)
		colors[i] = openrgb.Color{
			R: uint8(float64(baseColor.R) * r),
			G: uint8(float64(baseColor.G) * r),
			B: uint8(float64(baseColor.B) * r),
		}
	}
}

func intAbs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
