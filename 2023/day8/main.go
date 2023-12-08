package main

import (
	"log"
	"os"
	"strings"
)

var LineBreak = "\n"

func main() {
	data, err := os.ReadFile("2023/day8/input.txt")
	//data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	res := solve2(string(data))
	log.Println("result:", res)
}

func solve1(input string) int {
	byLine := strings.Split(input, LineBreak)

	instructions := strings.Split(byLine[0], "")

	points := map[string]Point{}
	for _, l := range byLine[1:] {
		if l == LineBreak || l == "" {
			continue
		}
		p := NewPoint(l)
		points[p.current] = p
	}

	steps := walk(points["AAA"], points, instructions)
	return steps
}

func walk(startPoint Point, points map[string]Point, instructions []string) int {
	steps := 0
	currentPoint := startPoint

	for _, in := range instructions {
		if in == "R" {
			currentPoint = points[currentPoint.right]
			steps++
		} else if in == "L" {
			currentPoint = points[currentPoint.left]
			steps++
		}
		if currentPoint.current == "ZZZ" {
			break
		}
	}
	if currentPoint.current != "ZZZ" {
		steps += walk(currentPoint, points, instructions)
	}
	return steps
}

// Part2

func solve2(input string) int64 {
	byLine := strings.Split(input, LineBreak)

	instructions := strings.Split(byLine[0], "")

	points := map[string]Point{}
	for _, l := range byLine[1:] {
		if l == LineBreak || l == "" {
			continue
		}
		p := NewPoint(l)
		points[p.current] = p
	}

	steps := walk3(getStartPoints(points), points, instructions)
	return steps
}

func walk3(startPoints []Point, points map[string]Point, instructions []string) int64 {
	currentPoints := startPoints

	var res []int64
	for _, p := range currentPoints {
		res = append(res, walk22(p, points, instructions))
	}

	sum := res[0]
	for i := 1; i < len(res); i = i + 1 {
		sum = lcm(sum, res[i])
	}

	return sum
}

func walk22(startPoint Point, points map[string]Point, instructions []string) int64 {
	steps := int64(0)
	currentPoint := startPoint

	for _, in := range instructions {
		if in == "R" {
			currentPoint = points[currentPoint.right]
			steps++
		} else if in == "L" {
			currentPoint = points[currentPoint.left]
			steps++
		}
		if currentPoint.current[2:] == "Z" {
			break
		}
	}
	if currentPoint.current[2:] != "Z" {
		steps += walk22(currentPoint, points, instructions)
	}
	return steps
}

func walk2(startPoints []Point, points map[string]Point, instructions []string) int64 {
	var steps int64
	steps = 0
	currentPoints := startPoints

	for !allEndWithZ(currentPoints) {
		for _, in := range instructions {
			if in == "R" {
				for i, cp := range currentPoints {
					currentPoints[i] = points[cp.right]
				}
				steps++
			} else if in == "L" {
				for i, cp := range currentPoints {
					currentPoints[i] = points[cp.left]
				}
				steps++
			}
			if allEndWithZ(currentPoints) {
				break
			}
		}
	}
	//if !allEndWithZ(currentPoints) {
	//	//log.Println("next round:", steps)
	//	steps += walk2(currentPoints, points, instructions)
	//}
	return steps
}

func allEndWithZ(points []Point) bool {
	for _, p := range points {
		if p.current[2:] != "Z" {
			return false
		}
	}
	return true
}

func oneEndsWithZ(points []Point) bool {
	for _, p := range points {
		if p.current[2:] == "Z" {
			return true
		}
	}
	return false
}

func getStartPoints(points map[string]Point) []Point {
	var starts []Point
	for _, p := range points {
		if p.current[2:] == "A" {
			starts = append(starts, p)
		}
	}
	return starts
}

func lcm(temp1 int64, temp2 int64) int64 {
	var lcmnum int64 = 1
	if temp1 > temp2 {
		lcmnum = temp1
	} else {
		lcmnum = temp2
	}

	for {
		if lcmnum%temp1 == 0 && lcmnum%temp2 == 0 {
			return lcmnum
		}
		lcmnum++
	}
}
