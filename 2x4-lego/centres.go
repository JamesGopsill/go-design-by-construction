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

				fromKey = fmt.Sprintf("%v_%v_%v_0", x, y, z)

				xCentre := float64(x) * float64(5)
				yCentre := float64(y)*float64(5) + float64(5)
				zCentre := float64(z) * float64(6)

				centre[0] = xCentre
				centre[1] = yCentre
				centre[2] = zCentre

				m.Store(fromKey, centre)

				fromKey = fmt.Sprintf("%v_%v_%v_1", x, y, z)

				xCentre = float64(x)*float64(5) + float64(5)
				yCentre = float64(y) * float64(5)
				zCentre = float64(z) * float64(6)

				centre[0] = xCentre
				centre[1] = yCentre
				centre[2] = zCentre

				m.Store(fromKey, centre)

			}
		}
	}

	return
}
