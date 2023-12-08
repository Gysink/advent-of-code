package main

import "testing"

func TestPart1(t *testing.T) {
	input := "" +
		"RL\n" +
		"\n" +
		"AAA = (BBB, CCC)\n" +
		"BBB = (DDD, EEE)\n" +
		"CCC = (ZZZ, GGG)\n" +
		"DDD = (DDD, DDD)\n" +
		"EEE = (EEE, EEE)\n" +
		"GGG = (GGG, GGG)\n" +
		"ZZZ = (ZZZ, ZZZ)"
	want := 2
	got := solve1(input)
	if got != want {
		t.Errorf("wrong steps: got=%d, want=%d", got, want)
	}
}

func Test2Part1(t *testing.T) {
	input := "" +
		"LLR\n" +
		"\n" +
		"AAA = (BBB, BBB)\n" +
		"BBB = (AAA, ZZZ)\n" +
		"ZZZ = (ZZZ, ZZZ)"

	want := 6
	got := solve1(input)
	if got != want {
		t.Errorf("wrong steps: got=%d, want=%d", got, want)
	}
}

// Part 2

func TestPart2(t *testing.T) {
	input := "" +
		"LR\n" +
		"\n" +
		"11A = (11B, XXX)\n" +
		"11B = (XXX, 11Z)\n" +
		"11Z = (11B, XXX)\n" +
		"22A = (22B, XXX)\n" +
		"22B = (22C, 22C)\n" +
		"22C = (22Z, 22Z)\n" +
		"22Z = (22B, 22B)\n" +
		"XXX = (XXX, XXX)"

	want := float64(6)
	got := solve2(input)
	if got != want {
		t.Errorf("wrong steps: got=%f, want=%f", got, want)
	}
}

func TestPart2_2(t *testing.T) {
	input := "" +
		"LR\n" +
		"\n" +
		"11A = (11B, XXX)\n" +
		"11B = (XXX, 11Z)\n" +
		"11Z = (11B, XXX)\n" +
		"22A = (22B, XXX)\n" +
		"22B = (22C, 22C)\n" +
		"22C = (22Z, 22Z)\n" +
		"22Z = (22B, 22B)\n" +
		"33A = (33B, XXX)\n" +
		"33B = (33C, 33C)\n" +
		"33C = (33Z, 33Z)\n" +
		"33Z = (33B, 33B)\n" +
		"XXX = (XXX, XXX)"

	want := float64(6)
	got := solve2(input)
	if got != want {
		t.Errorf("wrong steps: got=%f, want=%f", got, want)
	}
}
