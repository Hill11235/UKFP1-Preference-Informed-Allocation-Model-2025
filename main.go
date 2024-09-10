package main

import (
	"fmt"
	"sort"
	"strings"
	"time"

	"github.com/Hill11235/deanery-model/algo"
)

func main() {
	start := time.Now()

	ingester := algo.NewIngester()
	selections := algo.NewSelections(ingester)
	output := algo.MonteCarloPIA(selections, ingester)

	sortAndPrettyPrintMap(output, ingester)

	elapsed := time.Since(start)
	fmt.Printf("\nExecution time: %s\n", elapsed)
}

func sortAndPrettyPrintMap(m map[string]int, ingester *algo.Ingester) {
	type kv struct {
		Key   string
		Value int
	}

	var sortedSlice []kv
	for k, v := range ingester.Ranking {
		sortedSlice = append(sortedSlice, kv{k, v})
	}

	// Sort the slice by value in ascending order
	sort.Slice(sortedSlice, func(i, j int) bool {
		return sortedSlice[i].Value < sortedSlice[j].Value
	})

	// Determine the maximum width for each column
	maxKeyLen := 0
	maxCategoryLen := 15
	for k := range m {
		if len(k) > maxKeyLen {
			maxKeyLen = len(k)
		}
	}

	fmt.Printf("\n\n%-*s | %-*s | Probability\n", maxKeyLen, "Location", maxCategoryLen, "Initial Ranking")
	fmt.Println(strings.Repeat("-", maxKeyLen+maxCategoryLen+10)) // Adjust based on header length

	for idx, val := range sortedSlice {
		prob := float32(m[val.Key]) / algo.NumIterations
		if prob == 0.00 {
			fmt.Printf("%-*s | %-*v | %s\n", maxKeyLen, val.Key, maxCategoryLen, idx+1, "-")
		} else {
			fmt.Printf("%-*s | %-*v | %.2f\n", maxKeyLen, val.Key, maxCategoryLen, idx+1, float32(m[val.Key])/algo.NumIterations)
		}
	}
}
