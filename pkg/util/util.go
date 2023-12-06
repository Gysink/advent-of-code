package util

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func SplitByLine(input string) []string {
	var res []string
	splitted := strings.Split(input, fmt.Sprintln())
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

func StringsToIntArray(strings []string) []int {
	var ints []int
	for _, s := range strings {
		if s != "" {
			ints = append(ints, StringToInt(s))
		}
	}
	return ints
}
