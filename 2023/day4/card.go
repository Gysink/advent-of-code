package main

import (
	"advent-of-code/pkg/util"
	"log"
	"strings"
)

type Card struct {
	id   int
	name string
	wins int
}

func NewCard(input string) *Card {
	c := &Card{}
	if input != "" {
		c.parseFromInput(input)
	}
	return c
}

func (c *Card) getScore() int {
	if c.wins == 0 {
		return 0
	}
	return calcResult(c.wins)
}

func (c *Card) parseFromInput(input string) {
	doublePoint := strings.Split(input, ":")
	c.name = doublePoint[0]
	c.id = parseId(doublePoint[0])

	wn, cn := parseNumbers(doublePoint[1])
	c.wins = countWinningNumbers(wn, cn)
}

func countWinningNumbers(winningNumbers []int, cardNumbers []int) int {
	count := 0
	for _, w := range winningNumbers {
		for _, cn := range cardNumbers {
			if w == cn {
				count++
			}
		}
	}
	return count
}

func calcResult(winning int) int {
	result := 1
	for i := 0; i < (winning - 1); i++ {
		result = result * 2
	}
	return result
}

func parseNumbers(input string) ([]int, []int) {
	splittedInput := strings.Split(input, "|")
	winningNumbers := toIntArray(splittedInput[0])
	cardNumbers := toIntArray(splittedInput[1])
	return winningNumbers, cardNumbers
}

func parseId(input string) int {
	splitted := strings.Split(input, " ")
	splitted = splitted[1:]
	for _, i := range splitted {
		if i != "" {
			return util.StringToInt(i)
		}
	}
	log.Fatal("no id found")
	return 0
}

func toIntArray(input string) []int {
	var res []int
	for _, n := range strings.Split(input, " ") {
		if n != "" {
			res = append(res, util.StringToInt(n))
		}
	}
	return res
}
