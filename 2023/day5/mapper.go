package main

import (
	"advent-of-code/pkg/util"
	"fmt"
	"strings"
)

type Mapper struct {
	name     string
	mappings []*Mapping
}

func NewMapper(input string) *Mapper {
	if input == "" {
		return &Mapper{}
	}

	return parseFromInput(input)
}

func parseFromInput(input string) *Mapper {
	m := &Mapper{}
	lines := strings.Split(input, LineBreak)

	for i, l := range lines {
		if l == "" {
			continue
		}
		if i == 0 {
			m.name = l
			continue
		}

		values := strings.Split(l, " ")
		mapping := &Mapping{
			destination: util.StringToInt(values[0]),
			source:      util.StringToInt(values[1]),
			length:      util.StringToInt(values[2]),
		}
		m.mappings = append(m.mappings, mapping)
	}
	return m
}

type Mapping struct {
	destination int
	source      int
	length      int
}

func (mapping *Mapping) String() string {
	return fmt.Sprintf("des=%d, src=%d, len=%d", mapping.destination, mapping.source, mapping.length)
}

func (mapping *Mapping) mapInput(input int) (int, bool) {
	result := 0
	if input >= mapping.source && input < (mapping.source+mapping.length) {
		diff := mapping.destination - mapping.source
		result = input + diff
	}
	if result == 0 {
		return input, false
	}
	return result, true
}
