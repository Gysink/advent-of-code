package main

import (
	"advent-of-code/pkg/util"
	"testing"
)

func Test_Part1(t *testing.T) {
	input := "" +
		"..........." + util.GetLineBreak() +
		".....###.#." + util.GetLineBreak() +
		".###.##..#." + util.GetLineBreak() +
		"..#.#...#.." + util.GetLineBreak() +
		"....#.#...." + util.GetLineBreak() +
		".##..S####." + util.GetLineBreak() +
		".##..#...#." + util.GetLineBreak() +
		".......##.." + util.GetLineBreak() +
		".##.#.####." + util.GetLineBreak() +
		".##..##.##." + util.GetLineBreak() +
		"..........."
	want := 16
	got := solve1(input, 6)
	if got != want {
		t.Errorf("wrong result: want=%d, got=%d", want, got)
	}
}
