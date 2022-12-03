package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	b, _ := ioutil.ReadFile("03/test.txt")

	total := calculateTotal(b)
	log.Println(total)
}

func calculateTotal(input []byte) int {
	rucksacks := bytes.Split(input, []byte("\n"))
	var total int

	for _, v := range rucksacks {
		compartment1, compartment2 := v[:len(v)/2], v[len(v)/2:]

		common := findCommonItems(string(compartment1), string(compartment2))

		for _, item := range common {
			total += getPriority(item)
		}
	}

	return total
}

func findCommonItems(in ...string) []rune {
	common := make(map[rune]struct{})
	if len(in) == 0 {
		return []rune{}
	}

	for _, v := range in[0] {
		if _, ok := common[v]; ok {
			continue
		}

		if !stringsContainRune(in[1:], v) {
			continue
		}

		common[v] = struct{}{}
	}

	var items []rune
	for k := range common {
		items = append(items, k)
	}

	return items
}

func stringsContainRune(in []string, v rune) bool {
	for _, compare := range in {
		if !strings.ContainsRune(compare, v) {
			return false
		}
	}

	return true
}

func getPriority(r rune) int {
	if r >= 97 && r <= 122 {
		return int(r) - 96
	}

	if r >= 65 && r <= 90 {
		return 26 + int(r) - 64
	}

	panic(fmt.Sprintf("cannot prioritise item: %s", string(r)))
}
