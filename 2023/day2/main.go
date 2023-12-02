package main

import (
	"advent-of-code/pkg/util"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	bag := Round{
		red:   12,
		green: 13,
		blue:  14,
	}

	sum := getPossibleGamesIdSum(string(data), bag)
	log.Println("the result is:", sum)

	powerSum := getPowerOfLowestPossibles(string(data))
	log.Println("the result of lowest possibles is:", powerSum)
}

func getPossibleGamesIdSum(input string, bag Round) int {
	var games []Game
	lines := util.SplitByLine(input)
	for _, l := range lines {
		games = append(games, parseGame(l))
	}

	idSum := 0
	for _, g := range games {
		if g.isPossible(bag) {
			idSum = idSum + g.id
		}
	}
	return idSum
}

func getPowerOfLowestPossibles(input string) int {
	var games []Game
	lines := util.SplitByLine(input)
	for _, l := range lines {
		games = append(games, parseGame(l))
	}

	powerSum := 0
	for _, g := range games {
		lowestPossible := g.findLowestPossible()
		powerSum = powerSum + lowestPossible.getPower()
	}
	return powerSum
}

func parseGame(input string) Game {
	return Game{
		id:     parseId(input),
		rounds: parseRounds(input),
	}
}

func parseId(input string) int {
	input = strings.Replace(input, "Game ", "", 1)
	i := strings.Index(input, ":")
	res, err := strconv.Atoi(input[:i])
	if err != nil {
		log.Fatal(err)
	}
	return res
}

func parseRounds(input string) []Round {
	firstPart := strings.SplitN(input, ":", 2)
	roundsInput := strings.Split(firstPart[1], ";")
	var result []Round

	for _, i := range roundsInput {
		result = append(result, parseRound(i))
	}
	return result
}

func parseRound(input string) Round {
	result := Round{}

	commaSplitted := strings.Split(input, ",")
	for _, i := range commaSplitted {
		reds := parseColor(i, "red")
		if reds > 0 {
			result.red = reds
		}
		blues := parseColor(i, "blue")
		if blues > 0 {
			result.blue = blues
		}
		greens := parseColor(i, "green")
		if greens > 0 {
			result.green = greens
		}
	}
	return result
}

func parseColor(input string, color string) int {
	if strings.Contains(input, color) {
		input = strings.ReplaceAll(input, color, "")
		input = strings.ReplaceAll(input, " ", "")
		num, err := strconv.Atoi(input)
		if err != nil {
			log.Fatal(err)
		}
		return num
	}
	return 0
}
