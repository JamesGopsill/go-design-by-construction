package main

import (
	"fmt"
	"sync"
)

// moveMap creates the moves that one person can make from one brick to another
func noMoveMap() (m sync.Map) {

	fromBW := 0
	fromBH := 0
	toBW := 0
	toBH := 0
	xMin := 0
	xMax := 0
	yMin := 0
	yMax := 0

	var fromKey string
	var toMoves []string

	for fromX := 10; fromX <= 50; fromX++ {
		for fromY := 10; fromY <= 50; fromY++ {
			for fromZ := 10; fromZ <= 50; fromZ++ {

				toMoves = nil

				// 0 Degree from brick

				fromKey = fmt.Sprintf("%v_%v_%v_0", fromX, fromY, fromZ)
				fromBW = 2
				fromBH = 4

				// 0 degree to brick
				toBW = 2
				toBH = 4

				xMin = fromX - toBW + 1
				xMax = fromX + fromBW - 1
				yMin = fromY - toBH + 1
				yMax = fromY + fromBH - 1

				for toX := xMin; toX <= xMax; toX++ {
					for toY := yMin; toY <= yMax; toY++ {
						toMoves = append(toMoves, fmt.Sprintf("%v_%v_%v_0", toX, toY, fromZ))
					}
				}

				// 90 degree to brick
				toBW = 4
				toBH = 2

				xMin = fromX - toBW + 1
				xMax = fromX + fromBW - 1
				yMin = fromY - toBH + 1
				yMax = fromY + fromBH - 1

				for toX := xMin; toX <= xMax; toX++ {
					for toY := yMin; toY <= yMax; toY++ {
						toMoves = append(toMoves, fmt.Sprintf("%v_%v_%v_1", toX, toY, fromZ))
					}
				}

				// Store the combination
				m.Store(fromKey, toMoves)

				toMoves = nil

				fromKey = fmt.Sprintf("%v_%v_%v_1", fromX, fromY, fromZ)
				fromBW = 4
				fromBH = 2

				// 0 degree to brick
				toBW = 2
				toBH = 4

				xMin = fromX - toBW + 1
				xMax = fromX + fromBW - 1
				yMin = fromY - toBH + 1
				yMax = fromY + fromBH - 1

				for toX := xMin; toX <= xMax; toX++ {
					for toY := yMin; toY <= yMax; toY++ {
						toMoves = append(toMoves, fmt.Sprintf("%v_%v_%v_0", toX, toY, fromZ))
					}
				}

				// 90 degree to brick
				toBW = 4
				toBH = 2

				xMin = fromX - toBW + 1
				xMax = fromX + fromBW - 1
				yMin = fromY - toBH + 1
				yMax = fromY + fromBH - 1

				for toX := xMin; toX <= xMax; toX++ {
					for toY := yMin; toY <= yMax; toY++ {
						toMoves = append(toMoves, fmt.Sprintf("%v_%v_%v_1", toX, toY, fromZ))
					}
				}

				// Store the combination
				m.Store(fromKey, toMoves)

			}
		}
	}

	return
}
