package main

import (
	"log"
	"math"
	"os"
	"strings"
)

var LineBreak = "\r\n"

func main() {
	data, err := os.ReadFile("input.txt")
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

func solve2(input string) float64 {
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

	steps := calcLcmSum(getStartPoints(points), points, instructions)
	return steps
}

func calcLcmSum(startPoints []Point, points map[string]Point, instructions []string) float64 {
	var res []int64
	for _, p := range startPoints {
		res = append(res, walkEndsWith(p, points, instructions))
	}

	sum := float64(res[0])
	for i := 1; i < len(res); i = i + 1 {
		sum = lcm(sum, float64(res[i]))
	}
	return sum
}

func walkEndsWith(startPoint Point, points map[string]Point, instructions []string) int64 {
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
		steps += walkEndsWith(currentPoint, points, instructions)
	}
	return steps
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

func lcm(a, b float64) float64 {
	return (a * b) / gcd(a, b)
}

func gcd(a float64, b float64) float64 {
	if b == 0 {
		return a
	} else {
		return gcd(b, math.Mod(a, b))
	}
}
