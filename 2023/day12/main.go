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
	log.Println("result for day12:", res)

	res2 := solve2(string(data))
	log.Println("result for day12 part 2:", res2)
}

func solve1(input string) int {
	lines := util.SplitByLine(input)

	sum := 0
	for _, l := range lines {
		sp := strings.Split(l, " ")
		line := sp[0]
		nums := strings.Split(sp[1], ",")

		sum += calcArrangements(line, util.StringsToIntArray(nums))
	}

	return sum
}

func calcArrangements(line string, nums []int) int {
	return countPossible(line, nums)
}

func countPossible(s string, c []int) int {
	pos := 0
	// state is a tuple of 4 values
	cstates := map[[4]int]int{{0, 0, 0, 0}: 1}
	nstates := map[[4]int]int{}
	for len(cstates) > 0 {
		for state, num := range cstates {
			si, ci, cc, expdot := state[0], state[1], state[2], state[3]
			if si == len(s) {
				if ci == len(c) {
					pos += num
				}
				continue
			}
			switch {
			case (s[si] == '#' || s[si] == '?') && ci < len(c) && expdot == 0:
				// we are still looking for broken springs
				if s[si] == '?' && cc == 0 {
					// we are not in a run of broken springs, so ? can be working
					nstates[[4]int{si + 1, ci, cc, expdot}] += num
				}
				cc++
				if cc == c[ci] {
					// we've found the full next contiguous section of broken springs
					ci++
					cc = 0
					expdot = 1 // we only want a working spring next
				}
				nstates[[4]int{si + 1, ci, cc, expdot}] += num
			case (s[si] == '.' || s[si] == '?') && cc == 0:
				// we are not in a contiguous run of broken springs
				expdot = 0
				nstates[[4]int{si + 1, ci, cc, expdot}] += num
			}
		}
		cstates, nstates = nstates, cstates
		mapsClear(nstates)
	}
	return pos
}

func mapsClear[M ~map[K]V, K comparable, V any](m M) {
	for k := range m {
		delete(m, k)
	}
}

// part 2

func solve2(input string) int {
	lines := util.SplitByLine(input)

	sum := 0
	for _, l := range lines {
		sp := strings.Split(l, " ")
		line := sp[0]
		nums := strings.Split(sp[1], ",")
		line = unfoldLine(line)
		nums = unfoldNums(nums)
		sum += calcArrangements(line, util.StringsToIntArray(nums))
	}

	return sum
}

func unfoldNums(nums []string) []string {
	res := append([]string{}, nums...)
	res = append(res, nums...)
	res = append(res, nums...)
	res = append(res, nums...)
	res = append(res, nums...)
	return res
}

func unfoldLine(line string) string {
	return line + "?" + line + "?" + line + "?" + line + "?" + line
}
