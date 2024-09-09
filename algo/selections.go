package algo

// accept ingester and simulate student selections (equal to total of available positions)

type Selections struct {
	Rankings []Simulation
}

type Simulation struct {
	Ranking map[string]int
}

func NewSelections(ingester *Ingester) *Selections {
	// find total number of available places
	numPlaces := 0

	for _, val := range ingester.AvailablePositions {
		numPlaces += val
	}
	return nil
}
