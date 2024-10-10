package openrgb

import "time"

func Loop(c *Client, delay time.Duration, loop func(Controllers) error) error {
	cache := NewControllerCache(c)
	c.DeviceListUpdatedHandler(func(*DeviceListUpdatedResponse) {
		cache.Invalidate()
	})
	ticker := time.NewTicker(delay)
	defer ticker.Stop()
	for range ticker.C {
		controllers, err := cache.Controllers()
		if err != nil {
			return err
		}
		loop(controllers)
	}
	return nil
}
