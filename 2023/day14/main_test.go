package main

import (
	"advent-of-code/pkg/util"
	"testing"
)

func Test_Part1(t *testing.T) {
	input := "" +
		"O....#...." + util.GetLineBreak() +
		"O.OO#....#" + util.GetLineBreak() +
		".....##..." + util.GetLineBreak() +
		"OO.#O....O" + util.GetLineBreak() +
		".O.....O#." + util.GetLineBreak() +
		"O.#..O.#.#" + util.GetLineBreak() +
		"..O..#O..O" + util.GetLineBreak() +
		".......O.." + util.GetLineBreak() +
		"#....###.." + util.GetLineBreak() +
		"#OO..#...."
	want := 136
	got := solve1(input)
	if got != want {
		t.Errorf("wrong result: want=%d, got=%d", want, got)
	}
}

// Part 2

func Test_Part2(t *testing.T) {
	input := "" +
		"O....#...." + util.GetLineBreak() +
		"O.OO#....#" + util.GetLineBreak() +
		".....##..." + util.GetLineBreak() +
		"OO.#O....O" + util.GetLineBreak() +
		".O.....O#." + util.GetLineBreak() +
		"O.#..O.#.#" + util.GetLineBreak() +
		"..O..#O..O" + util.GetLineBreak() +
		".......O.." + util.GetLineBreak() +
		"#....###.." + util.GetLineBreak() +
		"#OO..#...."
	want := 64
	got := solve2(input, 1000000000)
	if got != want {
		t.Errorf("wrong result: want=%d, got=%d", want, got)
	}
}
