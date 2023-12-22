package main

import (
	"advent-of-code/pkg/util"
	"log"
	"math"
	"os"
	"sort"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	res := solve1(string(data))
	log.Println("result for day22:", res)

	res2 := solve2(string(data))
	log.Println("result for day22 part 2:", res2)
}

func solve1(input string) int {
	bricks := parseBricks(input)
	bricks = sortBricksAndFallDown(bricks)
	kSupportsV, vSupportsK := findSupporting(bricks)

	total := 0
	for i, _ := range bricks {
		if allNotSupported(kSupportsV[i], vSupportsK) {
			total++
		}
	}
	return total
}

func findSupporting(bricks [][]int) ([]map[int]bool, []map[int]bool) {
	kSupportsV := make([]map[int]bool, len(bricks))
	vSupportsK := make([]map[int]bool, len(bricks))

	for j, upper := range bricks {
		if vSupportsK[j] == nil {
			vSupportsK[j] = make(map[int]bool)
		}
		for i, lower := range bricks[:j] {
			if kSupportsV[i] == nil {
				kSupportsV[i] = make(map[int]bool)
			}
			if overlaps(lower, upper) && upper[2] == lower[5]+1 {
				kSupportsV[i][j] = true
				vSupportsK[j][i] = true
			}
		}
	}
	return kSupportsV, vSupportsK
}

func sortBricksAndFallDown(bricks [][]int) [][]int {
	sort.SliceStable(bricks, func(i, j int) bool {
		return bricks[i][2] < bricks[j][2]
	})

	for index, brick := range bricks {
		maxZ := 1
		for _, check := range bricks[:index] {
			if overlaps(brick, check) {
				maxZ = max(maxZ, check[5]+1)
			}
		}
		brick[5] -= brick[2] - maxZ
		brick[2] = maxZ
	}

	sort.SliceStable(bricks, func(i, j int) bool {
		return bricks[i][2] < bricks[j][2]
	})
	return bricks
}

func allNotSupported(kSupportsV map[int]bool, vSupportsK []map[int]bool) bool {
	for j, _ := range kSupportsV {
		if len(vSupportsK[j]) < 2 {
			return false
		}
	}
	return true
}

func overlaps(a, b []int) bool {
	return max(a[0], b[0]) <= min(a[3], b[3]) &&
		max(a[1], b[1]) <= min(a[4], b[4])
}

func parseBricks(input string) [][]int {
	lines := util.SplitByLine(input)

	bricks := make([][]int, len(lines))
	for i, l := range lines {
		sp := strings.Split(strings.Replace(l, "~", ",", 1), ",")

		bricks[i] = make([]int, len(sp))
		for j, num := range sp {
			bricks[i][j] = util.StringToInt(num)
		}
	}
	return bricks
}

// Part 2

func solve2(input string) int {
	bricks := parseBricks(input)
	bricks = sortBricksAndFallDown(bricks)
	kSupportsV, vSupportsK := findSupporting(bricks)

	total := 0
	for i, _ := range bricks {
		q := NewQueue()
		falling := make(map[int]bool)

		for j, _ := range kSupportsV[i] {
			if len(vSupportsK[j]) == 1 {
				q.push(j)
				falling[j] = true
			}
		}
		falling[i] = true

		for q.len() > 0 {
			j := q.pop()

			for k, _ := range kSupportsV[j] {
				if _, ok := falling[k]; !ok {
					if allIn(vSupportsK[k], falling) {
						q.push(k)
						falling[k] = true
					}
				}
			}
		}
		total += len(falling) - 1
	}
	return total
}

func allIn(vSupportsK map[int]bool, falling map[int]bool) bool {
	for i, _ := range vSupportsK {
		if _, ok := falling[i]; !ok {
			return false
		}
	}
	return true
}

// utils

func max(a, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func min(a, b int) int {
	return int(math.Min(float64(a), float64(b)))
}
