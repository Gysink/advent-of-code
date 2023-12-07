package main

import (
	"reflect"
	"testing"
)

func TestPoker(t *testing.T) {
	input := "" +
		"32T3K 765\n" +
		"T55J5 684\n" +
		"KK677 28\n" +
		"KTJJT 220\n" +
		"QQQJA 483\n" +
		""
	want := 6440
	got := solve1(input)
	if want != got {
		t.Errorf("wrong result: want=%d, got=%d", want, got)
	}
}

func TestNewCard(t *testing.T) {
	input := "A"
	want := &Card{value: "A", score: 14}
	got := NewCard(input)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("wrong card: want=%v, got=%v", want, got)
	}
}

func TestNewHand(t *testing.T) {
	input := "T55J5 684"
	want := &Hand{
		cards: []*Card{
			{value: "T", score: 10},
			{value: "5", score: 5},
			{value: "5", score: 5},
			{value: "J", score: 11},
			{value: "5", score: 5},
		},
		bid:   684,
		value: "T55J5",
		score: 3,
	}
	got := NewHand(input)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("wrong hand: want=%v, got=%v", want, got)
	}
}

// Part 2

func TestPoker_Part2(t *testing.T) {
	input := "" +
		"32T3K 765\n" +
		"T55J5 684\n" +
		"KK677 28\n" +
		"KTJJT 220\n" +
		"QQQJA 483\n" +
		""
	want := 5905
	got := solve2(input)
	if want != got {
		t.Errorf("wrong result: want=%d, got=%d", want, got)
	}
}
