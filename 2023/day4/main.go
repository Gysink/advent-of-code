package main

import (
	"advent-of-code/pkg/util"
	"log"
	"os"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	res := calcPoints(string(data))
	log.Println("the result (points) is:", res)
}

func calcPoints(input string) int {
	inputByLine := util.SplitByLine(input)

	var sum int
	for _, l := range inputByLine {
		card := NewCard(l)
		sum = sum + card.getResult()
	}
	return sum
}

// part 2

func calcTotalCards(input string) int {
	inputByLine := util.SplitByLine(input)

	var origCards []*Card
	for _, l := range inputByLine {
		origCards = append(origCards, NewCard(l))
	}

	allCards := calcCardsAndCopies(origCards)
	return len(allCards)
}

func calcCardsAndCopies(origCards []*Card) []*Card {
	cards := origCards
	for i, c := range cards {
		if c.countWinningNumbers() > 0 {
			copies := getCopies(i+1, c, origCards)
			//cards = append(cards, copies...)
			cards = append(cards, calcCardsAndCopies(copies)...)
		}
	}
	return cards
}

func getCopies(i int, card *Card, cards []*Card) []*Card {
	return cards[i : i+card.countWinningNumbers()]
}
