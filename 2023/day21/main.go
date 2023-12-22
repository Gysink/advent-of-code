package main

import (
	"advent-of-code/pkg/util"
	"log"
	"math"
	"os"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	res := solve1(string(data), 64)
	log.Println("result for day21:", res)

	res2 := solve2(string(data), 26501365)
	log.Println("result for day21 part 2:", res2)
}

const (
	plot = iota
	rock
)

type Dir struct {
	x, y int
}

var (
	Up    = Dir{x: -1, y: 0}
	Down  = Dir{x: 1, y: 0}
	Left  = Dir{x: 0, y: -1}
	Right = Dir{x: 0, y: 1}

	Dirs = []Dir{Up, Down, Left, Right}
)

type Point struct {
	x, y int
}

func solve1(input string, steps int) int {
	grid, start := createGrid(input)
	return fill(start.x, start.y, steps, grid)
}

func fill(startX, startY, steps int, grid [][]int) int {
	queue := NewQueue()
	ans := make(map[Point]bool)
	seen := make(map[Point]bool)
	queue.push(NewQueueItem(Point{x: startX, y: startY}, steps))

	for queue.len() > 0 {
		next := queue.pop()
		if next.steps%2 == 0 {
			ans[next.pos] = true
		}
		if next.steps == 0 {
			continue
		}

		for _, d := range Dirs {
			n := Point{x: next.pos.x + d.x, y: next.pos.y + d.y}
			if n.x < 0 || n.x >= len(grid) || n.y < 0 || n.y >= len(grid[0]) {
				continue
			}
			if grid[n.x][n.y] == rock {
				continue
			}
			if _, ok := seen[n]; ok {
				continue
			}
			seen[n] = true
			queue.push(NewQueueItem(n, next.steps-1))
		}
	}
	return len(ans)
}

func createGrid(input string) ([][]int, Point) {
	lines := util.SplitByLine(input)

	start := Point{}
	grid := make([][]int, len(lines))
	for x, l := range lines {
		grid[x] = make([]int, len(l))
		for y, c := range l {
			val := plot
			if c == '#' {
				val = rock
			}
			if c == 'S' {
				start = Point{x: x, y: y}
			}
			grid[x][y] = val
		}
	}
	return grid, start
}

// Part 2

func solve2(input string, steps int) int {
	grid, start := createGrid(input)

	// assert
	if len(grid) != len(grid[0]) {
		log.Fatalln("not square grid")
	}

	size := len(grid)

	// assert
	if start.x != size/2 || start.y != size/2 {
		log.Fatalln("start not in the middle")
	}
	if steps%size != size/2 {
		log.Fatalln("failed assumption")
	}

	gridWidth := steps/size - 1

	odd := int(math.Pow(float64(((gridWidth/2)*2)+1), 2))
	even := int(math.Pow(float64(((gridWidth+1)/2)*2), 2))

	oddPoints := fill(start.x, start.y, (size*2)+1, grid)
	evenPoints := fill(start.x, start.y, size*2, grid)

	cornerTop := fill(size-1, start.y, size-1, grid)
	cornerRight := fill(start.x, 0, size-1, grid)
	cornerBottom := fill(0, start.y, size-1, grid)
	cornerLeft := fill(start.x, size-1, size-1, grid)
	cornerSum := cornerTop + cornerRight + cornerBottom + cornerLeft

	smallTopRight := fill(size-1, 0, size/2-1, grid)
	smallTopLeft := fill(size-1, size-1, size/2-1, grid)
	smallBottomRight := fill(0, 0, size/2-1, grid)
	smallBottomLeft := fill(0, size-1, size/2-1, grid)
	smallSum := smallTopRight + smallTopLeft + smallBottomRight + smallBottomLeft

	largeTopRight := fill(size-1, 0, size*3/2-1, grid)
	largeTopLeft := fill(size-1, size-1, size*3/2-1, grid)
	largeBottomRight := fill(0, 0, size*3/2-1, grid)
	largeBottomLeft := fill(0, size-1, size*3/2-1, grid)
	largeSum := largeTopRight + largeTopLeft + largeBottomRight + largeBottomLeft

	oddSum := odd * oddPoints
	evenSum := even * evenPoints
	smallTotal := (gridWidth + 1) * smallSum
	largeTotal := gridWidth * largeSum
	return oddSum +
		evenSum +
		cornerSum +
		smallTotal +
		largeTotal
}
