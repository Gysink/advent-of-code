package main

import (
	"advent-of-code/pkg/util"
	"log"
	"os"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	res := solve1(string(data))
	log.Println("result for day14:", res)

	res2 := solve2(string(data), 1000000000)
	log.Println("result for day14 Part 2:", res2)
}

func solve1(input string) int {
	lineCount := len(util.SplitByLine(input))
	p := NewPlatform(input)
	p.tiltNorth()

	result := 0
	for _, r := range p.getAllRocks() {
		result += 1 * (lineCount - r.x)
	}
	return result
}

// Part 2

type plTouple struct {
	pT Platform
	iT int
}

func solve2(input string, cycles int) int {
	lineCount := len(util.SplitByLine(input))
	p := NewPlatform(input)

	visited := map[string]int{}
	var firstDuplicate *plTouple
	cycle := 0
	index := 0

	for {
		p.Step()
		if firstDuplicate != nil {
			if firstDuplicate.pT.String() == p.String() {
				cycle = index
				break
			}
		} else if i, ok := visited[p.String()]; ok {
			firstDuplicate = &plTouple{pT: p, iT: i}
		} else {
			visited[p.String()] = index + 1
		}
		index++
	}

	first := firstDuplicate.iT
	remainingCycles := (cycles - first) % (cycle - first)
	for i := 1; i < remainingCycles; i++ {
		p.Step()
	}

	result := 0
	for _, r := range p.getAllRocks() {
		result += 1 * (lineCount - r.x)
	}
	return result
}
