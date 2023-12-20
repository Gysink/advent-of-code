package main

import (
	"advent-of-code/pkg/util"
	"testing"
)

func Test_Part1(t *testing.T) {
	input := "" +
		"px{a<2006:qkq,m>2090:A,rfg}" + util.GetLineBreak() +
		"pv{a>1716:R,A}" + util.GetLineBreak() +
		"lnx{m>1548:A,A}" + util.GetLineBreak() +
		"rfg{s<537:gd,x>2440:R,A}" + util.GetLineBreak() +
		"qs{s>3448:A,lnx}" + util.GetLineBreak() +
		"qkq{x<1416:A,crn}" + util.GetLineBreak() +
		"crn{x>2662:A,R}" + util.GetLineBreak() +
		"in{s<1351:px,qqz}" + util.GetLineBreak() +
		"qqz{s>2770:qs,m<1801:hdj,R}" + util.GetLineBreak() +
		"gd{a>3333:R,R}" + util.GetLineBreak() +
		"hdj{m>838:A,pv}" + util.GetLineBreak() +
		util.GetLineBreak() +
		"{x=787,m=2655,a=1222,s=2876}" + util.GetLineBreak() +
		"{x=1679,m=44,a=2067,s=496}" + util.GetLineBreak() +
		"{x=2036,m=264,a=79,s=2244}" + util.GetLineBreak() +
		"{x=2461,m=1339,a=466,s=291}" + util.GetLineBreak() +
		"{x=2127,m=1623,a=2188,s=1013}"
	want := 19114
	got := solve1(input)
	if got != want {
		t.Errorf("wrong result: want=%d, got=%d", want, got)
	}
}

// Part 2

func Test_Part2(t *testing.T) {
	input := "" +
		"px{a<2006:qkq,m>2090:A,rfg}" + util.GetLineBreak() +
		"pv{a>1716:R,A}" + util.GetLineBreak() +
		"lnx{m>1548:A,A}" + util.GetLineBreak() +
		"rfg{s<537:gd,x>2440:R,A}" + util.GetLineBreak() +
		"qs{s>3448:A,lnx}" + util.GetLineBreak() +
		"qkq{x<1416:A,crn}" + util.GetLineBreak() +
		"crn{x>2662:A,R}" + util.GetLineBreak() +
		"in{s<1351:px,qqz}" + util.GetLineBreak() +
		"qqz{s>2770:qs,m<1801:hdj,R}" + util.GetLineBreak() +
		"gd{a>3333:R,R}" + util.GetLineBreak() +
		"hdj{m>838:A,pv}" + util.GetLineBreak() +
		util.GetLineBreak() +
		"{x=787,m=2655,a=1222,s=2876}" + util.GetLineBreak() +
		"{x=1679,m=44,a=2067,s=496}" + util.GetLineBreak() +
		"{x=2036,m=264,a=79,s=2244}" + util.GetLineBreak() +
		"{x=2461,m=1339,a=466,s=291}" + util.GetLineBreak() +
		"{x=2127,m=1623,a=2188,s=1013}"
	want := 167409079868000
	got := solve2(input)
	if got != want {
		t.Errorf("wrong result: want=%d, got=%d", want, got)
	}
}
