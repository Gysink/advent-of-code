package main

import (
	"advent-of-code/pkg/util"
	"reflect"
	"testing"
)

func Test_Part1(t *testing.T) {
	input := "" +
		"#.##..##." + util.GetLineBreak() +
		"..#.##.#." + util.GetLineBreak() +
		"##......#" + util.GetLineBreak() +
		"##......#" + util.GetLineBreak() +
		"..#.##.#." + util.GetLineBreak() +
		"..##..##." + util.GetLineBreak() +
		"#.#.##.#." + util.GetLineBreak() +
		util.GetLineBreak() +
		"#...##..#" + util.GetLineBreak() +
		"#....#..#" + util.GetLineBreak() +
		"..##..###" + util.GetLineBreak() +
		"#####.##." + util.GetLineBreak() +
		"#####.##." + util.GetLineBreak() +
		"..##..###" + util.GetLineBreak() +
		"#....#..#"
	want := 405
	got := solve1(input)
	if got != want {
		t.Errorf("wrong result: want=%d, got=%d", want, got)
	}
}

func Test_Part1_Vertical(t *testing.T) {
	input := "" +
		"#.##..##." + util.GetLineBreak() +
		"..#.##.#." + util.GetLineBreak() +
		"##......#" + util.GetLineBreak() +
		"##......#" + util.GetLineBreak() +
		"..#.##.#." + util.GetLineBreak() +
		"..##..##." + util.GetLineBreak() +
		"#.#.##.#."
	want := 5

	lines := util.SplitByLine(input)
	got := findReflections(rotateRight(lines), -1)
	if got != want {
		t.Errorf("wrong result: want=%d, got=%d", want, got)
	}
}

func Test_Part1_Horizontal(t *testing.T) {
	input := "" +
		"#...##..#" + util.GetLineBreak() +
		"#....#..#" + util.GetLineBreak() +
		"..##..###" + util.GetLineBreak() +
		"#####.##." + util.GetLineBreak() +
		"#####.##." + util.GetLineBreak() +
		"..##..###" + util.GetLineBreak() +
		"#....#..#"
	want := 4

	lines := util.SplitByLine(input)
	got := findReflections(lines, -1)
	if got != want {
		t.Errorf("wrong result: want=%d, got=%d", want, got)
	}
}

func Test_Part1_2(t *testing.T) {
	input := "" +
		"#####..#." + util.GetLineBreak() +
		"#..#.#.#." + util.GetLineBreak() +
		"#####..#." + util.GetLineBreak() +
		".##.##..#" + util.GetLineBreak() +
		"....#.###" + util.GetLineBreak() +
		"....#####" + util.GetLineBreak() +
		".##.##..#"
	want := 2
	got := solve1(input)
	if got != want {
		t.Errorf("wrong result: want=%d, got=%d", want, got)
	}
}

func Test_Part1_3(t *testing.T) {
	input := "" +
		"#.##..##." + util.GetLineBreak() +
		"..#.##.#." + util.GetLineBreak() +
		"##......#" + util.GetLineBreak() +
		"##......#" + util.GetLineBreak() +
		"..#.##.#." + util.GetLineBreak() +
		"..##..##." + util.GetLineBreak() +
		"#.#.##.#." + util.GetLineBreak() +
		util.GetLineBreak() +
		"#...##..#" + util.GetLineBreak() +
		"#....#..#" + util.GetLineBreak() +
		"..##..###" + util.GetLineBreak() +
		"#####.##." + util.GetLineBreak() +
		"#####.##." + util.GetLineBreak() +
		"..##..###" + util.GetLineBreak() +
		"#....#..#" + util.GetLineBreak() +
		util.GetLineBreak() +
		".#.##.#.#" + util.GetLineBreak() +
		".##..##.." + util.GetLineBreak() +
		".#.##.#.." + util.GetLineBreak() +
		"#......##" + util.GetLineBreak() +
		"#......##" + util.GetLineBreak() +
		".#.##.#.." + util.GetLineBreak() +
		".##..##.#" + util.GetLineBreak() +
		util.GetLineBreak() +
		"#..#....#" + util.GetLineBreak() +
		"###..##.." + util.GetLineBreak() +
		".##.#####" + util.GetLineBreak() +
		".##.#####" + util.GetLineBreak() +
		"###..##.." + util.GetLineBreak() +
		"#..#....#" + util.GetLineBreak() +
		"#..##...#" + util.GetLineBreak() +
		util.GetLineBreak() +
		"#.##..##." + util.GetLineBreak() +
		"..#.##.#." + util.GetLineBreak() +
		"##..#...#" + util.GetLineBreak() +
		"##...#..#" + util.GetLineBreak() +
		"..#.##.#." + util.GetLineBreak() +
		"..##..##." + util.GetLineBreak() +
		"#.#.##.#."
	want := 709
	got := solve1(input)
	if got != want {
		t.Errorf("wrong result: want=%d, got=%d", want, got)
	}
}

