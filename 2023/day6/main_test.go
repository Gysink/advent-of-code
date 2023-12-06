package main

import "testing"

func Test_Part1(t *testing.T) {
	input := "" +
		"Time:      7  15   30\n" +
		"Distance:  9  40  200"
	want := 288
	got := solve1(input)
	if got != want {
		t.Errorf("wrong result: want=%d, got=%d", want, got)
	}
}

func Test_Part2(t *testing.T) {
	input := "" +
		"Time:      7  15   30\n" +
		"Distance:  9  40  200"
	want := 71503
	got := solve2(input)
	if got != want {
		t.Errorf("wrong result: want=%d, got=%d", want, got)
	}
}
