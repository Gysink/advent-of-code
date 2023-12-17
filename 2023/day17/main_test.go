package main

import (
	"advent-of-code/pkg/util"
	"testing"
)

func Test_Part1(t *testing.T) {
	input := "" +
		"2413432311323" + util.GetLineBreak() +
		"3215453535623" + util.GetLineBreak() +
		"3255245654254" + util.GetLineBreak() +
		"3446585845452" + util.GetLineBreak() +
		"4546657867536" + util.GetLineBreak() +
		"1438598798454" + util.GetLineBreak() +
		"4457876987766" + util.GetLineBreak() +
		"3637877979653" + util.GetLineBreak() +
		"4654967986887" + util.GetLineBreak() +
		"4564679986453" + util.GetLineBreak() +
		"1224686865563" + util.GetLineBreak() +
		"2546548887735" + util.GetLineBreak() +
		"4322674655533"
	want := 102
	got := solve1(input)
	if got != want {
		t.Errorf("wrong result: want=%d, got=%d", want, got)
	}
}

// Part 2

func Test_Part2(t *testing.T) {
	input := "" +
		"2413432311323" + util.GetLineBreak() +
		"3215453535623" + util.GetLineBreak() +
		"3255245654254" + util.GetLineBreak() +
		"3446585845452" + util.GetLineBreak() +
		"4546657867536" + util.GetLineBreak() +
		"1438598798454" + util.GetLineBreak() +
		"4457876987766" + util.GetLineBreak() +
		"3637877979653" + util.GetLineBreak() +
		"4654967986887" + util.GetLineBreak() +
		"4564679986453" + util.GetLineBreak() +
		"1224686865563" + util.GetLineBreak() +
		"2546548887735" + util.GetLineBreak() +
		"4322674655533"
	want := 94
	got := solve2(input)
	if got != want {
		t.Errorf("wrong result: want=%d, got=%d", want, got)
	}
}
