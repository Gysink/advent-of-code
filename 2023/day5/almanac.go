package main

import (
	"advent-of-code/pkg/util"
	"strings"
)

type Almanac struct {
	seeds   []int
	mappers []*Mapper
}

func (a Almanac) getResult() int {
	res := -1
	for _, s := range a.seeds {
		calc := s
		for _, m := range a.mappers {
			isMapped := false
			for _, mapping := range m.mappings {
				calc, isMapped = mapping.mapInput(calc)
				if isMapped {
					break
				}
			}
		}
		if res == -1 {
			res = calc
			continue
		}
		if calc < res {
			res = calc
		}
	}
	return res

}

func NewAlmanac(input string) *Almanac {
	if input == "" {
		return &Almanac{}
	}
	return &Almanac{
		seeds: parseSeeds2(input),
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

func parseSeeds2(input string) []int {
	input = strings.Replace(input, "seeds: ", "", 1)
	var result []int
	numbers := strings.Split(input, " ")
	for i := 0; i < len(numbers); i = i + 2 {
		n := util.StringToInt(numbers[i])
		l := util.StringToInt(numbers[i+1])
		for j := 0; j < l; j++ {
			result = append(result, n+j)
		}
	}
	return result
}
