package main

import "strings"

type Point struct {
	current string
	left    string
	right   string
}

func NewPoint(input string) Point {
	return parseFromInput(input)
}

func parseFromInput(input string) Point {
	p := Point{}
	cs := strings.Split(input, "=")
	p.current = strings.Trim(cs[0], " ")

	directions := strings.Split(cs[1], ",")
	p.left = directions[0][2:]
	p.right = directions[1][1 : len(directions[1])-1]

	return p
}
