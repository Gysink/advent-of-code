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
