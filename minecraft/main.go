package main

import (
	"fmt"
	"log"
	"math"
	"sort"
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
	moves := moves()

	log.Println("Creating Centres Map")
	centres := centres()

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

	fmt.Printf("n,D,V,M,Rc,Rf,Pmin,Pmax,Pmean,G\n")
	var D float64
	var interfaces float64
	interfaces = 6.
	var interfacesUsed float64
	D = 1.
	interfacesUsed = 0.

	// Now run through the different combinations
	for i := 2; i <= 10; i++ {

		// initialise some vars
		currentKeyStartsWith := fmt.Sprintf("%02d", i)
		previousKeyStartsWith := fmt.Sprintf("%02d", i-1)
		var wg sync.WaitGroup

		// For each in the combination sync.Map
		cMap.Range(func(k, v interface{}) bool {

			// If the prefix is of a sequence from the previous combination set, increment a brick
			if strings.HasPrefix(k.(string), previousKeyStartsWith) {
				//fmt.Printf("Key: %v, Paths to Seq: %v \n", k.(string), v.(Combination).Paths)
				wg.Add(1)
				// Run Increment Brick Concurrently pointing to all the sync.Maps
				go incrementBrick(v.(Combination), &cMap, &moves, &centres, &wg)
			}

			return true
		})

		// Wait for all the go routines to complete
		wg.Wait()

		// ###############
		// Report

		// Report on the previous combination while incrementing the next
		// fmt.Printf("\n%02d brick combinations\n", i)
		var V float64
		var M float64
		Rc := 0.0
		var P []float64

		cMap.Range(func(k, v interface{}) bool {

			if strings.HasPrefix(k.(string), currentKeyStartsWith) {
				// spew.Dump(v.(Combination).Sequence)
				M++
				V += float64(v.(Combination).Paths)
				P = append(P, float64(v.(Combination).Paths))
			}

			return true
		})

		// D
		//fmt.Printf("Available Interfaces: %f\n", (interfaces*float64(i-1) - interfacesUsed))
		D = D * (interfaces*float64(i-1) - interfacesUsed)
		interfacesUsed += 2

		// Freedom
		Rf := (M / V) - (1. / V)
		Rf = Rf / (1. - (1. / V))

		// Paths
		Pmin, Pmax := minMax(P)
		Pmean := V / float64(len(P))

		// Gini
		Amin := (math.Pow(M, 2) + (V - M)) / 2.
		q, r := divmod(int64(V), int64(M))
		Amax := M * (((float64(q) * (M + 1)) / 2) + float64(r))
		A := calculateA(P)
		G := (A - Amin) / (Amax - Amin)

		// Print line as a csv line
		fmt.Printf("%d,%e,%e,%e,%e,%e,%e,%e,%e,%e\n", i, D, V, M, Rc, Rf, Pmin, Pmax, Pmean, G)

	}
}

func calculateA(P []float64) (A float64) {
	sort.Float64s(P)
	var cum float64
	for _, val := range P {
		cum += val
		A += cum
	}
	return
}
