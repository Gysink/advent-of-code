package main

import (
	"advent-of-code/pkg/util"
	"log"
	"os"
	"strings"
)

//var LineBreak = "\n"

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
	timesStrings := strings.Split(byLine[0], " ")
	distStrings := strings.Split(byLine[1], " ")
	times := util.StringsToIntArray(timesStrings[1:])
	dist := util.StringsToIntArray(distStrings[1:])

	var results []int
	for i, t := range times {
		count := 0
		for j := 1; j < t; j++ {
			mm := t - j
			if (mm * j) > dist[i] {
				count++
			}
		}
		results = append(results, count)
	}

	var sum int
	for _, r := range results {
		if sum == 0 {
			sum = r
			continue
		}
		sum = sum * r
	}

	return sum
}

func solve2(input string) int {
	byLine := strings.Split(input, LineBreak)
	timeVals := strings.ReplaceAll(byLine[0], "Time:", "")
	distVals := strings.ReplaceAll(byLine[1], "Distance:", "")

	time := util.StringToInt(strings.ReplaceAll(timeVals, " ", ""))
	dist := util.StringToInt(strings.ReplaceAll(distVals, " ", ""))

	count := 0
	for j := 1; j < time; j++ {
		mm := time - j
		if (mm * j) > dist {
			count++
		}
	}

	return count
}