func Test_Part1_4(t *testing.T) {
	input := "" +
		"###.##.##" + util.GetLineBreak() +
		"##.####.#" + util.GetLineBreak() +
		"##.#..#.#" + util.GetLineBreak() +
		"####..###" + util.GetLineBreak() +
		"....##..." + util.GetLineBreak() +
		"##.#..#.#" + util.GetLineBreak() +
		"...#..#.." + util.GetLineBreak() +
		"##..###.#" + util.GetLineBreak() +
		"##......#" + util.GetLineBreak() +
		"##......#" + util.GetLineBreak() +
		"..#.##.#." + util.GetLineBreak() +
		"...#..#.." + util.GetLineBreak() +
		"##.####.#" + util.GetLineBreak() +
		"....##..." + util.GetLineBreak() +
		"...####.." + util.GetLineBreak() +
		"....##..." + util.GetLineBreak() +
		"##.####.#" + util.GetLineBreak() +
		util.GetLineBreak() +
		".##.##...##...##." + util.GetLineBreak() +
		"#####..##..##..##" + util.GetLineBreak() +
		".....##..##..##.." + util.GetLineBreak() +
		".##.#.#.####.#.#." + util.GetLineBreak() +
		".##...#.#..#.#..." + util.GetLineBreak() +
		"....#..........#." + util.GetLineBreak() +
		"#..#..#......#..#" + util.GetLineBreak() +
		"....###.....####." + util.GetLineBreak() +
		".##...#.#..#.#..." + util.GetLineBreak() +
		".....#..####..#.." + util.GetLineBreak() +
		"#..#...##..##...#" + util.GetLineBreak() +
		"....#...#..#...#." + util.GetLineBreak() +
		"#..#.##########.#" + util.GetLineBreak() +
		"#..##...####...##" + util.GetLineBreak() +
		"#####.##.##.##.##"
	want := 3
	got := solve1(input)
	if got != want {
		t.Errorf("wrong result: want=%d, got=%d", want, got)
	}
}

func Test_RotateRight(t *testing.T) {
	input := []string{
		"1az",
		"2by",
		"3cx",
		"4dw",
	}
	want := []string{
		"4321",
		"dcba",
		"wxyz",
	}

	got := rotateRight(input)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("swap failed: want=%v, got=%v", want, got)
	}
}

// Part 2

func Test_Part2(t *testing.T) {
	input := "" +
		"#.##..##." + util.GetLineBreak() +
		"..#.##.#." + util.GetLineBreak() +
		"##......#" + util.GetLineBreak() +
		"##......#" + util.GetLineBreak() +
		"..#.##.#." + util.GetLineBreak() +
		"..##..##." + util.GetLineBreak() +
		"#.#.##.#." + util.GetLineBreak() +
		util.GetLineBreak() +
		"#...##..#" + util.GetLineBreak() +
		"#....#..#" + util.GetLineBreak() +
		"..##..###" + util.GetLineBreak() +
		"#####.##." + util.GetLineBreak() +
		"#####.##." + util.GetLineBreak() +
		"..##..###" + util.GetLineBreak() +
		"#....#..#"
	want := 400
	got := solve2(input)
	if got != want {
		t.Errorf("wrong result: want=%d, got=%d", want, got)
	}
}

