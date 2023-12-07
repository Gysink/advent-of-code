package main

import "advent-of-code/pkg/util"

type Card struct {
	value string
	score int
}

func NewCard(input string) *Card {
	return &Card{
		value: input,
		score: parseScore(input),
	}
}

func parseScore(input string) int {
	switch input {
	case "A":
		return 14
	case "K":
		return 13
	case "Q":
		return 12
	case "J":
		return 1
	case "T":
		return 10
	default:
		return util.StringToInt(input)
	}
}
