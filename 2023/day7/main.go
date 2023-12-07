package main

import (
	"log"
	"os"
	"sort"
	"strings"
)

var LineBreak = "\n"

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	res := solve2(string(data))
	log.Println("result:", res)
}

func solve1(input string) int {
	byLine := strings.Split(input, LineBreak)

	var hands []*Hand
	for _, l := range byLine {
		if l == "" {
			continue
		}
		hands = append(hands, NewHand(l))
	}

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].score == hands[j].score {
			return compare(hands[i], hands[j])
		}
		return hands[i].score < hands[j].score
	})

	sum := 0
	for i, h := range hands {
		sum += h.bid * (i + 1)
	}
	return sum
}

func solve2(input string) int {
	byLine := strings.Split(input, LineBreak)

	var hands []*Hand
	for _, l := range byLine {
		if l == "" {
			continue
		}
		hands = append(hands, NewHand(l))
	}

	sort.Slice(hands, func(i, j int) bool {
		if hands[i].score == hands[j].score {
			return compare(hands[i], hands[j])
		}
		return hands[i].score < hands[j].score
	})

	sum := 0
	for i, h := range hands {
		sum += h.bid * (i + 1)
	}
	return sum
}

func compare(i *Hand, j *Hand) bool {
	for i, c := range i.cards {
		if c.score == j.cards[i].score {
			continue
		}
		if c.score > j.cards[i].score {
			return false
		} else {
			return true
		}
	}
	return false
}
