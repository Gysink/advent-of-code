package main

import "testing"

func TestScores(t *testing.T) {
	tests := []struct {
		name  string
		input *Hand
		want  int
	}{
		{name: "1", input: parseFromInput("23456 1"), want: 0}, // lowest possible value
		{name: "2", input: parseFromInput("23457 1"), want: 0},
		{name: "3", input: parseFromInput("32T3K 1"), want: 1},
		{name: "4", input: parseFromInput("T55J5 1"), want: 3},
		{name: "5", input: parseFromInput("KK677 1"), want: 2},
		{name: "6", input: parseFromInput("KTJJT 1"), want: 2},
		{name: "7", input: parseFromInput("QQQJA 1"), want: 3},
		{name: "8", input: parseFromInput("AAAAA 1"), want: 6}, // highest possible value
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.input.score
			if got != tt.want {
				t.Errorf("wrong score, got=%d, want=%d", got, tt.want)
			}
		})
	}
}

// part 2

func TestScores2(t *testing.T) {
	tests := []struct {
		name  string
		input *Hand
		want  int
	}{
		{name: "1", input: parseFromInput("23456 1"), want: 0}, // lowest possible value
		{name: "2", input: parseFromInput("JJJJJ 1"), want: 6},
		{name: "3", input: parseFromInput("J8JJJ 1"), want: 6},
		{name: "4", input: parseFromInput("J82JJ 1"), want: 5},
		{name: "5", input: parseFromInput("J8288 1"), want: 5},
		{name: "6", input: parseFromInput("J828J 1"), want: 5},
		{name: "7", input: parseFromInput("J88JJ 1"), want: 6},
		{name: "8", input: parseFromInput("AAAAA 1"), want: 6}, // highest possible value
		{name: "8", input: parseFromInput("J5A7J 1"), want: 3},
		{name: "8", input: parseFromInput("J5352 1"), want: 3},
		{name: "8", input: parseFromInput("T55JT 1"), want: 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.input.score
			if got != tt.want {
				t.Errorf("wrong score, got=%d, want=%d", got, tt.want)
			}
		})
	}
}
