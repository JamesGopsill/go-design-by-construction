package main

import (
	"fmt"
	"sync"
)

// Centres creates a map of the centre positions for the bricks
func centres() (m sync.Map) {
	var fromKey string
	var centre []float64

	for x := 10; x <= 50; x++ {
		for y := 10; y <= 50; y++ {
			for z := 10; z <= 50; z++ {

				fromKey = fmt.Sprintf("%v_%v_%v", x, y, z)

				centre = nil
				centre = append(centre, float64(x))
				centre = append(centre, float64(y))
				centre = append(centre, float64(z))

				m.Store(fromKey, centre)

			}
		}
	}
	return
}
