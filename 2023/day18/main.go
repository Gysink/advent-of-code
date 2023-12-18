package main

import (
	"advent-of-code/pkg/util"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	res := solve1(string(data))
	log.Println("result for day18:", res)

	res2 := solve2(string(data))
	log.Println("result for day18 part 2:", res2)
}

func solve1(input string) int {
	trenches, total := parseInput(input)
	a := shoelace(trenches)
	tLen := float64(total)
	result := a - (tLen / 2) + 1
	result = result + tLen
	return int(result)
}

type Trench struct {
	x, y int
}

var dirs = map[string]Trench{
	"U": {x: -1, y: 0},
	"D": {x: 1, y: 0},
	"R": {x: 0, y: 1},
	"L": {x: 0, y: -1},
}

func parseInput(input string) ([]Trench, int) {
	var trenches []Trench
	trenches = append(trenches, Trench{x: 0, y: 0})
	lines := util.SplitByLine(input)

	total := 0
	for _, l := range lines {
		spl := strings.Split(l, " ")
		dir := spl[0]
		steps := util.StringToInt(spl[1])

		total += steps
		dr := dirs[dir].x
		dc := dirs[dir].y

		lX := trenches[len(trenches)-1].x
		lY := trenches[len(trenches)-1].y

		trenches = append(trenches, Trench{x: lX + (dr * steps), y: lY + (dc * steps)})
	}
	return trenches, total
}

func shoelace(trenches []Trench) float64 {
	result := float64(0)
	for i, t := range trenches {
		prevY := float64(0)
		nextY := float64(0)
		if i != 0 {
			prevY = float64(trenches[i-1].y)
		} else {
			prevY = float64(trenches[len(trenches)-1].y)
		}
		if i != len(trenches)-1 {
			nextY = float64(trenches[i+1].y)
		} else {
			nextY = float64(trenches[0].y)
		}
		result += float64(t.x) * (prevY - nextY)
	}
	return math.Abs(result) / 2
}

// Part 2

func solve2(input string) int {
	trenches, total := parseInput2(input)
	a := shoelace(trenches)
	tLen := float64(total)
	result := a - (tLen / 2) + 1
	result = result + tLen
	return int(result)
}

var dirs2 = map[int]Trench{
	0: {x: 0, y: 1},
	1: {x: 1, y: 0},
	2: {x: 0, y: -1},
	3: {x: -1, y: 0},
}

func parseInput2(input string) ([]Trench, int) {
	var trenches []Trench
	trenches = append(trenches, Trench{x: 0, y: 0})
	lines := util.SplitByLine(input)

	total := 0
	for _, l := range lines {
		spl := strings.Split(l, " ")
		hexF := spl[2][2:]
		hex := hexF[:len(hexF)-2]
		dir := util.StringToInt(hexF[5 : len(hexF)-1])
		n, err := strconv.ParseUint(hex, 16, 64)
		if err != nil {
			panic(err)
		}

		steps := int(n)
		total += steps
		dr := dirs2[dir].x
		dc := dirs2[dir].y

		lX := trenches[len(trenches)-1].x
		lY := trenches[len(trenches)-1].y

		trenches = append(trenches, Trench{x: lX + (dr * steps), y: lY + (dc * steps)})
	}
	return trenches, total
}
