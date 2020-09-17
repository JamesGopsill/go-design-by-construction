package main

import (
	"fmt"
	"sync"
)

// moveMap creates the moves that one person can make from one brick to another
func noMoveMap() (m sync.Map) {

	bw := 2 // brick height
	bh := 2 // brick width

	var fromKey string
	var toMoves []string

	for fromX := 10; fromX <= 50; fromX++ {
		for fromY := 10; fromY <= 50; fromY++ {
			for fromZ := 10; fromZ <= 50; fromZ++ {

				fromKey = fmt.Sprintf("%v_%v_%v", fromX, fromY, fromZ)
				toMoves = nil

				for toX := fromX - bw + 1; toX < fromX+bw; toX++ {
					for toY := fromY - bh + 1; toY < fromY+bh; toY++ {

						// Add the moves
						toMoves = append(toMoves, fmt.Sprintf("%v_%v_%v", toX, toY, fromZ))

					}
				}

				m.Store(fromKey, toMoves)
			}
		}
	}

	return
}
