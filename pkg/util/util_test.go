package util

import (
	"reflect"
	"testing"
)

func TestSplitByLine(t *testing.T) {
	input := "these\n" +
		"are\n" +
		"some\n" +
		"lines\n"
	want := []string{"these", "are", "some", "lines"}
	got := SplitByLine(input)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("wrong splitByLine, got=%s, want=%s", got, want)
	}
}

func TestStringToInt(t *testing.T) {
	input := "12"
	want := 12
	got := StringToInt(input)
	if got != want {
		t.Errorf("wrong convertion, got=%d, want=%d", got, want)
	}
}
