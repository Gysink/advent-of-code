package main

import (
	"advent-of-code/pkg/util"
	"strings"
)

type Hand struct {
	cards []*Card
	bid   int
	value string
	score int
}

func NewHand(input string) *Hand {
	return parseFromInput(input)
}

func parseFromInput(input string) *Hand {
	hand := &Hand{}
	splitted := strings.Split(input, " ")
	hand.bid = util.StringToInt(splitted[1])
	hand.value = splitted[0]

	for _, c := range splitted[0] {
		hand.cards = append(hand.cards, NewCard(string(c)))
	}
	hand.score = getScore(hand.cards, hand.value)
	return hand
}

func getScore(cards []*Card, val string) int {
	score := 0
	if isXOfAKind(cards, 5, val) {
		score += 6
	} else if isXOfAKind(cards, 4, val) {
		score += 5
	} else if isFullHouse(cards, val) {
		score += 4
	} else if isXOfAKind(cards, 3, val) {
		score += 3
	} else if isTwoPair(cards, val) {
		score += 2
	} else if isXOfAKind(cards, 2, val) {
		score += 1
	}
	return score
}

func isFullHouse(cards []*Card, val string) bool {
	counter := countSameCards(cards, val)
	has2Pair := false
	has3OfAKind := false
	hasSecond3OfAKind := false
	for _, v := range counter {
		if v == 3 && !has3OfAKind {
			has3OfAKind = true
		} else if v == 3 && has3OfAKind {
			hasSecond3OfAKind = true
		}
		if v == 2 {
			has2Pair = true
		}
	}
	return len(counter) == 2 && ((has3OfAKind && has2Pair) || (has3OfAKind && hasSecond3OfAKind))
}

func isTwoPair(cards []*Card, val string) bool {
	counter := countSameCards(cards, val)
	count := 0
	for _, v := range counter {
		if v == 2 {
			count++
		}
	}
	return count == 2
}

func isXOfAKind(cards []*Card, x int, val string) bool {
	counter := countSameCards(cards, val)
	for _, v := range counter {
		if v == x {
			return true
		}
	}
	return false
}

func countSameCards(cards []*Card, val string) map[string]int {
	counter := map[string]int{}
	for _, card := range cards {
		count := strings.Count(val, card.value)
		if card.value != "J" {
			countJ := strings.Count(val, "J")
			count += countJ
		}
		if count > 1 {
			counter[card.value] = count
		}
	}
	return counter
}
