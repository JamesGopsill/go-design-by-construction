package main

import (
	"fmt"
	"log"
	"math"
	"strings"
	"sync"
)

// Combination is a struct containing details of a combination
type Combination struct {
	Sequence []string
	Next     []string
	Dist     float64
	Paths    int64
}

func main() {

	log.Println("Minecraft Construction")

	log.Println("Creating Moves Map")
	moves := Moves()

	log.Println("Creating Centres Map")
	centres := Centres()

	// Creating the initial Combination
	var start Combination
	m, ok := moves.Load("30_30_30")
	if !ok {
		panic("Error: Centre key not found")
	}
	start.Sequence = append(start.Sequence, "30_30_30")
	start.Next = append(start.Next, m.([]string)...)
	start.Dist = 0.
	start.Paths = 1
	startKey := fmt.Sprintf("01_%v_%v", len(start.Next), start.Dist)

	// Add the first combination to the sync map
	var cMap sync.Map
	cMap.Store(startKey, start)

	// Now run through the different combinations
	for i := 2; i <= 14; i++ {

		// Report on the previous combination while incrementing the next
		fmt.Printf("\n%02d brick combinations\n", i-1)

		// initialise some vars
		previousKeyStartsWith := fmt.Sprintf("%02d", i-1)
		var wg sync.WaitGroup
		uniqueCombinations := 0
		var totalPaths int64

		// For each in the combination sync.Map
		cMap.Range(func(k, v interface{}) bool {

			// If the prefix is of a sequence from the previous combination set, increment a brick
			if strings.HasPrefix(k.(string), previousKeyStartsWith) {
				//fmt.Printf("Key: %v, Paths to Seq: %v \n", k.(string), v.(Combination).Paths)
				uniqueCombinations++
				totalPaths += v.(Combination).Paths
				wg.Add(1)
				// Run Increment Brick Concurrently pointing to all the sync.Maps
				go IncrementBrick(v.(Combination), &cMap, &moves, &centres, &wg)
			}

			return true
		})

		// Wait for all the go routines to complete
		wg.Wait()

		// Print out the facts from the brick combination set
		fmt.Printf("Unique Combinations: %v \n", uniqueCombinations)
		fmt.Printf("Paths: %v \n", totalPaths)
	}
}

// IncrementBrick increments the bricks to the combinations and adds them to the cMap
func IncrementBrick(
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

// stringInSlice checks for a string in a slice
func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// Moves creates the moves that one person can make from one brick to another
func Moves() (m sync.Map) {

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

// Centres creates a map of the centre positions for the bricks
func Centres() (m sync.Map) {
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
