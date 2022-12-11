package main

import (
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

func main() {
	b, _ := ioutil.ReadFile("08/test.txt")
	grid := loadGrid(b)

	log.Println("Visible:", grid.Visible())
	log.Println("Highest scenic score:", grid.HighestScenicScore())
}

func loadGrid(b []byte) Grid {
	var trees []Tree
	rows := strings.Split(string(b), "\n")
	for y, row := range rows {
		for x, tree := range row {
			size, err := strconv.ParseInt(string(tree), 10, 64)
			if err != nil {
				panic("unable to parse tree size")
			}

			trees = append(trees, Tree{
				Size: int(size),
				X:    x,
				Y:    y,
			})
		}
	}

	grid := NewGrid(trees)

	return grid
}

type Tree struct {
	Size int
	X    int
	Y    int
}

func (t Tree) VisibleBehind(trees []Tree) bool {
	for _, v := range trees {
		if v.Size >= t.Size {
			return false
		}
	}

	return true
}

type Grid struct {
	width  int
	height int
	trees  map[[2]int]Tree
}

func NewGrid(trees []Tree) Grid {
	var height, width int
	t := make(map[[2]int]Tree, len(trees))

	for _, v := range trees {
		if v.X > width {
			width = v.X
		}
		if v.Y > height {
			height = v.Y
		}
		t[[2]int{v.X, v.Y}] = v
	}

	return Grid{
		width:  width,
		height: height,
		trees:  t,
	}
}

func (g Grid) IsEdge(x, y int) bool {
	return x == 0 || y == 0 || x == g.width || y == g.height
}

func (g Grid) ViewsFrom(x, y int) [][]Tree {
	directions := make([][]Tree, 4)

	for i := x - 1; i >= 0; i-- {
		directions[0] = append(directions[0], g.trees[[2]int{i, y}])
	}
	for i := x + 1; i <= g.width; i++ {
		directions[1] = append(directions[1], g.trees[[2]int{i, y}])
	}
	for i := y - 1; i >= 0; i-- {
		directions[2] = append(directions[2], g.trees[[2]int{x, i}])
	}
	for i := y + 1; i <= g.height; i++ {
		directions[3] = append(directions[3], g.trees[[2]int{x, i}])
	}

	return directions
}

func (g Grid) VisibleAt(x, y int) bool {
	if g.IsEdge(x, y) {
		return true
	}

	target, ok := g.trees[[2]int{x, y}]
	if !ok {
		panic("unknown tree")
	}

	directions := g.ViewsFrom(x, y)
	for _, direction := range directions {
		if target.VisibleBehind(direction) {
			return true
		}
	}

	return false
}

func (g Grid) Visible() int {
	var visible int

	for _, v := range g.trees {
		if g.VisibleAt(v.X, v.Y) {
			visible++
		}
	}

	return visible
}

func (g Grid) ScenicScore(x, y int) int {
	target, ok := g.trees[[2]int{x, y}]
	if !ok {
		panic("unknown tree")
	}

	directions := g.ViewsFrom(x, y)
	scores := make([]int, 4)
	for k, v := range directions {
		for _, tree := range v {
			scores[k]++
			if tree.Size >= target.Size {
				break
			}
		}
	}

	score := 1
	for _, v := range scores {
		score *= v
	}

	return score
}

func (g Grid) HighestScenicScore() int {
	var scores []int
	for _, v := range g.trees {
		scores = append(scores, g.ScenicScore(v.X, v.Y))
	}

	sort.Slice(scores, func(i, j int) bool {
		return scores[i] > scores[j]
	})

	if len(scores) == 0 {
		return 0
	}

	return scores[0]
}
