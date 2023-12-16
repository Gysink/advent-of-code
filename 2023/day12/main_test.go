package main

import (
	"advent-of-code/pkg/util"
	"strconv"
	"testing"
)

func Test_Part1(t *testing.T) {
	input := "" +
		"???.### 1,1,3" + util.GetLineBreak() +
		".??..??...?##. 1,1,3" + util.GetLineBreak() +
		"?#?#?#?#?#?#?#? 1,3,1,6" + util.GetLineBreak() +
		"????.#...#... 4,1,1" + util.GetLineBreak() +
		"????.######..#####. 1,6,5" + util.GetLineBreak() +
		"?###???????? 3,2,1"
	want := 21
	got := solve1(input)
	if got != want {
		t.Errorf("wrong result: want=%d, got=%d", want, got)
	}
}

func TestPart1_TestCalcArrangements(t *testing.T) {
	tests := []struct {
		input string
		nums  []int
		want  int
	}{
		{input: "???.###", nums: []int{1, 1, 3}, want: 1},
		{input: ".??..??...?##.", nums: []int{1, 1, 3}, want: 4},
		{input: "?#?#?#?#?#?#?#?", nums: []int{1, 3, 1, 6}, want: 1},
		{input: "????.#...#...", nums: []int{4, 1, 1}, want: 1},
		{input: "????.######..#####.", nums: []int{1, 6, 5}, want: 4},
		{input: "?###????????", nums: []int{3, 2, 1}, want: 10},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := calcArrangements(tt.input, tt.nums)
			if got != tt.want {
				t.Errorf("wrong arrangements, got=%d, want=%d", got, tt.want)
			}
		})
	}
}

// Part 2

func Test_Part2(t *testing.T) {
	input := "" +
		"???.### 1,1,3" + util.GetLineBreak() +
		".??..??...?##. 1,1,3" + util.GetLineBreak() +
		"?#?#?#?#?#?#?#? 1,3,1,6" + util.GetLineBreak() +
		"????.#...#... 4,1,1" + util.GetLineBreak() +
		"????.######..#####. 1,6,5" + util.GetLineBreak() +
		"?###???????? 3,2,1"
	want := 525152
	got := solve2(input)
	if got != want {
		t.Errorf("wrong result: want=%d, got=%d", want, got)
	}
}
