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
	log.Println("result for day16:", res)

	res2 := solve2(string(data))
	log.Println("result for day16 part 2:", res2)
}

func solve1(input string) int {
	lines := util.SplitByLine(input)

	grid := make([][]Point, len(lines))
	for x, l := range lines {
		grid[x] = make([]Point, len(l))
		for y, p := range l {
			grid[x][y] = NewPoint(x, y, p)
		}
	}
	beamGrid(grid, beam{x: 0, y: 0, dir: leftToRight})
	return countEnergizedPoints(grid)
}

const (
	leftToRight = iota
	rightToLeft
	upwards
	downwards
)

type beam struct {
	x, y int
	dir  int
}

func beamGrid(grid [][]Point, startBeam beam) {
	rowLength := len(grid)
	colLength := len(grid[0])

	var beams []beam
	beams = append(beams, startBeam)

	for {
		if len(beams) == 0 {
			break
		}
		wasEnergized := grid[beams[0].x][beams[0].y].isEnergized
		grid[beams[0].x][beams[0].y].isEnergized = true

		switch grid[beams[0].x][beams[0].y].value {

		case mirror:
			switch beams[0].dir {
			case leftToRight:
				beams[0].dir = upwards
			case rightToLeft:
				beams[0].dir = downwards
			case upwards:
				beams[0].dir = leftToRight
			case downwards:
				beams[0].dir = rightToLeft
			}

		case mirror2:
			switch beams[0].dir {
			case leftToRight:
				beams[0].dir = downwards
			case rightToLeft:
				beams[0].dir = upwards
			case upwards:
				beams[0].dir = rightToLeft
			case downwards:
				beams[0].dir = leftToRight
			}

		case horizontalSplitter:
			var newDirBeam int
			var newDirNewBeam int
			noMatch := false
			switch beams[0].dir {
			case upwards:
				newDirBeam = rightToLeft
				newDirNewBeam = leftToRight
			case downwards:
				newDirBeam = leftToRight
				newDirNewBeam = rightToLeft
			default:
				noMatch = true
			}
			if !noMatch {
				if !wasEnergized {
					beams[0].dir = newDirBeam
					beams = append(beams, beam{x: beams[0].x, y: beams[0].y, dir: newDirNewBeam})
				} else {
					// we came to an end (loop)
					beams = beams[1:]
					continue
				}
			}

		case verticalSplitter:
			var newDirBeam int
			var newDirNewBeam int
			noMatch := false
			switch beams[0].dir {
			case rightToLeft:
				newDirBeam = upwards
				newDirNewBeam = downwards
			case leftToRight:
				newDirBeam = downwards
				newDirNewBeam = upwards
			default:
				noMatch = true
			}
			if !noMatch {
				if !wasEnergized {
					beams[0].dir = newDirBeam
					beams = append(beams, beam{x: beams[0].x, y: beams[0].y, dir: newDirNewBeam})
				} else {
					// we came to an end (loop)
					beams = beams[1:]
					continue
				}
			}

		case space:
		default:
			// go through
		}

		nX, nY := nextPoint(beams[0].x, beams[0].y, beams[0].dir)
		if nX >= rowLength || nY >= colLength || nX < 0 || nY < 0 {
			if len(beams) <= 1 {
				break
			}
			// we came to an end
			beams = beams[1:]
			continue
		}
		beams[0].x = nX
		beams[0].y = nY
	}

}

func nextPoint(x int, y int, dir int) (int, int) {
	switch dir {
	case leftToRight:
		return x, y + 1
	case rightToLeft:
		return x, y - 1
	case upwards:
		return x - 1, y
	case downwards:
		return x + 1, y
	}
	return -1, -1
}

// Part 2

func solve2(input string) int {
	lines := util.SplitByLine(input)

	grid := make([][]Point, len(lines))
	for x, l := range lines {
		grid[x] = make([]Point, len(l))
		for y, p := range l {
			grid[x][y] = NewPoint(x, y, p)
		}
	}
	highest := 0

	xMax := len(grid) - 1
	yMax := len(grid[0]) - 1

	// x:0, y:0-max
	for i := 0; i < yMax; i++ {
		newGrid := copyGrid(grid)
		beamGrid(newGrid, beam{x: 0, y: i, dir: downwards})
		ep := countEnergizedPoints(newGrid)
		if ep > highest {
			highest = ep
		}
	}

	// x:max, y:0-max
	for i := 0; i < yMax; i++ {
		newGrid := copyGrid(grid)
		beamGrid(newGrid, beam{x: xMax, y: i, dir: upwards})
		ep := countEnergizedPoints(newGrid)
		if ep > highest {
			highest = ep
		}
	}

	// x:0-max, y:0
	for i := 0; i < xMax; i++ {
		newGrid := copyGrid(grid)
		beamGrid(newGrid, beam{x: i, y: 0, dir: leftToRight})
		ep := countEnergizedPoints(newGrid)
		if ep > highest {
			highest = ep
		}
	}

	// x:0-max, y:max
	for i := 0; i < yMax; i++ {
		newGrid := copyGrid(grid)
		beamGrid(newGrid, beam{x: i, y: yMax, dir: rightToLeft})
		ep := countEnergizedPoints(newGrid)
		if ep > highest {
			highest = ep
		}
	}

	return highest
}

func copyGrid(grid [][]Point) [][]Point {
	c := make([][]Point, len(grid))
	for i, r := range grid {
		c[i] = make([]Point, len(r))
		for j, p := range r {
			c[i][j] = p
		}
	}
	return c
}

func countEnergizedPoints(grid [][]Point) int {
	sum := 0
	for _, r := range grid {
		for _, c := range r {
			if c.isEnergized {
				sum++
			}
		}
	}
	return sum
}
