package main

import (
	"strings"
	"testing"
)

func Test_Part1(t *testing.T) {
	input := "rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7"
	want := 1320
	got := solve1(input)
	if got != want {
		t.Errorf("wrong result: want=%d, got=%d", want, got)
	}
}

func Test_Part1_HashOfOneChar(t *testing.T) {
	input := "H"
	want := 200
	got := calcHash(input)
	if got != want {
		t.Errorf("wrong result: want=%d, got=%d", want, got)
	}
}

func Test_Part1_HashOfOneWord(t *testing.T) {
	input := "HASH"
	want := 52
	got := calcHash(input)
	if got != want {
		t.Errorf("wrong result: want=%d, got=%d", want, got)
	}
}

// Part 2

func Test_Part2(t *testing.T) {
	input := "rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7"
	want := 145
	got := solve2(input)
	if got != want {
		t.Errorf("wrong result: want=%d, got=%d", want, got)
	}
}

func Test_Part2_Boxes(t *testing.T) {
	input := "rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7"
	want := [][]Lens{
		{{label: "rn", value: 1}, {label: "cm", value: 2}},
		{},
		{},
		{{label: "ot", value: 7}, {label: "ab", value: 5}, {label: "pc", value: 6}},
	}
	got := processInputToBoxes(strings.Split(input, ","))
	assertBoxes(t, want, got)
}

func assertBoxes(t *testing.T, want [][]Lens, got [][]Lens) {
	t.Helper()
	for i, w := range want {
		for j, wl := range w {
			if wl != got[i][j] {
				t.Errorf("not right lens '%s' at pos '%d', got=%v", wl.label, wl.value, got[i][j])
			}
		}
	}
}
