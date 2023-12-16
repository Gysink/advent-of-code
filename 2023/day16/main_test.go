package main

import (
	"advent-of-code/pkg/util"
	"testing"
)

func Test_Part1(t *testing.T) {
	input := "" +
		".|...\\...." + util.GetLineBreak() +
		"|.-.\\....." + util.GetLineBreak() +
		".....|-..." + util.GetLineBreak() +
		"........|." + util.GetLineBreak() +
		".........." + util.GetLineBreak() +
		".........\\" + util.GetLineBreak() +
		"..../.\\\\.." + util.GetLineBreak() +
		".-.-/..|.." + util.GetLineBreak() +
		".|....-|.\\" + util.GetLineBreak() +
		"..//.|...."
	want := 46
	got := solve1(input)
	if got != want {
		t.Errorf("wrong result: want=%d, got=%d", want, got)
	}
}

func Test_Part1_Loop(t *testing.T) {
	input := "" +
		".........\\" + util.GetLineBreak() +
		"/........-" + util.GetLineBreak() +
		"-......../"
	want := 30
	got := solve1(input)
	if got != want {
		t.Errorf("wrong result: want=%d, got=%d", want, got)
	}
}

// Part 2

func Test_Part2(t *testing.T) {
	input := "" +
		".|...\\...." + util.GetLineBreak() +
		"|.-.\\....." + util.GetLineBreak() +
		".....|-..." + util.GetLineBreak() +
		"........|." + util.GetLineBreak() +
		".........." + util.GetLineBreak() +
		".........\\" + util.GetLineBreak() +
		"..../.\\\\.." + util.GetLineBreak() +
		".-.-/..|.." + util.GetLineBreak() +
		".|....-|.\\" + util.GetLineBreak() +
		"..//.|...."
	want := 51
	got := solve2(input)
	if got != want {
		t.Errorf("wrong result: want=%d, got=%d", want, got)
	}
}
