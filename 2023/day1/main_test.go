package main

import (
	"testing"
)

func TestGetFirstNum(t *testing.T) {
	input := "1abc2"
	want := "1"
	got := getFirstNum(input)
	if got != want {
		t.Errorf("wrong first num, got=%s want=%s", got, want)
	}
}

func TestGetLastNum(t *testing.T) {
	input := "1abc2"
	want := "2"
	got := getLastNum(input)
	if got != want {
		t.Errorf("wrong first num, got=%s want=%s", got, want)
	}
}

func TestStringToInt(t *testing.T) {
	input := "12"
	want := 12
	got := stringToInt(input)
	if got != want {
		t.Errorf("wrong convertion, got=%d, want=%d", got, want)
	}
}

func TestGetCalibrationOfLine(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{name: "first", input: "1abc2", want: 12},
		{name: "second", input: "pqr3stu8vwx", want: 38},
		{name: "multiple numbers", input: "a1b2c3d4e5f", want: 15},
		{name: "one number", input: "treb7uchet", want: 77},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getCalibrationOfLine(tt.input)
			if got != tt.want {
				t.Errorf("wrong calibration, got=%d, want=%d", got, tt.want)
			}
		})
	}
}

func TestGetSumOfCalibrations(t *testing.T) {
	input := "1abc2\npqr3stu8vwx\na1b2c3d4e5f\ntreb7uchet"
	want := 142
	got := getSumOfCalibrations(input)
	if got != want {
		t.Errorf("wrong sum, got=%d, want=%d", got, want)
	}
}

// part 2

func TestGetRealFirstNum(t *testing.T) {
	input := "two1nine"
	want := "2"
	got := getFirstNum(input)
	if got != want {
		t.Errorf("wrong first real num, got=%s want=%s", got, want)
	}
}

func TestGetRealLastNum(t *testing.T) {
	input := "two1nine"
	want := "9"
	got := getLastNum(input)
	if got != want {
		t.Errorf("wrong first real num, got=%s want=%s", got, want)
	}
}

func TestGetRealCalibrationOfLine(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{name: "1", input: "two1nine", want: 29},
		{name: "2", input: "eightwothree", want: 83},
		{name: "3", input: "abcone2threexyz", want: 13},
		{name: "4", input: "xtwone3four", want: 24},
		{name: "5", input: "4nineeightseven2", want: 42},
		{name: "6", input: "zoneight234", want: 14},
		{name: "7", input: "7pqrstsixteen", want: 76},
		{name: "8", input: "479", want: 49},
		{name: "9", input: "4threerrqxpmmzhrqht4two2", want: 42},
		{name: "10", input: "threetwo1drtzsixtwofourppvg", want: 34},
		{name: "11", input: "abc3abc", want: 33},
		{name: "12", input: "fh5", want: 55},
		{name: "13", input: "4one5twoxqfcflcbjqsixeightdlknnzdbzlrqfkhvm", want: 48},
		{name: "14", input: "44kmn", want: 44},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getCalibrationOfLine(tt.input)
			if got != tt.want {
				t.Errorf("wrong calibration, got=%d, want=%d", got, tt.want)
			}
		})
	}
}

func TestGetRealSumOfCalibrations(t *testing.T) {
	input := "two1nine\neightwothree\nabcone2threexyz\nxtwone3four\n4nineeightseven2\nzoneight234\n7pqrstsixteen"
	want := 281
	got := getSumOfCalibrations(input)
	if got != want {
		t.Errorf("wrong sum, got=%d, want=%d", got, want)
	}
}
