package main

import (
	"bytes"
	"io/ioutil"
	"log"
)

func main() {
	b, _ := ioutil.ReadFile("02/test.txt")

	total := getTotalScore(b)
	log.Println(total)
}

type Shape int

const (
	Rock Shape = iota
	Paper
	Scissors
)

type Result int

const (
	Win Result = iota
	Lose
	Draw
)

var Shapes = map[string]Shape{
	"A": Rock,
	"B": Paper,
	"C": Scissors,
}

var ShapeResponses = map[string]Shape{
	"X": Rock,
	"Y": Paper,
	"Z": Scissors,
}

var BaseScores = map[Shape]int{
	Rock:     1,
	Paper:    2,
	Scissors: 3,
}

var Winners = map[Shape]Shape{
	Rock:     Paper,
	Paper:    Scissors,
	Scissors: Rock,
}

var Scores = map[Result]int{
	Win:  6,
	Draw: 3,
	Lose: 0,
}

func getTotalScore(input []byte) int {
	score := 0
	games := bytes.Split(input, []byte("\n"))
	for _, game := range games {
		if len(game) == 0 {
			continue
		}

		shapes := bytes.Split(game, []byte(` `))
		theirShape := Shapes[string(shapes[0])]
		ourShape := ShapeResponses[string(shapes[1])]

		gameScore := scoreGame(ourShape, theirShape)
		score += gameScore
	}

	return score
}

func scoreGame(ours, theirs Shape) int {
	baseScore, ok := BaseScores[ours]
	if !ok {
		panic("oh noes")
	}

	if theirs == ours {
		return baseScore + Scores[Draw]
	}

	if ours == Winners[theirs] {
		return baseScore + Scores[Win]
	}

	return baseScore + Scores[Lose]
}
