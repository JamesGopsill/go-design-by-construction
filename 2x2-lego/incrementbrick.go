package main

import (
	"fmt"
	"math"
	"sync"
)

func incrementBrick(c Combination, cMap *sync.Map, moves *sync.Map, nomoves *sync.Map, cents *sync.Map, wg *sync.WaitGroup) {

	// Defer wg.Done until we complete our run through of the function
	defer wg.Done()

	for _, blockToAdd := range c.Moves {
		// Create a new combination
		var newC Combination

		// Append to the sequence
		newC.Sequence = append(c.Sequence, blockToAdd)

		// Update the no moves slice
		newC.NoMoves = c.NoMoves
		tmp, ok := nomoves.Load(blockToAdd)
		if !ok {
			panic("Error loading no moves")
		}
		nms, ok := tmp.([]string)
		if !ok {
			panic("Error loading no moves for the new block")
		}

		for _, nm := range nms {
			if !stringInSlice(nm, newC.NoMoves) {
				newC.NoMoves = append(newC.NoMoves, nm)
			}
		}

		// ###############
		// Get the next moves to go to
		var nextMoves []string

		// For each block in the sequence
		for _, block := range newC.Sequence {
			// Find the go to moves
			tmp, ok := moves.Load(block)
			if !ok {
				panic("Error loading movetoblocks")
			}
			potentialMoves, ok := tmp.([]string)
			if !ok {
				panic("Error converting moves")
			}

			for _, potentialMove := range potentialMoves {
				// Check if the move to block is not already in the sequence or in the next moves (i.e. duplicate)
				if !stringInSlice(potentialMove, newC.Sequence) || !stringInSlice(potentialMove, nextMoves) || !stringInSlice(potentialMove, newC.NoMoves) {
					nextMoves = append(nextMoves, potentialMove)
				}
			}
		}

		newC.Moves = nextMoves

		// ##############

		// Load the centre for the block
		tmp, ok = cents.Load(blockToAdd)
		if !ok {
			panic("Can't find centre")
		}
		u, ok := tmp.([3]float64)
		if !ok {
			panic("Centre conversion error")
		}
		// Get the pairwise distance from the previouse combination
		newC.Dist = c.Dist
		// Now add the pairwise distances of the new block to all the previous blocks
		for _, pb := range c.Sequence {
			tmp, ok = cents.Load(pb)
			if !ok {
				panic("Can't find centre")
			}
			v, ok := tmp.([3]float64)
			if !ok {
				panic("Centre conversion error")
			}

			first := math.Pow(u[0]-v[0], 2)
			second := math.Pow(u[1]-v[1], 2)
			third := math.Pow(u[2]-v[2], 2)

			// log.Printf("%v %v %v", first, second, third)

			// Sum the distances
			newC.Dist += first + second + third
		}

		// #########
		// log.Printf("%v %v %v", blockToAdd, len(newC.Moves), newC.Dist)

		// Take the previous number of paths from the last combination
		newC.Paths = c.Paths

		newCKey := fmt.Sprintf("%02d_%v_%v", len(newC.Sequence), len(newC.Moves), newC.Dist)

		// Check if it arleady exists or store a new combination
		e, exists := cMap.Load(newCKey)
		if exists {
			update, ok := e.(Combination)
			if !ok {
				panic("Error converting combination")
			}
			update.Paths = update.Paths + newC.Paths
			cMap.Store(newCKey, update)
		} else {
			cMap.Store(newCKey, newC)
		}
	}

	return
}
