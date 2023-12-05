package main

import (
	"advent-of-code/pkg/util"
	"strings"
)

type Almanac struct {
	seeds []int
}

func (a Almanac) getResult() int {
	return 0
}

func NewAlmanac(input string) *Almanac {
	if input == "" {
		return &Almanac{}
	}
	return &Almanac{
		seeds: parseSeeds(input),
	}
}

func parseSeeds(input string) []int {
	input = strings.Replace(input, "seeds: ", "", 1)
	var result []int
	numbers := strings.Split(input, " ")
	for _, n := range numbers {
		result = append(result, util.StringToInt(n))
	}
	return result
}
