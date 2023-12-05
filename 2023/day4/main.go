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
	total := calcTotalCards(string(data))
	log.Println("the total numbers of cards is:", total)
}

func calcPoints(input string) int {
	inputByLine := util.SplitByLine(input)

	var sum int
	for _, l := range inputByLine {
		card := NewCard(l)
		sum = sum + card.getScore()
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

	calc := map[int]int{1: 1}
	score := 0
	for _, c := range origCards {
		id := c.id
		score := c.wins

		if _, ok := calc[id]; !ok {
			calc[id] = 1
		}

		if score == 0 {
			continue
		}

		for i := id + 1; i <= id+score; i++ {
			if _, ok := calc[i]; !ok {
				calc[i] = 1
			}
			calc[i] += calc[id]
		}
	}

	for _, v := range calc {
		score += v
	}

	return score
}
