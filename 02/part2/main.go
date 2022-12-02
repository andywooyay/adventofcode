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

var Results = map[string]Result{
	"X": Lose,
	"Y": Draw,
	"Z": Win,
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

var Losers = map[Shape]Shape{
	Paper:    Rock,
	Scissors: Paper,
	Rock:     Scissors,
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

		result := Results[string(shapes[1])]

		ourShape := getCorrectShape(theirShape, result)

		gameScore := scoreGame(ourShape, theirShape)
		score += gameScore
	}

	return score
}

func getCorrectShape(theirs Shape, result Result) Shape {
	switch result {
	case Draw:
		return theirs
	case Win:
		return Winners[theirs]
	case Lose:
		return Losers[theirs]
	default:
		panic("oh noes")
	}
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
