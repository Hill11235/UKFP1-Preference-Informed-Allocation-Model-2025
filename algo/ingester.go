package algo

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

const (
	rankingPath            = "./data/ranking-input.json"
	ratiosPath             = "./data/competition-ratios.csv"
	availablePositionsPath = "./data/available-positions.csv"
)

type Ingester struct {
	Ranking            map[string]int
	Ratios             map[string]float32
	AvailablePositions map[string]int
}

func NewIngester() *Ingester {
	fmt.Println("Ingesting data...")
	ing := Ingester{
		Ranking:            generateRanking(),
		Ratios:             generateRatios(),
		AvailablePositions: generatePositions(),
	}

	return &ing
}

func generateRanking() map[string]int {
	byteValue, err := os.ReadFile(rankingPath)
	if err != nil {
		log.Fatal(err)
	}

	var ranking map[string]int

	err = json.Unmarshal(byteValue, &ranking)
	if err != nil {
		log.Fatal(err)
	}

	return ranking
}

func generateRatios() map[string]float32 {
	ratios := map[string]float32{}

	file, err := os.Open(ratiosPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break // End of file reached
		}
		if err != nil {
			log.Fatal(err)
		}

		var totalRatio float32 = 0.0
		var count float32 = 0.0
		for _, value := range record[1:] {
			var num float32
			if value == "-" {
				num = 0.0
			} else {
				f64, _ := strconv.ParseFloat(value, 32)
				num = float32(f64)
				count++
			}
			totalRatio += num
		}

		ratios[record[0]] = totalRatio / count
	}

	return ratios
}

func generatePositions() map[string]int {
	positions := map[string]int{}

	file, err := os.Open(availablePositionsPath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break // End of file reached
		}
		if err != nil {
			log.Fatal(err)
		}
		pos, _ := strconv.Atoi(record[1])

		positions[record[0]] = pos
	}

	return positions
}
