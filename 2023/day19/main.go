package main

import (
	"advent-of-code/pkg/util"
	"log"
	"math"
	"os"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	res := solve1(string(data))
	log.Println("result for day19:", res)

	res2 := solve2(string(data))
	log.Println("result for day19 part 2:", res2)
}

func solve1(input string) int {
	sp := strings.Split(input, util.GetLineBreak()+util.GetLineBreak())
	workflows := parseWorkflows(sp[0])
	parts := parseParts(sp[1])
	return processParts(parts, workflows)
}

func processParts(parts []Part, workflows map[string]Workflow) int {
	sum := 0
	for _, part := range parts {
		res := part.process(workflows)
		if res == "A" {
			sum += part.getSumRatings()
		}
	}
	return sum
}

func parseParts(input string) []Part {
	lines := strings.Split(input, util.GetLineBreak())

	var parts []Part
	for _, l := range lines {
		parts = append(parts, NewPart(l))
	}
	return parts
}

func parseWorkflows(input string) map[string]Workflow {
	lines := util.SplitByLine(input)

	workflows := make(map[string]Workflow)
	for _, l := range lines {
		fi := strings.Index(l, "{")
		name := l[:fi]
		workflows[name] = NewWorkflow(l)
	}
	return workflows
}

// Part 2

type Range struct {
	start, end int
}

func solve2(input string) int {
	sp := strings.Split(input, util.GetLineBreak()+util.GetLineBreak())
	workflows := parseWorkflows(sp[0])
	ranges := map[string]*Range{
		"x": {1, 4000},
		"m": {1, 4000},
		"a": {1, 4000},
		"s": {1, 4000}}
	return processRanges(ranges, "in", workflows)
}

func processRanges(ranges map[string]*Range, name string, workflows map[string]Workflow) int {
	if name == "R" {
		return 0
	}
	if name == "A" {
		return rangesProduct(ranges)
	}

	hasProcessedAll := false
	result := 0
	workflow := workflows[name]
	for _, rule := range workflow.rules {
		low := ranges[rule.input].start
		high := ranges[rule.input].end
		var trueRange *Range
		var falseRange *Range

		switch rule.operand {
		case "<":
			trueRange = &Range{start: low, end: min(rule.num-1, high)}
			falseRange = &Range{start: max(rule.num, low), end: high}
		case ">":
			trueRange = &Range{start: max(rule.num+1, low), end: high}
			falseRange = &Range{start: low, end: min(rule.num, high)}
		}

		if trueRange.start <= trueRange.end {
			newRanges := copyMap(ranges)
			newRanges[rule.input] = trueRange
			result += processRanges(newRanges, rule.target, workflows)
		}
		if falseRange.start <= falseRange.end {
			ranges = copyMap(ranges)
			ranges[rule.input] = falseRange
		} else {
			hasProcessedAll = true
			break
		}
	}
	if !hasProcessedAll {
		result += processRanges(ranges, workflow.elseWorkflow, workflows)
	}
	return result
}

func min(a int, b int) int {
	return int(math.Min(float64(a), float64(b)))
}

func max(a int, b int) int {
	return int(math.Max(float64(a), float64(b)))
}

func copyMap(ranges map[string]*Range) map[string]*Range {
	c := make(map[string]*Range, len(ranges))
	for k, v := range ranges {
		c[k] = &Range{
			start: v.start,
			end:   v.end,
		}
	}
	return c
}

func rangesProduct(ranges map[string]*Range) int {
	result := 1
	for _, r := range ranges {
		result *= r.end - r.start + 1
	}
	return result
}
