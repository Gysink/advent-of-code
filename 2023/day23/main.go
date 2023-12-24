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

	res := solve1(string(data))
	log.Println("result for day23:", res)

	res2 := solve2(string(data))
	log.Println("result for day23 part 2:", res2)
}

func solve1(input string) int {
	grid := parseGrid(input)
	start, end := findStartAndEnd(grid)
	queue := NewQueue()
	queue.push(NewQueueItem(start, 0, make([]Point, 0)))
	foundPaths := make([]int, 0)
	for queue.len() > 0 {
		p := queue.pop()
		if p.pos.isOn(end) {
			foundPaths = append(foundPaths, p.steps)
			continue
		}
		for _, d := range dirs {
			next := Point{x: p.pos.x + d.x, y: p.pos.y + d.y}
			if !grid.isValid(p.pos, next, p.path) {
				continue
			}
			queue.push(NewQueueItem(next, p.steps+1, append(p.path, p.pos)))
		}
	}
	return longest(foundPaths)
}

func longest(paths []int) int {
	l := 0
	for _, p := range paths {
		if p > l {
			l = p
		}
	}
	return l
}

func findStartAndEnd(grid Grid) (Point, Point) {
	start := Point{x: 0}
	end := Point{x: len(grid) - 1}
	for y, c := range grid[0] {
		if c == path {
			start.y = y
		}
	}
	for y, c := range grid[len(grid)-1] {
		if c == path {
			end.y = y
		}
	}
	return start, end
}

type Point struct {
	x, y int
}

type Grid [][]int

func (g Grid) isValid(current Point, next Point, walkedPath []Point) bool {
	if len(walkedPath) > 0 {
		for _, p := range walkedPath {
			if next.x == p.x && next.y == p.y {
				return false
			}
		}
	}
	if next.x <= 0 || next.x >= len(g) {
		return false
	}
	if next.y <= 0 || next.y >= len(g[0]) {
		return false
	}

	switch g[next.x][next.y] {
	case forest:
		return false
	case path:
		return true
	case upSlope:
		return current.x > next.x
	case downSlope:
		return current.x < next.x
	case rightSlope:
		return current.y < next.y
	case leftSlope:
		return current.y > next.y
	}
	log.Fatalln("could not determine valid path")
	return false
}

const (
	path       = '.'
	forest     = '#'
	upSlope    = '^'
	downSlope  = 'v'
	rightSlope = '>'
	leftSlope  = '<'
)

type Dir struct {
	x, y int
}

var (
	up    = Dir{x: -1, y: 0}
	down  = Dir{x: 1, y: 0}
	right = Dir{x: 0, y: 1}
	left  = Dir{x: 0, y: -1}
	dirs  = []Dir{up, down, left, right}
)

func (p Point) isOn(point Point) bool {
	if p.x != point.x {
		return false
	}
	if p.y != point.y {
		return false
	}
	return true
}

func parseGrid(input string) Grid {
	lines := util.SplitByLine(input)
	grid := make(Grid, len(lines))
	for x, r := range lines {
		grid[x] = make([]int, len(r))
		for y, c := range r {
			grid[x][y] = int(c)
		}
	}
	return grid
}

// Part 2

func solve2(input string) int {
	grid := parseGrid(input)
	start, end := findStartAndEnd(grid)

	points := []Point{start, end}
	for x, r := range grid {
		for y, ch := range r {
			if ch == forest {
				continue
			}
			neighbors := 0
			for _, d := range dirs {
				nx := x + d.x
				ny := y + d.y
				if 0 <= nx && nx < len(grid) && 0 <= ny && ny < len(grid[0]) &&
					grid[nx][ny] != forest {
					neighbors++
				}
			}
			if neighbors >= 3 {
				points = append(points, Point{x, y})
			}
		}
	}

	graph := NewGraph(points)
	for _, sp := range points {
		queue := NewQueue()
		queue.push(NewQueueItem(sp, 0, make([]Point, 0)))
		seen := make(map[Point]bool)
		seen[sp] = true
		for queue.len() > 0 {
			item := queue.pop()
			if item.steps != 0 && in(item, points) {
				graph.setSteps(sp, item.pos, item.steps)
				continue
			}
			for _, d := range dirs {
				nx := item.pos.x + d.x
				ny := item.pos.y + d.y
				np := Point{x: nx, y: ny}
				_, ok := seen[np]
				if 0 <= nx && nx < len(grid) && 0 <= ny && ny < len(grid[0]) &&
					grid[nx][ny] != forest && !ok {
					queue.push(NewQueueItem(np, item.steps+1, make([]Point, 0)))
					seen[np] = true
				}
			}
		}
	}
	seen := make(map[Point]bool)
	return dfs(start, end, seen, graph)
}

var seenG = make(map[Point]bool)

func dfs(point Point, end Point, seen map[Point]bool, graph Graph) int {
	if point.isOn(end) {
		return 0
	}
	m := 0
	seenG[point] = true
	for k, s := range graph.get(point) {
		if v, ok := seenG[k]; !ok && !v {
			m = max(m, dfs(k, end, seen, graph)+s)
		}
	}
	delete(seenG, point)
	return m
}

func max(a int, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func in(item QueueItem, points []Point) bool {
	for _, p := range points {
		if p.isOn(item.pos) {
			return true
		}
	}
	return false
}
