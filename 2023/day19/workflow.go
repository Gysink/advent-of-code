package main

import (
	"advent-of-code/pkg/util"
	"strings"
)

type Workflow struct {
	name         string
	elseWorkflow string
	rules        []Rule
}

func (w Workflow) process(p Part) string {
	for _, r := range w.rules {
		nextWorkflow := r.process(p)
		switch nextWorkflow {
		case "A":
			return "A"
		case "R":
			return "R"
		case "":
			continue
		default:
			return nextWorkflow
		}
	}
	return w.elseWorkflow
}

func NewWorkflow(input string) Workflow {
	fi := strings.Index(input, "{")
	li := strings.Index(input, "}")
	name := input[:fi]
	allRules := input[fi+1 : li]

	var rules []Rule
	elseWorkflow := ""
	ruleDescriptions := strings.Split(allRules, ",")
	for i, rd := range ruleDescriptions {
		if i == len(ruleDescriptions)-1 {
			elseWorkflow = rd
			break
		}
		rules = append(rules, NewRule(rd))
	}

	return Workflow{
		name:         name,
		elseWorkflow: elseWorkflow,
		rules:        rules,
	}
}

type Rule struct {
	input          string
	operand        string
	num            int
	targetWorkflow Workflow
	target         string
}

func (r Rule) process(p Part) string {
	in := 0

	switch r.input {
	case "x":
		in = p.x
	case "m":
		in = p.m
	case "a":
		in = p.a
	case "s":
		in = p.s
	}

	next := ""
	switch r.operand {
	case "<":
		if in < r.num {
			next = r.target
		}
	case ">":
		if in > r.num {
			next = r.target
		}
	}
	return next
}

func NewRule(input string) Rule {
	in := input[:1]
	operand := input[1:2]
	nt := strings.Split(input[2:], ":")
	num := nt[0]
	target := nt[1]

	return Rule{
		input:   in,
		operand: operand,
		num:     util.StringToInt(num),
		target:  target,
	}
}
