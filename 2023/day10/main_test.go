package main

import (
	"advent-of-code/pkg/util"
	"strconv"
	"testing"
)

func Test_Part1(t *testing.T) {
	input := "" +
		"....." + util.GetLineBreak() +
		".S-7." + util.GetLineBreak() +
		".|.|." + util.GetLineBreak() +
		".L-J." + util.GetLineBreak() +
		"....."
	want := 4
	got := solve1(input)
	if got != want {
		t.Errorf("wrong result: want=%d, got=%d", want, got)
	}
}

func Test_Part1_2(t *testing.T) {
	input := "" +
		"....." + util.GetLineBreak() +
		".F-S." + util.GetLineBreak() +
		".|.|." + util.GetLineBreak() +
		".|FJ." + util.GetLineBreak() +
		".LJ.."
	want := 5
	got := solve1(input)
	if got != want {
		t.Errorf("wrong result: want=%d, got=%d", want, got)
	}
}

func Test_Part1_3(t *testing.T) {
	input := "" +
		".F-7." + util.GetLineBreak() +
		".S||." + util.GetLineBreak() +
		".|.|." + util.GetLineBreak() +
		".|FJ." + util.GetLineBreak() +
		".LJ.."
	want := 6
	got := solve1(input)
	if got != want {
		t.Errorf("wrong result: want=%d, got=%d", want, got)
	}
}

func TestFindNext(t *testing.T) {
	topLeft := Point{symbol: "F", x: 0, y: 0}
	topMid := Point{symbol: "S", x: 0, y: 1}
	topRight := Point{symbol: "7", x: 0, y: 2}

	midLeft := Point{symbol: "|", x: 1, y: 0}
	midMid := Point{symbol: ".", x: 1, y: 1}
	midRight := Point{symbol: "|", x: 1, y: 2}

	bottomLeft := Point{symbol: "L", x: 2, y: 0}
	bottomMid := Point{symbol: "-", x: 2, y: 1}
	bottomRight := Point{symbol: "J", x: 2, y: 2}

	points := [][]Point{
		{topLeft, topMid, topRight},
		{midLeft, midMid, midRight},
		{bottomLeft, bottomMid, bottomRight},
	}

	tests := []struct {
		prev    Point
		current Point
		want    Point
	}{
		{prev: Point{}, current: topMid, want: topRight},
		{prev: topMid, current: topRight, want: midRight},
		{prev: topRight, current: midRight, want: bottomRight},
		{prev: midRight, current: bottomRight, want: bottomMid},
		{prev: bottomRight, current: bottomMid, want: bottomLeft},
		{prev: bottomMid, current: bottomLeft, want: midLeft},
		{prev: bottomLeft, current: midLeft, want: topLeft},
		{prev: midLeft, current: topLeft, want: topMid},
	}

	for i, tt := range tests {
		t.Run(strconv.Itoa(i+1), func(t *testing.T) {
			got := tt.current.next(points, tt.prev)
			if got.x != tt.want.x || got.y != tt.want.y || got.symbol != tt.want.symbol {
				//if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("wrong next point: want=%v, got=%v", tt.want, got)
			}
		})
	}
}

// Part 2

func Test_Part2(t *testing.T) {
	input := "" +
		"....." + util.GetLineBreak() +
		".S-7." + util.GetLineBreak() +
		".|.|." + util.GetLineBreak() +
		".L-J." + util.GetLineBreak() +
		"....."
	want := 1
	got := solve2(input)
	if got != want {
		t.Errorf("wrong result: want=%d, got=%d", want, got)
	}
}

func Test_Part2_2(t *testing.T) {
	input := "" +
		"..........." + util.GetLineBreak() +
		".S-------7." + util.GetLineBreak() +
		".|F-----7|." + util.GetLineBreak() +
		".||.....||." + util.GetLineBreak() +
		".||.....||." + util.GetLineBreak() +
		".|L-7.F-J|." + util.GetLineBreak() +
		".|..|.|..|." + util.GetLineBreak() +
		".L--J.L--J." + util.GetLineBreak() +
		"..........."
	want := 4
	got := solve2(input)
	if got != want {
		t.Errorf("wrong result: want=%d, got=%d", want, got)
	}
}

func Test_Part2_3(t *testing.T) {
	input := "" +
		".F----7F7F7F7F-7...." + util.GetLineBreak() +
		".|F--7||||||||FJ...." + util.GetLineBreak() +
		".||.FJ||||||||L7...." + util.GetLineBreak() +
		"FJL7L7LJLJ||LJ.L-7.." + util.GetLineBreak() +
		"L--J.L7...LJS7F-7L7." + util.GetLineBreak() +
		"....F-J..F7FJ|L7L7L7" + util.GetLineBreak() +
		"....L7.F7||L7|.L7L7|" + util.GetLineBreak() +
		".....|FJLJ|FJ|F7|.LJ" + util.GetLineBreak() +
		"....FJL-7.||.||||..." + util.GetLineBreak() +
		"....L---J.LJ.LJLJ..."
	want := 8
	got := solve2(input)
	if got != want {
		t.Errorf("wrong result: want=%d, got=%d", want, got)
	}
}

func Test_Part2_4(t *testing.T) {
	input := "" +
		"FF7FSF7F7F7F7F7F---7" + util.GetLineBreak() +
		"L|LJ||||||||||||F--J" + util.GetLineBreak() +
		"FL-7LJLJ||||||LJL-77" + util.GetLineBreak() +
		"F--JF--7||LJLJ7F7FJ-" + util.GetLineBreak() +
		"L---JF-JLJ.||-FJLJJ7" + util.GetLineBreak() +
		"|F|F-JF---7F7-L7L|7|" + util.GetLineBreak() +
		"|FFJF7L7F-JF7|JL---7" + util.GetLineBreak() +
		"7-L-JL7||F7|L7F-7F7|" + util.GetLineBreak() +
		"L.L7LFJ|||||FJL7||LJ" + util.GetLineBreak() +
		"L7JLJL-JLJLJL--JLJ.L"
	want := 10
	got := solve2(input)
	if got != want {
		t.Errorf("wrong result: want=%d, got=%d", want, got)
	}
}
