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
	log.Println("result day11:", res)

	res2 := solve2(string(data))
	log.Println("result day11 part 2:", res2)
}

func solve1(input string) int {
	return calc(input, 2)
}

func solve2(input string) int {
	return calc(input, 1_000_000)
}

type galaxy struct {
	x int
	y int
}

type pair struct {
	a, b galaxy
}

func (p pair) distance() int {
	dY := p.b.y - p.a.y
	dX := p.b.x - p.a.x

	if dY < 0 {
		dY = -dY
	}
	if dX < 0 {
		dX = -dX
	}
	return dX + dY
}

func calc(input string, adj int) int {
	rows := util.SplitByLine(input)
	emptyR, emptyC := getEmptyRowsAndCols(rows)

	galaxies := findGalaxies(rows, adj-1, emptyR, emptyC)
	pairs := toPairs(galaxies)

	var sum int
	for _, p := range pairs {
		sum += p.distance()
	}
	return sum
}

func findGalaxies(rows []string, adj int, emptyR []bool, emptyC []bool) []galaxy {
	var result []galaxy
	iCount := 0

	for i, row := range rows {
		if emptyR[i] {
			iCount++
		}
		jCount := 0
		for j := range row {
			if emptyC[j] {
				jCount++
			}
			if row[j] == '#' {
				result = append(result, galaxy{
					x: i + iCount*adj,
					y: j + jCount*adj,
				})
			}
		}
	}
	return result
}

func toPairs(g []galaxy) []pair {
	var out []pair
	for i := 0; i < len(g)-1; i++ {
		for j := i + 1; j < len(g); j++ {
			out = append(out, pair{g[i], g[j]})
		}
	}
	return out
}

func getEmptyRowsAndCols(rows []string) ([]bool, []bool) {
	emptyR := make([]bool, len(rows))
	emptyC := make([]bool, len(rows[0]))
	counts := make([]int, len(rows[0]))
	for i, row := range rows {
		var full bool
		for j := range row {
			if row[j] == '.' {
				counts[j]++
			} else {
				full = true
			}
		}
		if !full {
			emptyR[i] = true
		}
	}
	for i := range counts {
		emptyC[i] = counts[i] == len(rows)
	}
	return emptyR, emptyC
}
