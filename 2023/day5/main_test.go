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

func TestMapping_Result(t *testing.T) {
	input := "seeds: 79 14 55 13\n" +
		"\n" +
		"seed-to-soil map:\n" +
		"50 98 2\n" +
		"52 50 48\n" +
		"\n" +
		"soil-to-fertilizer map:\n" +
		"0 15 37\n" +
		"37 52 2\n" +
		"39 0 15\n" +
		"\n" +
		"fertilizer-to-water map:\n" +
		"49 53 8\n" +
		"0 11 42\n" +
		"42 0 7\n" +
		"57 7 4\n" +
		"\n" +
		"water-to-light map:\n" +
		"88 18 7\n" +
		"18 25 70\n" +
		"\n" +
		"light-to-temperature map:\n" +
		"45 77 23\n" +
		"81 45 19\n" +
		"68 64 13\n" +
		"\n" +
		"temperature-to-humidity map:\n" +
		"0 69 1\n" +
		"1 0 69\n" +
		"\n" +
		"humidity-to-location map:\n" +
		"60 56 37\n" +
		"56 93 4"
	a := parseAlmanacFromInput(input)
	assertNum(t, 35, a.getResult())
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

// part 2

func TestParseInput_AlmanacSeeds_Part2(t *testing.T) {
	input := "seeds: 79 2 55 3"
	almanac := NewAlmanac(input)
	assertSeeds(t, []int{79, 80, 55, 56, 57}, almanac.seeds)
}

func TestMapping_Result2(t *testing.T) {
	input := "seeds: 79 14 55 13\n" +
		"\n" +
		"seed-to-soil map:\n" +
		"50 98 2\n" +
		"52 50 48\n" +
		"\n" +
		"soil-to-fertilizer map:\n" +
		"0 15 37\n" +
		"37 52 2\n" +
		"39 0 15\n" +
		"\n" +
		"fertilizer-to-water map:\n" +
		"49 53 8\n" +
		"0 11 42\n" +
		"42 0 7\n" +
		"57 7 4\n" +
		"\n" +
		"water-to-light map:\n" +
		"88 18 7\n" +
		"18 25 70\n" +
		"\n" +
		"light-to-temperature map:\n" +
		"45 77 23\n" +
		"81 45 19\n" +
		"68 64 13\n" +
		"\n" +
		"temperature-to-humidity map:\n" +
		"0 69 1\n" +
		"1 0 69\n" +
		"\n" +
		"humidity-to-location map:\n" +
		"60 56 37\n" +
		"56 93 4"
	a := parseAlmanacFromInput(input)
	assertNum(t, 35, a.getResult())
}
