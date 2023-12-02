package main

import (
	"reflect"
	"testing"
)

func TestParseId(t *testing.T) {
	input := "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
	want := 1
	got := parseId(input)
	if got != want {
		t.Errorf("wrong parsed ID, got=%d, want=%d", got, want)
	}
}

func TestParseRounds(t *testing.T) {
	input := "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
	want := []Round{
		{red: 4, blue: 3, green: 0},
		{red: 1, blue: 6, green: 2},
		{red: 0, blue: 0, green: 2},
	}
	got := parseRounds(input)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("wrong parsed rounds, got=%v, want=%v", got, want)
	}
}

func TestParseGame(t *testing.T) {
	input := "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
	want := Game{
		id: 1,
		rounds: []Round{
			{red: 4, blue: 3, green: 0},
			{red: 1, blue: 6, green: 2},
			{red: 0, blue: 0, green: 2},
		},
	}
	got := parseGame(input)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("wrong parsed game, got=%v, want=%v", got, want)
	}
}

func TestGamePossible(t *testing.T) {
	input := Round{
		red:   12,
		green: 13,
		blue:  14,
	}
	got := Game{
		id: 1,
		rounds: []Round{
			{red: 4, blue: 3, green: 0},
			{red: 1, blue: 6, green: 2},
			{red: 0, blue: 0, green: 2},
		},
	}.isPossible(input)
	if got != true {
		t.Errorf("game not possible but should be, got=%v, want=%v", got, true)
	}
}

func TestGameNotPossible(t *testing.T) {
	input := Round{
		red:   12,
		green: 13,
		blue:  14,
	}
	got := Game{
		id: 1,
		rounds: []Round{
			{red: 20, blue: 6, green: 8},
			{red: 4, blue: 5, green: 13},
			{red: 1, blue: 0, green: 5},
		},
	}.isPossible(input)
	if got != false {
		t.Errorf("game possible but should not be, got=%v, want=%v", got, true)
	}
}

func TestGetPossibleGamesIdSum(t *testing.T) {
	input := "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green\n" +
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue\n" +
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red\n" +
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red\n" +
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"
	bag := Round{
		red:   12,
		green: 13,
		blue:  14,
	}
	want := 8
	got := getPossibleGamesIdSum(input, bag)
	if got != want {
		t.Errorf("wrong sumg: got=%d, want=%d", got, want)
	}
}

// part 2

func TestFindLowestPossible(t *testing.T) {
	input := Game{
		id: 1,
		rounds: []Round{
			{red: 4, blue: 3, green: 0},
			{red: 1, blue: 6, green: 2},
			{red: 0, blue: 0, green: 2},
		},
	}
	want := Round{
		red:   4,
		green: 2,
		blue:  6,
	}
	got := input.findLowestPossible()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("wrong lowest possible: got=%v, want=%v", got, want)
	}
}

func TestCalcPowerOfLowestPossible(t *testing.T) {
	input := "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green\n" +
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue\n" +
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red\n" +
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red\n" +
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"
	want := 2286
	got := getPowerOfLowestPossibles(input)
	if got != want {
		t.Errorf("wrong power of lowest: got=%d, want=%d", got, want)
	}
}
