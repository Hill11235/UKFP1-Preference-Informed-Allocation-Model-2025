package main

import (
	"fmt"

	"github.com/Hill11235/deanery-model/algo"
)

func main() {
	// TODO: track how long this whole simulation takes

	ingester := algo.NewIngester()
	fmt.Println(ingester.AvailablePositions)

	// TODO: create selections
	// TODO: run pia algo
	// TODO: print the outcome
}
