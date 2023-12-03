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

	sum := calcEnginePartsSum(string(data))
	log.Println("the result is:", sum)

	sum2 := calcEnginePartsSum_Part2(string(data))
	log.Println("the result of part 2 is:", sum2)
}

func calcEnginePartsSum(input string) int {
	values := parseInput(input)
	sum := 0
	for _, v := range values {
		if !v.isNum() {
			neighbours := v.getNumNeighbours(values)
			for _, n := range neighbours {
				sum = sum + n.getNum()
			}
		}
	}
	return sum
}

func calcEnginePartsSum_Part2(input string) int {
	values := parseInput(input)
	sum := 0
	for _, v := range values {
		if !v.isNum() {
			neighbours := v.getNumNeighbours(values)
			if v.value == "*" && len(neighbours) == 2 {
				sum = sum + (neighbours[0].getNum() * neighbours[1].getNum())
			}
		}
	}
	return sum
}

func parseInput(input string) []ValueCord {
	inputByLine := util.SplitByLine(input)
	var values []ValueCord
	for i, l := range inputByLine {
		values = append(values, parseValues(l, i)...)
	}
	return values
}

func parseValues(input string, line int) []ValueCord {
	input = input + "."
	var nums []ValueCord
	var buffer string
	startX := -1
	for i, c := range input {
		if !isAtAnEnd(string(c), buffer) {
			buffer = buffer + string(c)
			if startX == -1 {
				startX = i
			}
		} else {
			if buffer != "" {
				nums = append(nums, ValueCord{
					value:  buffer,
					startX: startX,
					endX:   i,
					line:   line,
				})
				if c == '.' {
					buffer = ""
					startX = -1
				} else {
					buffer = string(c)
					startX = i
				}
			}
		}
	}
	return nums
}

func isAtAnEnd(c, buffer string) bool {
	if c == "." {
		return true
	}
	if buffer != "" {
		if util.IsNum(c) && !util.IsNum(buffer) {
			return true
		}
		if !util.IsNum(c) && util.IsNum(buffer) {
			return true
		}
		if !util.IsNum(c) && !util.IsNum(buffer) {
			return true
		}
	}
	return false
}
