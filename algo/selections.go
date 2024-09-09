package algo

import (
	"fmt"
	"sort"
)

type Selections struct {
	Rankings []Simulation
}

type Simulation struct {
	Ranking map[string]int
}

func NewSelections(ingester *Ingester) *Selections {
	numPlaces := 0

	for _, val := range ingester.AvailablePositions {
		numPlaces += val
	}

	// generate first preferences for each sim based on competition ratios
	// subsequent places - to start with use relative popularity minus oversubscribed?
	return nil
}

func generateSimulationsWithFirstChoice(ingester *Ingester) Selections {
	selections := Selections{}
	fmt.Println(ingester.AvailablePositions)

	return selections
}

func completeSubsequentChoices(selections *Selections, ingester *Ingester) {
}

func sortRatios(ratios map[string]float32) []string {
	keys := make([]string, 0, len(ratios))

	for key := range ratios {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return ratios[keys[i]] < ratios[keys[j]]
	})

	return keys
}
