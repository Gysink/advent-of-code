package util

import (
	"log"
	"math"
	"runtime"
	"strconv"
	"strings"
)

func SplitByLine(input string) []string {
	var res []string
	splitted := strings.Split(input, GetLineBreak())
	for _, s := range splitted {
		if s != "" {
			res = append(res, s)
		}
	}
	return res
}

func IsNum(c string) bool {
	_, err := strconv.Atoi(c)
	return err == nil
}

func StringToInt(input string) int {
	res, err := strconv.Atoi(input)
	if err != nil {
		log.Fatalln("cannot convert string to int:", input)
	}
	return res
}

func StringToFloat(input string) float64 {
	res, err := strconv.Atoi(input)
	if err != nil {
		log.Fatalln("cannot convert string to int:", input)
	}
	return float64(res)
}

func StringsToIntArray(strings []string) []int {
	var ints []int
	for _, s := range strings {
		if s != "" {
			ints = append(ints, StringToInt(s))
		}
	}
	return ints
}

func GetLineBreak() string {
	if runtime.GOOS == "windows" {
		return "\r\n"
	}
	return "\n"
}

func CalculateLCM(a, b float64) float64 {
	return (a * b) / CalculateGCD(a, b)
}

func CalculateGCD(a float64, b float64) float64 {
	if b == 0 {
		return a
	} else {
		return CalculateGCD(b, math.Mod(a, b))
	}
}