func Test_Part2_2(t *testing.T) {
	input := "" +
		"#.##..##." + util.GetLineBreak() +
		"..#.##.#." + util.GetLineBreak() +
		"##......#" + util.GetLineBreak() +
		"##......#" + util.GetLineBreak() +
		"..#.##.#." + util.GetLineBreak() +
		"..##..##." + util.GetLineBreak() +
		"#.#.##.#." + util.GetLineBreak() +
		util.GetLineBreak() +
		"#...##..#" + util.GetLineBreak() +
		"#....#..#" + util.GetLineBreak() +
		"..##..###" + util.GetLineBreak() +
		"#####.##." + util.GetLineBreak() +
		"#####.##." + util.GetLineBreak() +
		"..##..###" + util.GetLineBreak() +
		"#....#..#" + util.GetLineBreak() +
		util.GetLineBreak() +
		".#.##.#.#" + util.GetLineBreak() +
		".##..##.." + util.GetLineBreak() +
		".#.##.#.." + util.GetLineBreak() +
		"#......##" + util.GetLineBreak() +
		"#......##" + util.GetLineBreak() +
		".#.##.#.." + util.GetLineBreak() +
		".##..##.#" + util.GetLineBreak() +
		util.GetLineBreak() +
		"#..#....#" + util.GetLineBreak() +
		"###..##.." + util.GetLineBreak() +
		".##.#####" + util.GetLineBreak() +
		".##.#####" + util.GetLineBreak() +
		"###..##.." + util.GetLineBreak() +
		"#..#....#" + util.GetLineBreak() +
		"#..##...#" + util.GetLineBreak() +
		util.GetLineBreak() +
		"#.##..##." + util.GetLineBreak() +
		"..#.##.#." + util.GetLineBreak() +
		"##..#...#" + util.GetLineBreak() +
		"##...#..#" + util.GetLineBreak() +
		"..#.##.#." + util.GetLineBreak() +
		"..##..##." + util.GetLineBreak() +
		"#.#.##.#."
	want := 1400
	got := solve2(input)
	if got != want {
		t.Errorf("wrong result: want=%d, got=%d", want, got)
	}
}

func Test_Part2_3(t *testing.T) {
	input := "" +
		"###.##.##" + util.GetLineBreak() +
		"##.####.#" + util.GetLineBreak() +
		"##.#..#.#" + util.GetLineBreak() +
		"####..###" + util.GetLineBreak() +
		"....##..." + util.GetLineBreak() +
		"##.#..#.#" + util.GetLineBreak() +
		"...#..#.." + util.GetLineBreak() +
		"##..###.#" + util.GetLineBreak() +
		"##......#" + util.GetLineBreak() +
		"##......#" + util.GetLineBreak() +
		"..#.##.#." + util.GetLineBreak() +
		"...#..#.." + util.GetLineBreak() +
		"##.####.#" + util.GetLineBreak() +
		"....##..." + util.GetLineBreak() +
		"...####.." + util.GetLineBreak() +
		"....##..." + util.GetLineBreak() +
		"##.####.#" + util.GetLineBreak() +
		util.GetLineBreak() +
		".##.##...##...##." + util.GetLineBreak() +
		"#####..##..##..##" + util.GetLineBreak() +
		".....##..##..##.." + util.GetLineBreak() +
		".##.#.#.####.#.#." + util.GetLineBreak() +
		".##...#.#..#.#..." + util.GetLineBreak() +
		"....#..........#." + util.GetLineBreak() +
		"#..#..#......#..#" + util.GetLineBreak() +
		"....###.....####." + util.GetLineBreak() +
		".##...#.#..#.#..." + util.GetLineBreak() +
		".....#..####..#.." + util.GetLineBreak() +
		"#..#...##..##...#" + util.GetLineBreak() +
		"....#...#..#...#." + util.GetLineBreak() +
		"#..#.##########.#" + util.GetLineBreak() +
		"#..##...####...##" + util.GetLineBreak() +
		"#####.##.##.##.##"
	want := 15
	got := solve2(input)
	if got != want {
		t.Errorf("wrong result: want=%d, got=%d", want, got)
	}
}
