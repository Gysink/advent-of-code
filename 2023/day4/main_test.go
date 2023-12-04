package main

import (
	"reflect"
	"testing"
)

func TestNewCard_ResultIsZero(t *testing.T) {
	card := NewCard("")
	assertResult(t, 0, card.getResult())
}

func TestNewCard_FromInput_ParseNumbers(t *testing.T) {
	input := "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"
	card := NewCard(input)
	assertNumbers(t, []int{41, 48, 83, 86, 17}, card.winningNumbers)
	assertNumbers(t, []int{83, 86, 6, 31, 17, 9, 48, 53}, card.cardNumbers)
}

func TestCard_OnXWinningAndYCard_ResultIsZero(t *testing.T) {
	card := &Card{
		winningNumbers: []int{1},
		cardNumbers:    []int{2},
	}
	assertResult(t, 0, card.getResult())
}

func TestCard_OnXWinningAndXCard_ResultIsOne(t *testing.T) {
	card := &Card{
		winningNumbers: []int{1},
		cardNumbers:    []int{1},
	}
	assertResult(t, 1, card.getResult())
}

func TestCard_On2Winning_ResultIsTwo(t *testing.T) {
	card := &Card{
		winningNumbers: []int{1, 2},
		cardNumbers:    []int{1, 2},
	}
	assertResult(t, 2, card.getResult())
}

func TestCard_On4Winning_ResultIsEight(t *testing.T) {
	card := &Card{
		winningNumbers: []int{1, 2, 3, 4},
		cardNumbers:    []int{1, 2, 3, 4},
	}
	assertResult(t, 8, card.getResult())
}

func TestFromInput_CalcPoints(t *testing.T) {
	input := "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53\n" +
		"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19\n" +
		"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1\n" +
		"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83\n" +
		"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36\n" +
		"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11"
	assertResult(t, 13, calcPoints(input))
}

func assertNumbers(t *testing.T, want []int, got []int) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("wrong numbers: got=%v, want=%v", got, want)
	}
}

func assertResult(t *testing.T, want int, got int) {
	t.Helper()
	if want != got {
		t.Errorf("wrong result: got=%d, want=%d", got, want)
	}
}

// part 2

//func TestFromInput_CalcTotalCards_One(t *testing.T) {
//	input := "Card 1: 87 83 26 28 32 | 87 30 70 12 93 22 82 36\n" + // 1 matching number
//		"Card 2: 31 18 13 56 72 | 74 77 10 23 35 67 36 11" // 0 matching number
//	assertResult(t, 3, calcTotalCards(input))
//}
//
//func TestFromInput_CalcTotalCards(t *testing.T) {
//	input := "Card 1: 13 32 20 16 61 | 61 30 68 82 17 32 24 19\n" + // 2 matching numbers
//		"Card 2: 87 83 26 28 32 | 87 30 70 12 93 22 82 36\n" + // 1 matching number
//		"Card 3: 31 18 13 56 72 | 74 77 10 23 35 67 36 11" // 0 matching number
//	assertResult(t, 7, calcTotalCards(input))
//}
//
//func TestFromInput_CalcTotalCards22(t *testing.T) {
//	input := "Card 1: 1 2 | 1 2\n" + // 2 matching numbers
//		"Card 2: 1 | 1\n" + // 1 matching number
//		"Card 3: 1 | 1\n" + // 1 matching number
//		"Card 4: 1 | 2" // 0 matching number
//	assertResult(t, 12, calcTotalCards(input))
//}
//
//func TestFromInput_CalcTotalCards2(t *testing.T) {
//	input := "Card 1: 1 2 3 | 1 2 3\n" + // 3 matching numbers
//		"Card 2: 1 2 | 1 2\n" + // 2 matching number
//		"Card 3: 1 | 1\n" + // 1 matching number
//		"Card 4: 72 | 74" // 1 matching number
//	assertResult(t, 14, calcTotalCards(input))
//}
//
////func TestFromInput_CalcTotalCards_30(t *testing.T) {
////	input := "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53\n" +
////		"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19\n" +
////		"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1\n" +
////		"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83\n" +
////		"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36\n" +
////		"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11"
////	assertResult(t, 30, calcTotalCards(input))
////}
