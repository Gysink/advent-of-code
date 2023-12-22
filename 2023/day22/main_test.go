package main

import (
	"advent-of-code/pkg/util"
	"testing"
)

func Test_Part1(t *testing.T) {
	input := "" +
		"1,0,1~1,2,1" + util.GetLineBreak() +
		"0,0,2~2,0,2" + util.GetLineBreak() +
		"0,2,3~2,2,3" + util.GetLineBreak() +
		"0,0,4~0,2,4" + util.GetLineBreak() +
		"2,0,5~2,2,5" + util.GetLineBreak() +
		"0,1,6~2,1,6" + util.GetLineBreak() +
		"1,1,8~1,1,9"
	want := 5
	got := solve1(input)
	if got != want {
		t.Errorf("wrong result: want=%d, got=%d", want, got)
	}
}

// Part 2

func Test_Part2(t *testing.T) {
	input := "" +
		"1,0,1~1,2,1" + util.GetLineBreak() +
		"0,0,2~2,0,2" + util.GetLineBreak() +
		"0,2,3~2,2,3" + util.GetLineBreak() +
		"0,0,4~0,2,4" + util.GetLineBreak() +
		"2,0,5~2,2,5" + util.GetLineBreak() +
		"0,1,6~2,1,6" + util.GetLineBreak() +
		"1,1,8~1,1,9"
	want := 7
	got := solve2(input)
	if got != want {
		t.Errorf("wrong result: want=%d, got=%d", want, got)
	}
}
