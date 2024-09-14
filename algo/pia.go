package algo

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	NumIterations = 10000
	NumWorkers    = 10
)

func MonteCarloPIA(selections *Selections, ingester *Ingester) map[string]int {
	destinations := generateEmptyRankingMap()
	var wg sync.WaitGroup
	fmt.Printf("Running the algorithm %v times and estimating probabilities...", NumIterations)

	results := make(chan string, NumIterations)

	for i := 0; i < NumWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for j := 0; j < NumIterations/NumWorkers; j++ {
				r := rand.New(rand.NewSource(time.Now().UnixNano()))
				output := pia(selections, ingester, r)
				results <- output
			}
		}()
	}

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		destinations[result]++
	}

	return destinations
}

func pia(selections *Selections, ingester *Ingester, r *rand.Rand) string {
	numStudents := len(selections.Rankings)
	order := generateRandomOrder(numStudents, r)
	allocatedLocations := generateEmptyRankingMap()
	allocatedStudents := make(map[int]struct{}, numStudents)
	ownRanking := rand.Intn(numStudents)

	for _, studentNum := range order {
		studentRanking := selections.Rankings[studentNum]
		if studentNum == ownRanking {
			studentRanking = ingester.Ranking
		}

		firstChoice := getChoice(studentRanking, 1)

		if allocatedLocations[firstChoice] < ingester.AvailablePositions[firstChoice] {
			allocatedLocations[firstChoice]++
			allocatedStudents[studentNum] = struct{}{}
			if studentNum == ownRanking {
				return firstChoice
			}
		}
	}

	for len(allocatedStudents) != numStudents {
		for _, studentNum := range order {
			if _, allocated := allocatedStudents[studentNum]; allocated {
				continue
			}
			studentRanking := selections.Rankings[studentNum]
			if studentNum == ownRanking {
				studentRanking = ingester.Ranking
			}

			var choice string

			for i := range len(studentRanking) {
				choice = getChoice(studentRanking, i)

				if allocatedLocations[choice] < ingester.AvailablePositions[choice] {
					allocatedLocations[choice]++
					allocatedStudents[studentNum] = struct{}{}
					if studentNum == ownRanking {
						return choice
					}
					break
				}
			}
		}
	}

	return "Error"
}

func getChoice(ranking map[string]int, val int) string {
	for k, v := range ranking {
		if v == val {
			return k
		}
	}

	return "MISSING"
}

func generateRandomOrder(length int, r *rand.Rand) []int {
	slice := make([]int, length)

	for i := range length {
		slice[i] = i
	}

	r.Shuffle(length, func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})

	return slice
}
