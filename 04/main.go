package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	b, _ := ioutil.ReadFile("04/test.txt")
	full, partial := getOverlaps(b)

	log.Println("Total full:", full)
	log.Println("Total partial:", partial)
}

func getOverlaps(input []byte) (full, partial int) {
	pairs := bytes.Split(input, []byte("\n"))
	for _, v := range pairs {
		assignments := bytes.Split(v, []byte(`,`))
		range1 := assignmentToRange(string(assignments[0]))
		range2 := assignmentToRange(string(assignments[1]))

		if assignmentsFullyOverlap(range1, range2) {
			full++
		}
		if assignmentsPartiallyOverlap(range1, range2) {
			partial++
		}
	}

	return
}

func assignmentToRange(assignment string) [2]int64 {
	minMax := strings.Split(assignment, "-")
	min, max, err := getMinMax(minMax)
	if err != nil {
		panic(err)
	}

	return [2]int64{min, max}
}

func getMinMax(minMax []string) (int64, int64, error) {
	if len(minMax) != 2 {
		return 0, 0, fmt.Errorf("expected range in format `n-m`")
	}

	min, err := strconv.ParseInt(minMax[0], 10, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid min value: %w", err)
	}

	max, err := strconv.ParseInt(minMax[1], 10, 64)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid max value: %w", err)
	}

	return min, max, nil
}

func assignmentsFullyOverlap(range1, range2 [2]int64) bool {
	if range1[0] <= range2[0] && range1[1] >= range2[1] {
		return true
	}

	if range2[0] <= range1[0] && range2[1] >= range1[1] {
		return true
	}

	return false
}

func assignmentsPartiallyOverlap(range1, range2 [2]int64) bool {
	if range1[0] <= range2[0] && range1[1] >= range2[0] {
		return true
	}

	if range2[0] <= range1[0] && range2[1] >= range1[0] {
		return true
	}

	return false
}
