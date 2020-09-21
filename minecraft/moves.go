package main

import (
	"fmt"
	"sync"
)

// Moves creates the moves that one person can make from one brick to another
func moves() (m sync.Map) {

	var fromKey string
	var toMoves []string

	for x := 10; x <= 50; x++ {
		for y := 10; y <= 50; y++ {
			for z := 10; z <= 50; z++ {

				fromKey = fmt.Sprintf("%v_%v_%v", x, y, z)
				toMoves = nil

				toMoves = append(toMoves, fmt.Sprintf("%v_%v_%v", x+1, y, z))
				toMoves = append(toMoves, fmt.Sprintf("%v_%v_%v", x-1, y, z))
				toMoves = append(toMoves, fmt.Sprintf("%v_%v_%v", x, y+1, z))
				toMoves = append(toMoves, fmt.Sprintf("%v_%v_%v", x, y-1, z))
				toMoves = append(toMoves, fmt.Sprintf("%v_%v_%v", x, y, z+1))
				toMoves = append(toMoves, fmt.Sprintf("%v_%v_%v", x, y, z-1))

				m.Store(fromKey, toMoves)

			}
		}
	}

	return
}
