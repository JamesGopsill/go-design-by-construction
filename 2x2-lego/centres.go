package main

import (
	"fmt"
	"sync"
)

func centres() (m sync.Map) {
	var fromKey string
	var centre [3]float64

	for x := 10; x <= 50; x++ {
		for y := 10; y <= 50; y++ {
			for z := 10; z <= 50; z++ {

				fromKey = fmt.Sprintf("%v_%v_%v", x, y, z)

				xCentre := float64(x) * float64(5) // + float64(2.5) - just a translation relative to the index so can omit
				yCentre := float64(y) * float64(5) // - float64(2.5)
				zCentre := float64(z) * float64(6) // - float64(3)

				centre[0] = xCentre //append(centre, xCentre)
				centre[1] = yCentre //append(centre, yCentre)
				centre[2] = zCentre //append(centre, zCentre)

				m.Store(fromKey, centre)

			}
		}
	}

	return
}
