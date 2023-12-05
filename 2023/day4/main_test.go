package main

import (
	"reflect"
	"testing"
)

func TestNewCard_ResultIsZero(t *testing.T) {
	card := NewCard("")
	assertInt(t, 0, card.getScore())
}

func TestNewCard_FromInput_ParseNumbers(t *testing.T) {
	input := "41 48 83 86 17 | 83 86  6 31 17  9 48 53"
	wn, cn := parseNumbers(input)
	assertNumbers(t, []int{41, 48, 83, 86, 17}, wn)
	assertNumbers(t, []int{83, 86, 6, 31, 17, 9, 48, 53}, cn)
}

func TestNewCard_FromInput_ParseCardId(t *testing.T) {
	input := "Card  1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"
	card := NewCard(input)
	assertInt(t, 1, card.id)
}

func TestCard_OnXWinningAndYCard_ResultIsZero(t *testing.T) {
	card := &Card{
		wins: 0,
	}
	assertInt(t, 0, card.getScore())
}

func TestCard_OnXWinningAndXCard_ResultIsOne(t *testing.T) {
	card := &Card{
		wins: 1,
	}
	assertInt(t, 1, card.getScore())
}

func TestCard_On2Winning_ResultIsTwo(t *testing.T) {
	card := &Card{
		wins: 2,
	}
	assertInt(t, 2, card.getScore())
}

func TestCard_On4Winning_ResultIsEight(t *testing.T) {
	card := &Card{
		wins: 4,
	}
	assertInt(t, 8, card.getScore())
}

func TestFromInput_CalcPoints(t *testing.T) {
	input := "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53\n" +
		"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19\n" +
		"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1\n" +
		"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83\n" +
		"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36\n" +
		"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11"
	assertInt(t, 13, calcPoints(input))
}

func assertNumbers(t *testing.T, want []int, got []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("wrong numbers: got=%v, want=%v", got, want)
	}
}

func assertInt(t *testing.T, want int, got int) {
	t.Helper()
	if want != got {
		t.Errorf("wrong result: got=%d, want=%d", got, want)
	}
}

// part 2

func TestFromInput_CalcTotalCards_2_0(t *testing.T) {
	input := "Card 1: 87 83 26 28 32 | 87 30 70 12 93 22 82 36\n" + // 1 matching number
		"Card 2: 31 18 13 56 72 | 74 77 10 23 35 67 36 11" // 0 matching number
	assertInt(t, 3, calcTotalCards(input))
}

func TestFromInput_CalcTotalCards_2_1_0(t *testing.T) {
	input := "Card 1: 13 32 20 16 61 | 61 30 68 82 17 32 24 19\n" + // 2 matching numbers
		"Card 2: 87 83 26 28 32 | 87 30 70 12 93 22 82 36\n" + // 1 matching number
		"Card 3: 31 18 13 56 72 | 74 77 10 23 35 67 36 11" // 0 matching number
	assertInt(t, 7, calcTotalCards(input))
}

func TestFromInput_CalcTotalCards_2_1_1_0(t *testing.T) {
	input := "Card 1: 1 2 | 1 2\n" + // 2 matching numbers
		"Card 2: 1 | 1\n" + // 1 matching number
		"Card 3: 1 | 1\n" + // 1 matching number
		"Card 4: 1 | 2" // 0 matching number
	assertInt(t, 12, calcTotalCards(input))
}

func TestFromInput_CalcTotalCards_3_2_1_0(t *testing.T) {
	input := "Card 1: 1 2 3 | 1 2 3\n" + // 3 matching numbers
		"Card 2: 1 2 | 1 2\n" + // 2 matching number
		"Card 3: 1 | 1\n" + // 1 matching number
		"Card 4: 1 | 99" // 0 matching number
	assertInt(t, 15, calcTotalCards(input))
}

func TestFromInput_CalcTotalCards_30(t *testing.T) {
	input := "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53\n" +
		"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19\n" +
		"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1\n" +
		"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83\n" +
		"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36\n" +
		"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11"
	assertInt(t, 30, calcTotalCards(input))
}
