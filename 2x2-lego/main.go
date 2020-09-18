package main

import (
	"fmt"
	"log"
	"strings"
	"sync"
)

// Combination is a struct containing details of a combination
type Combination struct {
	Sequence []string
	Moves    []string
	NoMoves  []string
	Dist     float64
	Paths    int64
}

func main() {
	log.Println("2x2 Lego Construction")

	startPos := "30_30_30"

	log.Println("Creating Maps")
	moves := moveMap()
	nomoves := noMoveMap()
	cents := centres()

	// #####################

	// Creating the initial Combination
	var start Combination

	tmp, ok := moves.Load("30_30_30")
	if !ok {
		panic("Error: Centre key not found")
	}
	move, ok := tmp.([]string)
	if !ok {
		panic("Error in typing")
	}
	tmp, ok = nomoves.Load(startPos)
	if !ok {
		panic("Error: Centre key not found")
	}
	nomove, ok := tmp.([]string)
	if !ok {
		panic("Error in typing")
	}
	// spew.Dump(move)
	//spew.Dump(nomove)

	start.Sequence = append(start.Sequence, "30_30_30")
	start.Moves = append(start.Moves, move...)
	start.NoMoves = append(start.NoMoves, nomove...)
	start.Dist = 0.
	start.Paths = 1
	startKey := fmt.Sprintf("01_%v_%v", len(start.Moves), start.Dist)

	// Add the first combination to the sync map
	var cMap sync.Map
	cMap.Store(startKey, start)

	// #################

	// Now run through the different combinations
	for i := 2; i <= 9; i++ {

		// initialise some vars
		currentKeyStartsWith := fmt.Sprintf("%02d", i)
		previousKeyStartsWith := fmt.Sprintf("%02d", i-1)
		var wg sync.WaitGroup

		// For each in the combination sync.Map
		cMap.Range(func(k, v interface{}) bool {

			// If the prefix is of a sequence from the previous combination set, increment a brick
			if strings.HasPrefix(k.(string), previousKeyStartsWith) {

				// fmt.Printf("Key: %v, Paths to Seq: %v \n", k.(string), v.(Combination).Paths)

				wg.Add(1)
				go incrementBrick(v.(Combination), &cMap, &moves, &nomoves, &cents, &wg)

			}

			return true
		})

		wg.Wait()

		// ###############
		// Report

		// Report on the previous combination while incrementing the next
		fmt.Printf("\n%02d brick combinations\n", i)

		uniqueCombinations := 0
		var totalPaths int64

		cMap.Range(func(k, v interface{}) bool {

			// If the prefix is of a sequence from the previous combination set, increment a brick
			if strings.HasPrefix(k.(string), currentKeyStartsWith) {
				// spew.Dump(v.(Combination).Sequence)
				uniqueCombinations++
				totalPaths += v.(Combination).Paths
			}

			return true
		})

		// Print out the facts from the brick combination set
		fmt.Printf("Unique Combinations: %v \n", uniqueCombinations)
		fmt.Printf("Paths: %v \n", totalPaths)

	}
}
