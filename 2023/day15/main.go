package main

import (
	"advent-of-code/pkg/util"
	"log"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	//data, err := os.ReadFile("2023/day15/input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	res := solve1(string(data))
	log.Println("result for day15:", res)

	res2 := solve2(string(data))
	log.Println("result for day15 part 2:", res2)

}

func solve1(input string) int {
	blocks := strings.Split(input, ",")
	sum := 0
	for _, b := range blocks {
		sum += calcHash(b)
	}
	return sum
}

func calcHash(input string) int {
	value := 0
	for _, c := range input {
		value += int(c)
		value = value * 17
		value = value % 256
	}
	return value
}

// Part 2

type Lens struct {
	label string
	value int
}

func newBox(val string) Lens {
	sp := strings.Split(val, "=")
	return Lens{
		label: sp[0],
		value: util.StringToInt(sp[1]),
	}
}

func solve2(input string) int {
	blocks := strings.Split(input, ",")
	boxes := processInputToBoxes(blocks)

	sum := 0
	for bI, b := range boxes {
		if len(b) > 0 {
			for lI, l := range b {
				sum += (bI + 1) * (lI + 1) * l.value
			}
		}
	}
	return sum
}

func processInputToBoxes(input []string) [][]Lens {
	boxes := make([][]Lens, 256)

	for _, val := range input {
		if strings.Contains(val, "=") {
			b := newBox(val)
			boxI := calcHash(b.label)
			index := indexOf(boxes[boxI], b.label)
			if index != -1 {
				boxes[boxI][index].value = b.value
			} else {
				boxes[boxI] = append(boxes[boxI], b)
			}
		} else {
			label := strings.Replace(val, "-", "", 1)
			boxI := calcHash(label)
			index := indexOf(boxes[boxI], label)
			if index != -1 {
				boxes[boxI] = append(boxes[boxI][:index], boxes[boxI][index+1:]...)
			}
		}
	}
	return boxes
}

func indexOf(boxes []Lens, label string) int {
	for i, b := range boxes {
		if b.label == label {
			return i
		}
	}
	return -1
}
