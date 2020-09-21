package main

import (
	"fmt"
	"math"
	"sync"
)

// IncrementBrick increments the bricks to the combinations and adds them to the cMap
func incrementBrick(
	c Combination,
	cMap *sync.Map,
	moves *sync.Map,
	centres *sync.Map,
	wg *sync.WaitGroup) {

	// Defer wg.Done until we complete our run through of the function
	defer wg.Done()

	// For each value in the c.Next combi
	for _, b := range c.Next {

		// Create a new combination
		var newC Combination

		// Append to the sequence
		newC.Sequence = append(c.Sequence, b)

		// ###############
		// Get the next positions to go to
		var next []string

		// For each block in the sequence
		for _, block := range newC.Sequence {
			// Find the go to moves
			movetoblocks, ok := moves.Load(block)
			if !ok {
				panic("Error loading movetoblocks")
			}
			for _, movetoblock := range movetoblocks.([]string) {
				// Check if the move to block is not already in the sequence or in the next moves (i.e. duplicate)
				if !stringInSlice(movetoblock, newC.Sequence) || !stringInSlice(movetoblock, next) {
					next = append(next, movetoblock)
				}
			}
		}
		newC.Next = next
		// ################

		// Load the centre for the block
		u, _ := centres.Load(b)
		// Get the pairwise distance from the previouse combination
		newC.Dist = c.Dist
		// Now add the pairwise distances of the new block to all the previous blocks
		for _, pb := range c.Sequence {
			v, _ := centres.Load(pb)
			first := math.Pow(u.([]float64)[0]-v.([]float64)[0], 2)
			second := math.Pow(u.([]float64)[1]-v.([]float64)[1], 2)
			third := math.Pow(u.([]float64)[2]-v.([]float64)[2], 2)
			// Sum the distances
			newC.Dist += first + second + third
		}

		// Take the previous number of paths from the last combination
		newC.Paths = c.Paths

		// Create the unique key for the combination
		newCKey := fmt.Sprintf("%02d_%v_%v", len(newC.Sequence), len(newC.Next), newC.Dist)

		// Check if it arleady exists or store a new combination
		e, exists := cMap.Load(newCKey)
		if exists {
			update := e.(Combination)
			update.Paths = update.Paths + newC.Paths
			cMap.Store(newCKey, update)
		} else {
			cMap.Store(newCKey, newC)
		}
	}

	return
}
