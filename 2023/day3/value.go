package main

import (
	"advent-of-code/pkg/util"
	"strconv"
)

type ValueCord struct {
	value  string
	startX int
	endX   int
	line   int
}

func (c ValueCord) isNum() bool {
	return util.IsNum(c.value)
}

func (c ValueCord) getNumNeighbours(values []ValueCord) []ValueCord {
	var neighbours []ValueCord

	for _, v := range values {
		if v.isNum() {
			if c.line == v.line {
				if c.startX == v.endX {
					neighbours = append(neighbours, v)
					continue
				}
				if c.endX == v.startX {
					neighbours = append(neighbours, v)
					continue
				}
			}
			if c.line+1 == v.line || c.line-1 == v.line {
				if c.startX <= v.endX && c.endX >= v.startX {
					neighbours = append(neighbours, v)
					continue
				}
				if c.startX+1 >= v.endX && c.endX+1 <= v.startX {
					neighbours = append(neighbours, v)
					continue
				}
			}
		}
	}
	return neighbours
}

func (c ValueCord) getNum() int {
	if c.isNum() {
		num, _ := strconv.Atoi(c.value)
		return num
	}
	return 0
}
