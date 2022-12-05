package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func main() {
	b, _ := ioutil.ReadFile("05/test.txt")
	result1, result2 := getResults(b)

	log.Println(result1, result2)
}

func getResults(b []byte) (string, string) {
	parts := bytes.Split(b, []byte("\n\n"))

	instructions := parseInstructions(parts[1])

	crates1 := parseStack(parts[0])
	crates2 := parseStack(parts[0])

	for _, v := range instructions {
		v.apply9000(crates1)
		v.apply9001(crates2)
	}

	return crates1.topString(), crates2.topString()
}

func parseStack(stacks []byte) stack {
	rows := strings.Split(string(stacks), "\n")
	if len(rows) == 0 {
		return nil
	}

	stackMap := make(map[int][]string)
	indexes := getColumnIndexes(rows[len(rows)-1])

	for _, v := range rows[:len(rows)-1] {
		var rowCols []string
		for i := 0; i+3 <= len(v); i += 4 {
			crate := strings.TrimSpace(v[i : i+3])
			rowCols = append(rowCols, crate)
		}

		for k := range rowCols {
			if rowCols[k] == "" {
				continue
			}
			stackMap[indexes[k]] = append(stackMap[indexes[k]], rowCols[k])
		}
	}

	return stackMap
}

func getColumnIndexes(row string) []int {
	columnNames := strings.Fields(row)

	var indexes []int
	for _, v := range columnNames {
		intval, _ := strconv.ParseInt(v, 10, 64)
		indexes = append(indexes, int(intval))
	}

	return indexes
}

var pattern = regexp.MustCompile(`move ([0-9]+) from ([0-9]+) to ([0-9]+)`)

func parseInstructions(instructions []byte) []instruction {
	i := bytes.Split(instructions, []byte("\n"))
	var ins []instruction
	for _, v := range i {
		matches := pattern.FindSubmatch(v)
		if len(matches) != 4 {
			panic(fmt.Sprintf("can't parse instruction: %s", v))
		}

		move, _ := strconv.ParseInt(string(matches[1]), 10, 64)
		from, _ := strconv.ParseInt(string(matches[2]), 10, 64)
		to, _ := strconv.ParseInt(string(matches[3]), 10, 64)
		ins = append(ins, instruction{
			move: int(move),
			from: int(from),
			to:   int(to),
		})
	}

	return ins
}

type instruction struct {
	move int
	from int
	to   int
}

func (ins instruction) apply9000(s stack) {
	for i := 0; i < ins.move; i++ {
		crate := s[ins.from][0]
		s[ins.from] = s[ins.from][1:]
		s[ins.to] = append([]string{crate}, s[ins.to]...)
	}
}

func (ins instruction) apply9001(s stack) {
	var crates []string
	crates = append(crates, s[ins.from][0:ins.move]...)
	s[ins.from] = s[ins.from][ins.move:]
	s[ins.to] = append(crates, s[ins.to]...)
}

type stack map[int][]string

func (s stack) top() []string {
	var keys []int
	for col := range s {
		keys = append(keys, col)
	}

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	var top []string
	for _, col := range keys {
		if len(s[col]) > 0 {
			top = append(top, s[col][0])
		}
	}

	return top
}

func (s stack) topString() string {
	top := strings.Join(s.top(), "")
	top = strings.ReplaceAll(top, "[", "")
	top = strings.ReplaceAll(top, "]", "")

	return top
}
