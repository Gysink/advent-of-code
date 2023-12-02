package util

import (
	"fmt"
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
