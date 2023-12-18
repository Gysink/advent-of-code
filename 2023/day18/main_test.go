package main

import (
	"advent-of-code/pkg/util"
	"testing"
)

func Test_Part1(t *testing.T) {
	input := "" +
		"R 6 (#70c710)" + util.GetLineBreak() +
		"D 5 (#0dc571)" + util.GetLineBreak() +
		"L 2 (#5713f0)" + util.GetLineBreak() +
		"D 2 (#d2c081)" + util.GetLineBreak() +
		"R 2 (#59c680)" + util.GetLineBreak() +
		"D 2 (#411b91)" + util.GetLineBreak() +
		"L 5 (#8ceee2)" + util.GetLineBreak() +
		"U 2 (#caa173)" + util.GetLineBreak() +
		"L 1 (#1b58a2)" + util.GetLineBreak() +
		"U 2 (#caa171)" + util.GetLineBreak() +
		"R 2 (#7807d2)" + util.GetLineBreak() +
		"U 3 (#a77fa3)" + util.GetLineBreak() +
		"L 2 (#015232)" + util.GetLineBreak() +
		"U 2 (#7a21e3)"
	want := 62
	got := solve1(input)
	if got != want {
		t.Errorf("wrong result: want=%d, got=%d", want, got)
	}
}

// Part 2

func Test_Part2(t *testing.T) {
	input := "" +
		"R 6 (#70c710)" + util.GetLineBreak() +
		"D 5 (#0dc571)" + util.GetLineBreak() +
		"L 2 (#5713f0)" + util.GetLineBreak() +
		"D 2 (#d2c081)" + util.GetLineBreak() +
		"R 2 (#59c680)" + util.GetLineBreak() +
		"D 2 (#411b91)" + util.GetLineBreak() +
		"L 5 (#8ceee2)" + util.GetLineBreak() +
		"U 2 (#caa173)" + util.GetLineBreak() +
		"L 1 (#1b58a2)" + util.GetLineBreak() +
		"U 2 (#caa171)" + util.GetLineBreak() +
		"R 2 (#7807d2)" + util.GetLineBreak() +
		"U 3 (#a77fa3)" + util.GetLineBreak() +
		"L 2 (#015232)" + util.GetLineBreak() +
		"U 2 (#7a21e3)"
	want := 952408144115
	got := solve2(input)
	if got != want {
		t.Errorf("wrong result: want=%d, got=%d", want, got)
	}
}
