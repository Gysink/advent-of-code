package main

import (
	"advent-of-code/pkg/util"
	"strconv"
	"testing"
)

func Test_Part1(t *testing.T) {
	input := "" +
		"0 3 6 9 12 15" + util.GetLineBreak() +
		"1 3 6 10 15 21" + util.GetLineBreak() +
		"10 13 16 21 30 45"
	want := 114
	got := solve1(input)
	if got != want {
		t.Errorf("wrong result: want=%d, got=%d", want, got)
	}
}

func Test_Part1_Negatives(t *testing.T) {
	input := "" +
		"0 3 6 9 12 15" + util.GetLineBreak() +
		"1 3 6 10 15 21" + util.GetLineBreak() +
		"-2 2 6 10 14 18"
	want := 68
	got := solve1(input)
	if got != want {
		t.Errorf("wrong result: want=%d, got=%d", want, got)
	}
}

// part 2

func Test_Part2(t *testing.T) {
	input := "" +
		"0 3 6 9 12 15" + util.GetLineBreak() +
		"1 3 6 10 15 21" + util.GetLineBreak() +
		"10 13 16 21 30 45"
	want := 2
	got := solve2(input)
	if got != want {
		t.Errorf("wrong result: want=%d, got=%d", want, got)
	}
}

func Test_Part2_Negatives(t *testing.T) {
	input := "" +
		"0 3 6 9 12 15" + util.GetLineBreak() +
		"1 3 6 10 15 21" + util.GetLineBreak() +
		"-2 2 6 10 14 18"
	want := -9
	got := solve2(input)
	if got != want {
		t.Errorf("wrong result: want=%d, got=%d", want, got)
	}
}

func Test_Part2_Negatives2(t *testing.T) {
	input := "" +
		"-6 -11 -16 -21 -26"
	want := -1
	got := solve2(input)
	if got != want {
		t.Errorf("wrong result: want=%d, got=%d", want, got)
	}
}

func Test_Part2_Negatives3(t *testing.T) {
	input := "" +
		"137 145 70 -153 -589"
	want := 111
	got := solve2(input)
	if got != want {
		t.Errorf("wrong result: want=%d, got=%d", want, got)
	}
}

func TestDiff(t *testing.T) {
	tests := []struct {
		a    float64
		b    float64
		want float64
	}{
		{a: 0, b: 3, want: 3},
		{a: 3, b: 0, want: -3},
		{a: 0, b: 1, want: 1},
		{a: 1, b: 2, want: 1},
		{a: 1, b: 1, want: 0},
		{a: 0, b: 2, want: 2},
		{a: 2, b: 0, want: -2},
		{a: -2, b: 3, want: 5},
		{a: 5, b: 10, want: 5},
		{a: -5, b: -3, want: 2},
		{a: 4, b: -2, want: -6},
		{a: -3, b: -5, want: -2},
		{a: 14, b: 7, want: -7},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i+1), func(t *testing.T) {
			got := getDiff(tt.a, tt.b)
			if got != tt.want {
				t.Errorf("wrong diff, got=%f, want=%f", got, tt.want)
			}
		})
	}
}
