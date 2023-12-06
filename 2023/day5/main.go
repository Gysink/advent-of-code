package main

import (
	"log"
	"os"
	"strings"
)

var LineBreak = "\r\n"

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	a := parseAlmanacFromInput(string(data))
	log.Println("result is:", a.getResult())
}

func parseAlmanacFromInput(input string) *Almanac {
	blocks := strings.Split(input, LineBreak+LineBreak)
	almanac := NewAlmanac(blocks[0])

	for _, b := range blocks[1:] {
		almanac.mappers = append(almanac.mappers, NewMapper(b))
	}

	return almanac
}
