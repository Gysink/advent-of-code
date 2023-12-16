package main

import (
	"advent-of-code/pkg/util"
	"log"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	res := solve1(string(data))
	log.Println("result for day13:", res)

	res2 := solve2(string(data))
	log.Println("result for day13 part 2:", res2)
}

func solve1(input string) int {
	blocks := strings.Split(input, util.GetLineBreak()+util.GetLineBreak())

	sum := 0
	for _, b := range blocks {
		lines := util.SplitByLine(b)
		sum += findReflections(rotateRight(lines), -1)
		sum += 100 * findReflections(lines, -1)
	}
	return sum
}

func findReflections(lines []string, prevFound int) int {
	foundIndex := 0
	for i, l := range lines {
		if (i+1) < len(lines) && l == lines[i+1] {
			if i == (prevFound - 1) {
				continue
			}
			if isCompleteRefl(lines, i) {
				foundIndex = i + 1
			}
		}
	}
	return foundIndex
}

func isCompleteRefl(lines []string, index int) bool {
	halfLen := (len(lines) - 1) / 2
	checked := 0
	for i := 1; i < halfLen; i++ {
		if (index-i) >= 0 && (index+i+1) < len(lines) {
			if lines[index-i] != lines[index+i+1] {
				return false
			}
			checked++
		}
	}
	return true
}

func rotateRight(lines []string) []string {
	var result []string
	for i := 0; i < len(lines[0]); i++ {
		newLine := ""
		for j := len(lines) - 1; j >= 0; j-- {
			newLine += string(lines[j][i])
		}
		result = append(result, newLine)
	}
	return result
}

// part 2

func solve2(input string) int {
	blocks := strings.Split(input, util.GetLineBreak()+util.GetLineBreak())

	sum := 0
	for _, b := range blocks {
		lines := util.SplitByLine(b)
		highestHor := 0
		highestVert := 0
		oldFindingHor := findReflections(lines, -1)
		oldFindingVert := findReflections(rotateRight(lines), -1)
		for i := 0; i < len(lines); i++ {
			for j := 0; j < len(lines[0]); j++ {
				a := alternate(lines, i, j)
				refs := findReflections(a, oldFindingHor)
				if refs > highestHor {
					highestHor = refs
				}
				refsVert := findReflections(rotateRight(a), oldFindingVert)
				if refsVert > highestVert {
					highestVert = refsVert
				}
			}
		}
		sum += 100 * highestHor
		if highestHor == 0 {
			sum += highestVert
		}
	}
	return sum
}

func alternate(lines []string, i int, j int) []string {
	result := make([]string, len(lines))
	copy(result, lines)
	symbol := result[i][j]
	switch symbol {
	case '.':
		result[i] = result[i][:j] + "#" + result[i][j+1:]
	case '#':
		result[i] = result[i][:j] + "." + result[i][j+1:]
	}
	return result
}
