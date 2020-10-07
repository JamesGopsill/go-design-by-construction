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
	Moves    []string
	NoMoves  []string
	Dist     float64
	Paths    int64
}

func main() {
	log.Println("2x4 Lego Construction")

	startPos := "30_30_30_0"

	log.Println("Creating Maps")
	moves := moveMap()
	// spew.Dump(moves.Load("30_30_30_0"))
	nomoves := noMoveMap()
	// spew.Dump(nomoves.Load("30_30_30_0"))
	cents := centres()

	// #####################

	// Creating the initial Combination
	var start Combination

	tmp, ok := moves.Load("30_30_30_0")
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

	start.Sequence = append(start.Sequence, "30_30_30_0")
	start.Moves = append(start.Moves, move...)
	start.NoMoves = append(start.NoMoves, nomove...)
	start.Dist = 0.
	start.Paths = 1

	// spew.Dump(start)

	startKey := fmt.Sprintf("01_%v_%v_%v", len(start.Moves), start.Dist, 0)

	// Add the first combination to the sync map
	var cMap sync.Map
	cMap.Store(startKey, start)

	fmt.Printf("n,D,V,M,Rc,Rf,Pmin,Pmax,Pmean,G\n")
	var D float64
	var interfaces float64
	interfaces = 184.
	var interfacesUsed float64
	D = 1.
	interfacesUsed = 0.

	// #################

	// Now run through the different combinations
	for i := 2; i <= 6; i++ {

		// initialise some vars
		currentKeyStartsWith := fmt.Sprintf("%02d", i)
		previousKeyStartsWith := fmt.Sprintf("%02d", i-1)
		var wg sync.WaitGroup

		// For each in the combination sync.Map
		cMap.Range(func(k, v interface{}) bool {

			// If the prefix is of a sequence from the previous combination set, increment a brick
			if strings.HasPrefix(k.(string), previousKeyStartsWith) {

				// fmt.Printf("Key: %v, Paths to Seq: %v \n", k.(string), v.(Combination).Paths)
				// fmt.Printf("Key: %v\n", k.(string))

				wg.Add(1)
				go incrementBrick(v.(Combination), &cMap, &moves, &nomoves, &cents, &wg)

			}

			return true
		})

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
		Pmin, Pmax := minMax(P)
		Pmean := V / float64(len(P))
		Amin := (math.Pow(M, 2) + (V - M)) / 2.
		q, r := divmod(int64(V), int64(M))
		Amax := M * (((float64(q) * (M + 1)) / 2) + float64(r))
		A := calculateA(P)
		G := (A - Amin) / (Amax - Amin)

		// Print line as a csv line
		fmt.Printf("%d,%e,%e,%e,%e,%e,%e,%e,%e,%e\n", i, D, V, M, Rc, Rf, Pmin, Pmax, Pmean, G)

	}
}

//#################################

func calculateA(P []float64) (A float64) {
	sort.Float64s(P)
	var cum float64
	for _, val := range P {
		cum += val
		A += cum
	}
	return
}

func minMax(array []float64) (min float64, max float64) {
	min = array[0]
	max = array[0]
	for _, value := range array {
		if max < value {
			max = value
		}
		if min > value {
			min = value
		}
	}
	return
}

func divmod(numerator, denominator int64) (quotient, remainder int64) {
	quotient = numerator / denominator // integer division, decimals are truncated
	remainder = numerator % denominator
	return
}
