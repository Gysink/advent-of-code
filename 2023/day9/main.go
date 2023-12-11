package main

import (
	"advent-of-code/pkg/util"
	"log"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	res := solve2(string(data))
	log.Println("result:", res)
}

func solve1(input string) int {
	byLines := util.SplitByLine(input)

	var results []int
	for _, l := range byLines {
		numStrings := strings.Split(l, " ")
		results = append(results, calcNextVal(util.StringsToIntArray(numStrings)))
	}

	sum := 0
	for _, r := range results {
		sum += r
	}
	return sum
}

func calcNextVal(nums []int) int {
	var result int
	var steps [][]int
	steps = append(steps, nums)
	currentNums := nums
	for !allZeros(currentNums) {
		currentNums = getDiffs(currentNums)
		steps = append(steps, currentNums)
	}

	for _, s := range steps {
		result += s[len(s)-1]
	}

	return result
}

func getDiffs(nums []int) []int {
	var diffs []int
	for i := 0; i < len(nums)-1; i++ {
		diff := nums[i+1] - nums[i]
		diffs = append(diffs, diff)
	}
	return diffs
}

func allZeros(steps []int) bool {
	for _, s := range steps {
		if s != 0 {
			return false
		}
	}
	return true
}

// part 2

func solve2(input string) int {
	byLines := util.SplitByLine(input)

	var results []int
	for _, l := range byLines {
		numStrings := strings.Split(l, " ")
		results = append(results, calcNextVal2(util.StringsToIntArray(numStrings)))
	}

	sum := 0
	for _, r := range results {
		sum += r
	}
	return sum
}

func calcNextVal2(nums []int) int {
	var result float64
	var steps [][]int
	steps = append(steps, nums)
	currentNums := nums
	for !allZeros(currentNums) {
		currentNums = getDiffs(currentNums)
		steps = append(steps, currentNums)
	}

	for i := len(steps) - 1; i >= 0; i-- {
		diff := getDiff(result, float64(steps[i][0]))
		result = diff
	}
	return int(result)
}

func getDiff(a float64, b float64) float64 {
	if b <= 0 {
		return b - a
	}

	if a > b {
		return (a - b) * -1
	} else {
		return b - a
	}
}
