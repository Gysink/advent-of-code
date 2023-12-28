package main

import (
	"advent-of-code/pkg/util"
	"testing"
)

func Test_Part1(t *testing.T) {
	input := "" +
		"19, 13, 30 @ -2,  1, -2" + util.GetLineBreak() +
		"18, 19, 22 @ -1, -1, -2" + util.GetLineBreak() +
		"20, 25, 34 @ -2, -2, -4" + util.GetLineBreak() +
		"12, 31, 28 @ -1, -2, -1" + util.GetLineBreak() +
		"20, 19, 15 @  1, -5, -3"
	want := 2
	got := solve1(input, 7, 27)
	if got != want {
		t.Errorf("wrong result: want=%d, got=%d", want, got)
	}
}
