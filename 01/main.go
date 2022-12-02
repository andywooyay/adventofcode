package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
)

func main() {
	b, _ := ioutil.ReadFile("01/test.txt")
	totals := getElfCalorieTotals(b)
	sort.Slice(totals, func(i, j int) bool {
		return totals[i] > totals[j]
	})

	log.Println("Total:", totals[0])
	log.Println("Top three:", totals[0]+totals[1]+totals[2])
}

func getElfCalorieTotals(input []byte) []int {
	var elves []int

	chunks := bytes.Split(input, []byte("\n\n"))
	for _, v := range chunks {
		var elf int
		items := bytes.Split(v, []byte("\n"))
		for _, calories := range items {
			itemCals, _ := strconv.ParseInt(string(calories), 10, 64)
			elf += int(itemCals)
		}

		elves = append(elves, elf)
	}

	return elves
}
