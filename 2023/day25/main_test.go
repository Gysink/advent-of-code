package main

import (
	"advent-of-code/pkg/util"
	"testing"
)

func Test_Part1(t *testing.T) {
	input := "" +
		"jqt: rhn xhk nvd" + util.GetLineBreak() +
		"rsh: frs pzl lsr" + util.GetLineBreak() +
		"xhk: hfx" + util.GetLineBreak() +
		"cmg: qnr nvd lhk bvb" + util.GetLineBreak() +
		"rhn: xhk bvb hfx" + util.GetLineBreak() +
		"bvb: xhk hfx" + util.GetLineBreak() +
		"pzl: lsr hfx nvd" + util.GetLineBreak() +
		"qnr: nvd" + util.GetLineBreak() +
		"ntq: jqt hfx bvb xhk" + util.GetLineBreak() +
		"nvd: lhk" + util.GetLineBreak() +
		"lsr: lhk" + util.GetLineBreak() +
		"rzs: qnr cmg lsr rsh" + util.GetLineBreak() +
		"frs: qnr lhk lsr"
	want := 54
	got := solve1(input)
	if got != want {
		t.Errorf("wrong result: want=%d, got=%d", want, got)
	}
}
