package main

import (
	"log"

	"github.com/davecgh/go-spew/spew"
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
	moves := moveMap()
	nomoves := noMoveMap()

	move, ok := moves.Load("30_30_30")
	move = move.([]string)
	if !ok {
		panic("Error: Centre key not found")
	}
	nomove, ok := nomoves.Load("30_30_30")
	nomove = nomove.([]string)
	if !ok {
		panic("Error: Centre key not found")
	}
	spew.Dump(move)
	spew.Dump(nomove)
}
