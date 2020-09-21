package main

import (
	"fmt"
	"sync"
)

// moveMap creates the moves that one person can make from one brick to another
func moveMap() (m sync.Map) {

	// bw := 2 // brick height
	// bh := 2 // brick width

	var fromKey string
	var toMoves []string

	for fromX := 10; fromX <= 50; fromX++ {
		for fromY := 10; fromY <= 50; fromY++ {
			for fromZ := 10; fromZ <= 50; fromZ++ {

				fromKey = fmt.Sprintf("%v_%v_%v", fromX, fromY, fromZ)
				toMoves = nil

				// All potential moves
				/*
					for toX := fromX - bw + 1; toX < fromX+bw; toX++ {
						for toY := fromY - bh + 1; toY < fromY+bh; toY++ {

							// Add the moves
							toMoves = append(toMoves, fmt.Sprintf("%v_%v_%v", toX, toY, fromZ+1))
							//toMoves = append(toMoves, fmt.Sprintf("%v_%v_%v", toX, toY, fromZ-1))

						}
					}
				*/

				// To studs of existing brick
				toMoves = append(toMoves, fmt.Sprintf("%v_%v_%v", fromX-1, fromY, fromZ+1))
				toMoves = append(toMoves, fmt.Sprintf("%v_%v_%v", fromX-1, fromY, fromZ-1))

				toMoves = append(toMoves, fmt.Sprintf("%v_%v_%v", fromX+1, fromY, fromZ+1))
				toMoves = append(toMoves, fmt.Sprintf("%v_%v_%v", fromX+1, fromY, fromZ-1))

				toMoves = append(toMoves, fmt.Sprintf("%v_%v_%v", fromX, fromY+1, fromZ+1))
				toMoves = append(toMoves, fmt.Sprintf("%v_%v_%v", fromX, fromY+1, fromZ-1))

				toMoves = append(toMoves, fmt.Sprintf("%v_%v_%v", fromX, fromY-1, fromZ+1))
				toMoves = append(toMoves, fmt.Sprintf("%v_%v_%v", fromX, fromY-1, fromZ-1))

				m.Store(fromKey, toMoves)
			}
		}
	}

	return
}
