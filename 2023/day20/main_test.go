package main

import (
	"advent-of-code/pkg/util"
	"testing"
)

func Test_Part1(t *testing.T) {
	input := "" +
		"broadcaster -> a, b, c" + util.GetLineBreak() +
		"%a -> b" + util.GetLineBreak() +
		"%b -> c" + util.GetLineBreak() +
		"%c -> inv" + util.GetLineBreak() +
		"&inv -> a"
	want := 32_000_000
	got := solve1(input)
	if got != want {
		t.Errorf("wrong result: want=%d, got=%d", want, got)
	}
}

func Test_Part1_2(t *testing.T) {
	input := "" +
		"broadcaster -> a" + util.GetLineBreak() +
		"%a -> inv, con" + util.GetLineBreak() +
		"&inv -> b" + util.GetLineBreak() +
		"%b -> con" + util.GetLineBreak() +
		"&con -> output"
	want := 11_687_500
	got := solve1(input)
	if got != want {
		t.Errorf("wrong result: want=%d, got=%d", want, got)
	}
}
