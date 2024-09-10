package algo

import (
	"fmt"
	"math"
	"sort"
)

type Selections struct {
	Rankings []map[string]int
}

func NewSelections(ingester *Ingester) *Selections {
	fmt.Println("Simulating student rankings...")
	selections := generateSimulationsWithFirstChoice(ingester)
	completeSubsequentChoices(&selections, ingester)

	return &selections
}

func generateSimulationsWithFirstChoice(ingester *Ingester) Selections {
	numPlaces := 0

	for _, val := range ingester.AvailablePositions {
		numPlaces += val
	}

	rankings := make([]map[string]int, 0, numPlaces)

	for location, num := range ingester.AvailablePositions {
		numFirstChoices := float32(num) * ingester.Ratios[location]

		for range int(math.Round(float64(numFirstChoices))) {
			emptyRanking := generateEmptyRankingMap()
			emptyRanking[location] = 1
			rankings = append(rankings, emptyRanking)
		}
	}

	selections := Selections{rankings}

	return selections
}

// Critical, will determine probabilities beyond the first choice
// Using relative popularity minus oversubscribed
// This has the Northern Ireland problem, relatively popular first choice, unlikely to be popular beyond first choice
func completeSubsequentChoices(selections *Selections, ingester *Ingester) {
	locationsByPopularity := sortRatios(ingester.Ratios)

	for _, ranking := range selections.Rankings {
		count := 2
		for _, loc := range locationsByPopularity {
			if ranking[loc] == 1 || ingester.Ratios[loc] > 1 {
				continue
			}
			ranking[loc] = count
			count++
		}
		for _, loc := range locationsByPopularity {
			if ranking[loc] != 0 {
				continue
			}
			ranking[loc] = count
			count++
		}
	}
}

func sortRatios(ratios map[string]float32) []string {
	keys := make([]string, 0, len(ratios))

	for key := range ratios {
		keys = append(keys, key)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return ratios[keys[i]] > ratios[keys[j]]
	})

	return keys
}

func generateEmptyRankingMap() map[string]int {
	return map[string]int{
		"East of England":    0,
		"KSS":                0,
		"LNR":                0,
		"London":             0,
		"North West England": 0,
		"Northern Ireland":   0,
		"Northern":           0,
		"Oxford":             0,
		"Peninsula":          0,
		"Scotland":           0,
		"Severn":             0,
		"Trent":              0,
		"Wales":              0,
		"Wessex":             0,
		"West Mids Central":  0,
		"West Mids North":    0,
		"West Mids South":    0,
		"Yorks & Humber":     0,
	}
}
