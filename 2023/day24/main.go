package main

import (
	"advent-of-code/pkg/util"
	"log"
	"os"
	"slices"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	res := solve1(string(data), 200000000000000, 400000000000000)
	log.Println("result for day24:", res)

	res2 := solve2(string(data))
	log.Println("result for day24 part 2:", res2)
}

func solve1(input string, lb, ub float64) int {
	lines := util.SplitByLine(input)
	var hailstones []Hailstone
	for _, l := range lines {
		hailstones = append(hailstones, NewHailstone(l))
	}
	total := 0

	for i, h := range hailstones {
		for _, h2 := range hailstones[:i] {
			a1, b1, c1 := h.a, h.b, h.c
			a2, b2, c2 := h2.a, h2.b, h2.c

			if a1*b2 == b1*a2 {
				continue
			}

			x := (c1*b2 - c2*b1) / (a1*b2 - a2*b1)
			y := (c2*a1 - c1*a2) / (a1*b2 - a2*b1)

			if lb <= x && x <= ub && lb <= y && y <= ub {
				if (x-h.sx)*h.vx >= 0 && (y-h.sy)*h.vy >= 0 && (x-h2.sx)*h2.vx >= 0 && (y-h2.sy)*h2.vy >= 0 {
					total++
				}
			}
		}
	}
	return total
}

// Part 2

type Vector3 struct {
	x, y, z float64
}

func solve2(input string) int {
	lines := util.SplitByLine(input)

	var hailstones []Hailstone
	for _, l := range lines {
		hailstones = append(hailstones, NewHailstone(l))
	}

	maybeX, maybeY, maybeZ := []int{}, []int{}, []int{}
	for i := 0; i < len(hailstones)-1; i++ {
		for j := i + 1; j < len(hailstones); j++ {
			a, b := hailstones[i], hailstones[j]
			if a.vx == b.vx {
				nextMaybe := findMatchingVel(int(b.sx-a.sx), int(a.vx))
				if len(maybeX) == 0 {
					maybeX = nextMaybe
				} else {
					maybeX = getIntersect(maybeX, nextMaybe)
				}
			}
			if a.vy == b.vy {
				nextMaybe := findMatchingVel(int(b.sy-a.sy), int(a.vy))
				if len(maybeY) == 0 {
					maybeY = nextMaybe
				} else {
					maybeY = getIntersect(maybeY, nextMaybe)
				}
			}
			if a.vz == b.vz {
				nextMaybe := findMatchingVel(int(b.sz-a.sz), int(a.vz))
				if len(maybeZ) == 0 {
					maybeZ = nextMaybe
				} else {
					maybeZ = getIntersect(maybeZ, nextMaybe)
				}
			}
		}
	}

	var result = 0
	if len(maybeX) == len(maybeY) && len(maybeY) == len(maybeZ) && len(maybeZ) == 1 {
		// only one possible velocity in all dimensions
		rockVel := Vector3{float64(maybeX[0]), float64(maybeY[0]), float64(maybeZ[0])}
		hailStoneA, hailStoneB := hailstones[0], hailstones[1]
		mA := (hailStoneA.vy - rockVel.y) / (hailStoneA.vx - rockVel.x)
		mB := (hailStoneB.vy - rockVel.y) / (hailStoneB.vx - rockVel.x)
		cA := hailStoneA.sy - (mA * hailStoneA.sx)
		cB := hailStoneB.sy - (mB * hailStoneB.sx)
		xPos := (cB - cA) / (mA - mB)
		yPos := mA*xPos + cA
		time := (xPos - hailStoneA.sx) / (hailStoneA.vx - rockVel.x)
		zPos := hailStoneA.sz + (hailStoneA.vz-rockVel.z)*time
		result = int(xPos + yPos + zPos)
	}

	return result
}

func findMatchingVel(dvel, pv int) []int {
	match := []int{}
	for v := -1000; v < 1000; v++ {
		if v != pv && dvel%(v-pv) == 0 {
			match = append(match, v)
		}
	}
	return match
}

func getIntersect(a, b []int) []int {
	result := []int{}
	for _, val := range a {
		if slices.Contains(b, val) {
			result = append(result, val)
		}
	}
	return result
}
