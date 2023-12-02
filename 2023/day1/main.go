package main

import (
	"advent-of-code/pkg/util"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type found struct {
	i int
	l int
}

var searchables = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9",
	"one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	sum := getSumOfCalibrations(string(data))
	log.Println("the result is:", sum)
}

func getSumOfCalibrations(input string) int {
	inputByLine := util.SplitByLine(input)
	var sum int
	for _, l := range inputByLine {
		cal := getCalibrationOfLine(l)
		sum = sum + cal
	}
	return sum
}

func getCalibrationOfLine(input string) int {
	firstNumString := getFirstNum(input)
	lastNumString := getLastNum(input)
	return stringToInt(fmt.Sprint(firstNumString, lastNumString))
}

func stringToInt(input string) int {
	res, err := strconv.Atoi(input)
	if err != nil {
		log.Fatalln("cannot convert string to int:", input)
	}
	return res
}

func getFirstNum(input string) string {
	lowestFound := found{
		i: 999,
		l: 0,
	}
	for _, s := range searchables {
		index := strings.Index(input, s)
		if index != -1 && index < lowestFound.i {
			lowestFound = found{i: index, l: len(s)}
		}
	}
	result := input[lowestFound.i : lowestFound.i+lowestFound.l]
	if lowestFound.l > 1 {
		return replaceTextWithNumber(result)
	}
	return result
}

func getLastNum(input string) string {
	highestFound := found{
		i: 0,
		l: 0,
	}
	for _, s := range searchables {
		index := strings.LastIndex(input, s)
		if index != -1 && index >= highestFound.i {
			highestFound = found{i: index, l: len(s)}
		}
	}
	result := input[highestFound.i : highestFound.i+highestFound.l]
	if highestFound.l > 1 {
		return replaceTextWithNumber(result)
	}
	return result
}

func replaceTextWithNumber(input string) string {
	switch input {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	default:
		return ""
	}
}
