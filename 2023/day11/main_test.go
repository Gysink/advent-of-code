package main

import (
	"advent-of-code/pkg/util"
	"testing"
)

func Test_Part1(t *testing.T) {
	input := "" +
		"...#......" + util.GetLineBreak() +
		".......#.." + util.GetLineBreak() +
		"#........." + util.GetLineBreak() +
		".........." + util.GetLineBreak() +
		"......#..." + util.GetLineBreak() +
		".#........" + util.GetLineBreak() +
		".........#" + util.GetLineBreak() +
		".........." + util.GetLineBreak() +
		".......#.." + util.GetLineBreak() +
		"#...#....."
	want := 374
	got := solve1(input)
	if got != want {
		t.Errorf("wrong result: want=%d, got=%d", want, got)
	}
}

// Part 2

func Test_Part2(t *testing.T) {
	input := "" +
		"...#......" + util.GetLineBreak() +
		".......#.." + util.GetLineBreak() +
		"#........." + util.GetLineBreak() +
		".........." + util.GetLineBreak() +
		"......#..." + util.GetLineBreak() +
		".#........" + util.GetLineBreak() +
		".........#" + util.GetLineBreak() +
		".........." + util.GetLineBreak() +
		".......#.." + util.GetLineBreak() +
		"#...#....."
	want := 8410
	got := calc(input, 100)
	if got != want {
		t.Errorf("wrong result: want=%d, got=%d", want, got)
	}
}
