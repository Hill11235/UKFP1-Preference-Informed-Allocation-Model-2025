package main

import (
	"fmt"

	"github.com/Hill11235/deanery-model/algo"
)

func main() {
	// TODO: track how long this whole simulation takes

	ingester := algo.NewIngester()
	selections := algo.NewSelections(ingester)
	fmt.Println(selections)
	// TODO: run pia algo
	// TODO: print the outcome
}
