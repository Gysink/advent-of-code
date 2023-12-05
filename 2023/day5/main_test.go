package main

import (
	"reflect"
	"testing"
)

func TestNewAlmanac_EmptyResult(t *testing.T) {
	almanac := NewAlmanac("")
	assertResult(t, 0, almanac.getResult())
}

func TestParseInput_AlmanacSeeds(t *testing.T) {
	input := "seeds: 79 14 55 13"
	almanac := NewAlmanac(input)
	assertSeeds(t, []int{79, 14, 55, 13}, almanac.seeds)
}

func assertSeeds(t *testing.T, want []int, got []int) {
	t.Helper()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("wrong seeds: want=%v, got=%v", want, got)
	}
}

func assertResult(t *testing.T, want int, got int) {
	t.Helper()
	if want != got {
		t.Errorf("wrong result: want=%d, got=%d", want, got)
	}
}
