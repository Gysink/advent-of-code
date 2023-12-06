package main

import (
	"reflect"
	"testing"
)

func TestParseInput_Map(t *testing.T) {
	input := "seed-to-soil map:\n" +
		"50 98 2\n" +
		"52 50 48"

	mapper := NewMapper(input)

	assertMappings(t, []*Mapping{
		{
			destination: 50,
			source:      98,
			length:      2,
		},
		{
			destination: 52,
			source:      50,
			length:      48,
		},
	}, mapper.mappings)
}

func TestParseInput(t *testing.T) {
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
		"humidity-to-location map:\n" +
		"60 56 37"

	got := parseAlmanacFromInput(input)

	assertMappers(t, []*Mapper{
		{
			name: "seed-to-soil map:",
			mappings: []*Mapping{
				{
					destination: 50,
					source:      98,
					length:      2,
				},
				{
					destination: 52,
					source:      50,
					length:      48,
				},
			},
		},
		{
			name: "soil-to-fertilizer map:",
			mappings: []*Mapping{
				{
					destination: 0,
					source:      15,
					length:      37,
				},
				{
					destination: 37,
					source:      52,
					length:      2,
				},
				{
					destination: 39,
					source:      0,
					length:      15,
				},
			},
		},
		{
			name: "humidity-to-location map:",
			mappings: []*Mapping{
				{
					destination: 60,
					source:      56,
					length:      37,
				},
			},
		},
	}, got.mappers)
}

// Seed number 79 corresponds to soil number 81.
// seed-to-soil map:
// 52 50 48
func TestMappingSeed(t *testing.T) {
	input := 79
	mapping := &Mapping{
		destination: 52,
		source:      50,
		length:      48,
	}
	want := 81
	got, _ := mapping.mapInput(input)
	assertNum(t, want, got)
}

func TestMappingSeed_Almanac(t *testing.T) {
	a := &Almanac{
		seeds: []int{79},
		mappers: []*Mapper{
			{
				name: "seed-to-soil map:",
				mappings: []*Mapping{
					{
						destination: 50,
						source:      98,
						length:      2,
					},
					{
						destination: 52,
						source:      50,
						length:      48,
					},
				},
			},
		},
	}
	assertNum(t, 81, a.getResult())
}

func assertNum(t *testing.T, want int, got int) {
	t.Helper()
	if got != want {
		t.Errorf("wrong mapped input: want=%d, got=%d", want, got)
	}
}

func assertMappers(t *testing.T, want []*Mapper, got []*Mapper) {
	t.Helper()
	for i, w := range want {
		assertMappings(t, w.mappings, got[i].mappings)
	}
}

func assertMappings(t *testing.T, want []*Mapping, got []*Mapping) {
	t.Helper()
	if !reflect.DeepEqual(want, got) {
		t.Errorf("wrong mappings: want=%v, got=%v", want, got)
	}
}
