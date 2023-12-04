package main

import (
	"advent-of-code/pkg/util"
	"strings"
)

type Card struct {
	name           string
	winningNumbers []int
	cardNumbers    []int
}

func NewCard(input string) *Card {
	c := &Card{}
	if input != "" {
		c.parseFromInput(input)
	}
	return c
}

func (c *Card) getResult() int {
	winning := c.countWinningNumbers()
	if winning == 0 {
		return 0
	}
	return calcResult(winning)
}

func (c *Card) countWinningNumbers() int {
	count := 0
	for _, w := range c.winningNumbers {
		for _, cn := range c.cardNumbers {
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

func (c *Card) parseFromInput(input string) {
	doublePoint := strings.Split(input, ":")
	c.name = doublePoint[0]
	splittedInput := strings.Split(doublePoint[1], "|")
	c.winningNumbers = toIntArray(splittedInput[0])
	c.cardNumbers = toIntArray(splittedInput[1])
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
