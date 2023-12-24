package main

import (
	"advent-of-code/pkg/util"
	"testing"
)

func Test_Part1(t *testing.T) {
	input := "" +
		"#.#####################" + util.GetLineBreak() +
		"#.......#########...###" + util.GetLineBreak() +
		"#######.#########.#.###" + util.GetLineBreak() +
		"###.....#.>.>.###.#.###" + util.GetLineBreak() +
		"###v#####.#v#.###.#.###" + util.GetLineBreak() +
		"###.>...#.#.#.....#...#" + util.GetLineBreak() +
		"###v###.#.#.#########.#" + util.GetLineBreak() +
		"###...#.#.#.......#...#" + util.GetLineBreak() +
		"#####.#.#.#######.#.###" + util.GetLineBreak() +
		"#.....#.#.#.......#...#" + util.GetLineBreak() +
		"#.#####.#.#.#########v#" + util.GetLineBreak() +
		"#.#...#...#...###...>.#" + util.GetLineBreak() +
		"#.#.#v#######v###.###v#" + util.GetLineBreak() +
		"#...#.>.#...>.>.#.###.#" + util.GetLineBreak() +
		"#####v#.#.###v#.#.###.#" + util.GetLineBreak() +
		"#.....#...#...#.#.#...#" + util.GetLineBreak() +
		"#.#########.###.#.#.###" + util.GetLineBreak() +
		"#...###...#...#...#.###" + util.GetLineBreak() +
		"###.###.#.###v#####v###" + util.GetLineBreak() +
		"#...#...#.#.>.>.#.>.###" + util.GetLineBreak() +
		"#.###.###.#.###.#.#v###" + util.GetLineBreak() +
		"#.....###...###...#...#" + util.GetLineBreak() +
		"#####################.#"
	want := 94
	got := solve1(input)
	if got != want {
		t.Errorf("wrong result: want=%d, got=%d", want, got)
	}
}

// Part 2

func Test_Part2(t *testing.T) {
	input := "" +
		"#.#####################" + util.GetLineBreak() +
		"#.......#########...###" + util.GetLineBreak() +
		"#######.#########.#.###" + util.GetLineBreak() +
		"###.....#.>.>.###.#.###" + util.GetLineBreak() +
		"###v#####.#v#.###.#.###" + util.GetLineBreak() +
		"###.>...#.#.#.....#...#" + util.GetLineBreak() +
		"###v###.#.#.#########.#" + util.GetLineBreak() +
		"###...#.#.#.......#...#" + util.GetLineBreak() +
		"#####.#.#.#######.#.###" + util.GetLineBreak() +
		"#.....#.#.#.......#...#" + util.GetLineBreak() +
		"#.#####.#.#.#########v#" + util.GetLineBreak() +
		"#.#...#...#...###...>.#" + util.GetLineBreak() +
		"#.#.#v#######v###.###v#" + util.GetLineBreak() +
		"#...#.>.#...>.>.#.###.#" + util.GetLineBreak() +
		"#####v#.#.###v#.#.###.#" + util.GetLineBreak() +
		"#.....#...#...#.#.#...#" + util.GetLineBreak() +
		"#.#########.###.#.#.###" + util.GetLineBreak() +
		"#...###...#...#...#.###" + util.GetLineBreak() +
		"###.###.#.###v#####v###" + util.GetLineBreak() +
		"#...#...#.#.>.>.#.>.###" + util.GetLineBreak() +
		"#.###.###.#.###.#.#v###" + util.GetLineBreak() +
		"#.....###...###...#...#" + util.GetLineBreak() +
		"#####################.#"
	want := 154
	got := solve2(input)
	if got != want {
		t.Errorf("wrong result: want=%d, got=%d", want, got)
	}
}
