package algo

import (
	"math/rand"
	"time"
)

const NumIterations = 10000

func MonteCarloPIA(selections *Selections, ingester *Ingester) map[string]int {
	destinations := generateEmptyRankingMap()

	for range NumIterations {
		output := Pia(selections, ingester)
		destinations[output]++
	}

	return destinations
}

func Pia(selections *Selections, ingester *Ingester) string {
	numStudents := len(selections.Rankings)
	order := generateRandomOrder(numStudents)
	allocatedLocations := generateEmptyRankingMap()
	allocatedStudents := make([]int, 0, numStudents)
	rand.Seed(time.Now().UnixNano())
	ownRanking := rand.Intn(numStudents)

	for _, studentNum := range order {
		studentRanking := selections.Rankings[studentNum]
		if studentNum == ownRanking {
			studentRanking = ingester.Ranking
		}

		firstChoice := getChoice(studentRanking, 1)

		if allocatedLocations[firstChoice] < ingester.AvailablePositions[firstChoice] {
			allocatedLocations[firstChoice]++
			allocatedStudents = append(allocatedStudents, studentNum)
			if studentNum == ownRanking {
				return firstChoice
			}
		}
	}

	for len(allocatedStudents) != numStudents {
		for _, studentNum := range order {
			if containsInt(allocatedStudents, studentNum) {
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
					allocatedStudents = append(allocatedStudents, studentNum)
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

func containsInt(slice []int, element int) bool {
	for _, v := range slice {
		if v == element {
			return true
		}
	}
	return false
}

func getChoice(ranking map[string]int, val int) string {
	for k, v := range ranking {
		if v == val {
			return k
		}
	}

	return "MISSING"
}

func generateRandomOrder(length int) []int {
	slice := make([]int, length)

	for i := range length {
		slice[i] = i
	}

	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(length, func(i, j int) {
		slice[i], slice[j] = slice[j], slice[i]
	})

	return slice
}
