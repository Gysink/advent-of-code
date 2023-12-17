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

	res := solve1(string(data))
	log.Println("result for day17:", res)

	res2 := solve2(string(data))
	log.Println("result for day17 part 2:", res2)
}

func solve1(input string) int {
	board, maxX, maxY := createBoard(input)
	target := NewPosition(maxX, maxY)

	return findShortest(board, target, 1, 3)
}

func createBoard(input string) (map[Position]int, int, int) {
	lines := util.SplitByLine(input)

	board := make(map[Position]int, len(lines)*len(lines[0]))
	for x, l := range lines {
		for y, p := range l {
			board[NewPosition(x, y)] = util.StringToInt(string(p))
		}
	}
	return board, len(lines) - 1, len(lines[0]) - 1
}

func findShortest(board map[Position]int, target Position, minStraight, maxStraight int) int {
	type state struct {
		loc      Location
		straight int
	}
	type entry struct {
		state
		heatLoss int
	}

	q := NewPriorityQueue[entry](func(a, b entry) int {
		return a.heatLoss - b.heatLoss
	})

	q.Push(entry{
		state: state{
			loc:      NewLocation(0, 1, Right),
			straight: 1,
		},
	})
	q.Push(entry{
		state: state{
			loc:      NewLocation(1, 0, Down),
			straight: 1,
		},
	})
	visited := make(map[state]int)

	for !q.IsEmpty() {
		e := q.Pop()
		pos := e.loc.Pos

		if _, exists := board[pos]; !exists {
			continue
		}

		heat := board[pos] + e.heatLoss
		if pos == target {
			return heat // shortest path
		}

		if v, exists := visited[e.state]; exists {
			if v <= heat {
				continue
			}
		}
		visited[e.state] = heat

		if e.straight >= minStraight {
			q.Push(entry{
				state: state{
					loc:      e.loc.Turn(Left, 1),
					straight: 1,
				},
				heatLoss: heat,
			})

			q.Push(entry{
				state: state{
					loc:      e.loc.Turn(Right, 1),
					straight: 1,
				},
				heatLoss: heat,
			})
		}

		if e.straight < maxStraight {
			q.Push(entry{
				state: state{
					loc:      e.loc.Straight(1),
					straight: e.straight + 1,
				},
				heatLoss: heat,
			})
		}
	}
	log.Fatalln("no path found")
	return 0
}

// Part 2

func solve2(input string) int {
	board, maxX, maxY := createBoard(input)
	target := NewPosition(maxX, maxY)

	return findShortest(board, target, 4, 10)
}
