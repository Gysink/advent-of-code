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
	log.Println("result for day10:", res)

	res2 := solve2(string(data))
	log.Println("result for day10 part 2:", res2)
}

func solve1(input string) int {
	points := parsePoints(input)
	res, _ := walk(points)
	return res / 2
}

func parsePoints(input string) [][]Point {
	byLine := util.SplitByLine(input)
	points := make([][]Point, len(byLine))
	for x, l := range byLine {
		points[x] = make([]Point, len(l))
		for y, c := range l {
			points[x][y] = NewPoint(string(c), x, y)
		}
	}
	return points
}

func walk(points [][]Point) (int, []Point) {
	var path []Point
	counter := 0
	start := findStart(points)

	prev := start
	current := start.next(points, Point{})
	current.viewed = true
	counter++

	path = append(path, start)
	path = append(path, current)
	for current.symbol != "S" {
		t := current
		current = current.next(points, prev)
		path = append(path, current)
		prev = t
		counter++
	}
	return counter, path
}

func findStart(points [][]Point) Point {
	for _, l := range points {
		for _, p := range l {
			if p.symbol == "S" {
				return p
			}
		}
	}
	log.Fatalln("no start found")
	return Point{}
}

// Part 2

func solve2(input string) int {
	points := parsePoints(input)
	_, path := walk(points)
	start := findStart(points)
	points[start.x][start.y].symbol = convertStart(path)
	return countInterior(points, seen(path))
}

type Location struct {
	X, Y int
}

func countInterior(rows [][]Point, loop map[Location]Point) int {
	var count int
	for i := 1; i < len(rows)-1; i++ {
		var in bool
		var lastVertex string
		for j := 0; j < len(rows[i])-1; j++ {
			if _, ok := loop[Location{i, j}]; ok {
				switch rows[i][j].symbol {
				case "|":
					in = !in
					lastVertex = ""
				case "L", "F":
					lastVertex = rows[i][j].symbol
				case "J":
					if lastVertex == "F" {
						in = !in
					}
					lastVertex = ""
				case "7":
					if lastVertex == "L" {
						in = !in
					}
					lastVertex = ""
				}
			} else {
				lastVertex = ""
				if in {
					count++
				}
			}
		}
	}
	return count
}

func convertStart(path []Point) string {
	first := path[1]
	last := path[len(path)-2] // the final node is the start node, so the second to last node is actually last

	k := Location{first.x - last.x, first.y - last.y}
	switch k {
	case Location{2, 0}, Location{-2, 0}:
		return "|"
	case Location{0, 2}, Location{0, -2}:
		return "-"
	case Location{-1, -1}:
		return "L"
	case Location{1, -1}:
		return "J"
	case Location{1, 1}:
		return "7"
	case Location{-1, 1}:
		return "F"
	}
	return ""
}

func seen(points []Point) map[Location]Point {
	m := make(map[Location]Point)
	for _, p := range points {
		m[Location{p.x, p.y}] = p
	}
	return m
}
