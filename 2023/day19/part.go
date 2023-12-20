package main

import (
	"advent-of-code/pkg/util"
	"strings"
)

type Part struct {
	x, m, a, s int
}

func (p Part) process(workflows map[string]Workflow) string {
	result := "in"
	for result != "A" && result != "R" {
		result = workflows[result].process(p)
	}
	return result
}

func (p Part) getSumRatings() int {
	return p.x + p.m + p.a + p.s
}

func NewPart(input string) Part {
	input = input[1 : len(input)-1]
	sp := strings.Split(input, ",")
	return Part{
		x: util.StringToInt(sp[0][2:]),
		m: util.StringToInt(sp[1][2:]),
		a: util.StringToInt(sp[2][2:]),
		s: util.StringToInt(sp[3][2:]),
	}
}
