package main

import (
	"reflect"
	"testing"
)

func TestFindValues_OnlyNumbers(t *testing.T) {
	input := "467..114.."
	want := []ValueCord{
		{value: "467", startX: 0, endX: 3, line: 0},
		{value: "114", startX: 5, endX: 8, line: 0},
	}
	got := parseValues(input, 0)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("wrong parsed values, got=%v, want=%v", got, want)
	}
}

func TestFindValues_OnlySymbols(t *testing.T) {
	input := ".....*...."
	want := []ValueCord{
		{value: "*", startX: 5, endX: 6, line: 0},
	}
	got := parseValues(input, 0)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("wrong parsed values, got=%v, want=%v", got, want)
	}
}

func TestFindValues_NumbersAndSymbols(t *testing.T) {
	input := ".34..*...."
	want := []ValueCord{
		{value: "34", startX: 1, endX: 3, line: 0},
		{value: "*", startX: 5, endX: 6, line: 0},
	}
	got := parseValues(input, 0)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("wrong parsed values, got=%v, want=%v", got, want)
	}
}

func TestFindValues_NumberNextToSymbol(t *testing.T) {
	input := ".344*....."
	want := []ValueCord{
		{value: "344", startX: 1, endX: 4, line: 0},
		{value: "*", startX: 4, endX: 5, line: 0},
	}
	got := parseValues(input, 0)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("wrong parsed values, got=%v, want=%v", got, want)
	}
}

func TestFindValues_MultipleLines(t *testing.T) {
	input := "467..114..\n" +
		"...*......\n" +
		"..35..633."
	want := []ValueCord{
		{value: "467", startX: 0, endX: 3, line: 0},
		{value: "114", startX: 5, endX: 8, line: 0},
		{value: "*", startX: 3, endX: 4, line: 1},
		{value: "35", startX: 2, endX: 4, line: 2},
		{value: "633", startX: 6, endX: 9, line: 2},
	}
	got := parseInput(input)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("wrong parsed values, got=%v, want=%v", got, want)
	}
}

func TestIsNumber_Number(t *testing.T) {
	input := ValueCord{value: "34", startX: 3, endX: 4, line: 1}
	if input.isNum() != true {
		t.Errorf("wrong type: should be a number")
	}
}

func TestIsNumber_Symbol(t *testing.T) {
	input := ValueCord{value: "*", startX: 3, endX: 4, line: 1}
	if input.isNum() != false {
		t.Errorf("wrong type: should not be a number")
	}
}

func TestGetNeighboursOnXAxis_Left(t *testing.T) {
	// 467.*44..
	values := []ValueCord{
		{value: "467", startX: 0, endX: 3, line: 0},
		{value: "*", startX: 4, endX: 5, line: 0},
		{value: "44", startX: 5, endX: 7, line: 0},
	}
	input := ValueCord{value: "*", startX: 4, endX: 5, line: 0}

	want := []ValueCord{
		{value: "44", startX: 5, endX: 7, line: 0},
	}
	got := input.getNumNeighbours(values)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("wrong neighbours: got=%v, want=%v", got, want)
	}
}

func TestGetNeighboursOnXAxis_Right(t *testing.T) {
	// 467*.44..
	values := []ValueCord{
		{value: "467", startX: 0, endX: 3, line: 0},
		{value: "*", startX: 3, endX: 4, line: 0},
		{value: "44", startX: 5, endX: 7, line: 0},
	}
	input := ValueCord{value: "*", startX: 3, endX: 4, line: 0}

	want := []ValueCord{
		{value: "467", startX: 0, endX: 3, line: 0},
	}
	got := input.getNumNeighbours(values)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("wrong neighbours: got=%v, want=%v", got, want)
	}
}

func TestGetNeighboursOnYAxis_Top(t *testing.T) {
	// 467..44..
	// ..*......
	values := []ValueCord{
		{value: "467", startX: 0, endX: 3, line: 0},
		{value: "44", startX: 5, endX: 7, line: 0},
		{value: "*", startX: 2, endX: 3, line: 1},
	}
	input := ValueCord{value: "*", startX: 2, endX: 3, line: 1}

	want := []ValueCord{
		{value: "467", startX: 0, endX: 3, line: 0},
	}
	got := input.getNumNeighbours(values)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("wrong neighbours: got=%v, want=%v", got, want)
	}
}

func TestGetNeighboursOnYAxis_Bottom(t *testing.T) {
	// ..*......
	// 467..44..
	values := []ValueCord{
		{value: "*", startX: 2, endX: 3, line: 0},
		{value: "467", startX: 0, endX: 3, line: 1},
		{value: "44", startX: 5, endX: 7, line: 1},
	}
	input := ValueCord{value: "*", startX: 2, endX: 3, line: 0}

	want := []ValueCord{
		{value: "467", startX: 0, endX: 3, line: 1},
	}
	got := input.getNumNeighbours(values)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("wrong neighbours: got=%v, want=%v", got, want)
	}
}

func TestGetNeighboursOnYAxis_Diagonal(t *testing.T) {
	// 467..44..
	// ...*.....
	values := []ValueCord{
		{value: "467", startX: 0, endX: 3, line: 0},
		{value: "44", startX: 5, endX: 7, line: 0},
		{value: "*", startX: 3, endX: 4, line: 1},
	}
	input := ValueCord{value: "*", startX: 3, endX: 4, line: 1}

	want := []ValueCord{
		{value: "467", startX: 0, endX: 3, line: 0},
	}
	got := input.getNumNeighbours(values)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("wrong neighbours: got=%v, want=%v", got, want)
	}
}

func TestGetSumOfEngineParts(t *testing.T) {
	input := "467..114..\n" +
		"...*......\n" +
		"..35..633.\n" +
		"......#...\n" +
		"617*......\n" +
		".....+.58.\n" +
		"..592.....\n" +
		"......755.\n" +
		"...$.*....\n" +
		".664.598.."
	want := 4361
	got := calcEnginePartsSum(input)
	if got != want {
		t.Errorf("wrong sum: got=%d, want=%d", got, want)
	}
}

func TestGetSumOfEngineParts_Variant(t *testing.T) {
	input :=
		"....33+58.\n" +
			"..592.....\n" +
			"......755.\n" +
			"...$.*....\n" +
			".664.598.."
	want := 2108
	got := calcEnginePartsSum(input)
	if got != want {
		t.Errorf("wrong sum: got=%d, want=%d", got, want)
	}
}

func TestGetSumOfEngineParts_Long(t *testing.T) {
	input := "" +
		".....137.........340................961.................383*.......295........................506...........*.....................*....*....\n" + // 383 207 622 30 216
		"......../..670..............52.....*.......802..950.........207..........334..377.625.....871...*........622..160.130..#.....295.30...216...\n" + // 137 961 34 506 20 804
		".............*............=.*.......34.941.=.......*............&.12......*....*...........*...20..356...........*......804.*..............." // 670 52 802 950 334 377 871 160 130 295
	want := 8561
	got := calcEnginePartsSum(input)
	if got != want {
		t.Errorf("wrong sum: got=%d, want=%d", got, want)
	}
}

func TestGetSumOfEngineParts_Variant2(t *testing.T) {
	input := "" +
		"..........\n" +
		"...59.22..\n" +
		".....*....\n" +
		"...2...3..\n" +
		".........."
	want := 81
	got := calcEnginePartsSum(input)
	if got != want {
		t.Errorf("wrong sum: got=%d, want=%d", got, want)
	}
}

func TestGetSumOfEngineParts_Variant3(t *testing.T) {
	input := "" +
		"..........\n" +
		"..59...22.\n" +
		".....*....\n" +
		"....2.3...\n" +
		".........."
	want := 5
	got := calcEnginePartsSum(input)
	if got != want {
		t.Errorf("wrong sum: got=%d, want=%d", got, want)
	}
}

func TestGetSumOfEngineParts_Variant4(t *testing.T) {
	input := "" +
		"..........\n" +
		"...11.10..\n" +
		"....2*4...\n" +
		"....2.3...\n" +
		".........."
	want := 32
	got := calcEnginePartsSum(input)
	if got != want {
		t.Errorf("wrong sum: got=%d, want=%d", got, want)
	}
}

func TestGetSumOfEngineParts_Variant5(t *testing.T) {
	input := "" +
		"..........\n" +
		"...11.22..\n" +
		".....*/...\n" +
		".........."
	want := 55
	got := calcEnginePartsSum(input)
	if got != want {
		t.Errorf("wrong sum: got=%d, want=%d", got, want)
	}
}

func TestGetSumOfEngineParts_Variant6(t *testing.T) {
	input := "" +
		"12.......*..\n" + // 34
		"+.........34\n" + // 12
		".......-12..\n" + // 12
		"..78........\n" +
		"..*....60...\n" + // 78 78
		"78..........\n" +
		".......23...\n" +
		"....90*12...\n" + // 90 23 13
		"............\n" +
		"2.2......12.\n" +
		".*.........*\n" + // 2 2 1 1  12 56
		"1.1.......56"
	want := 413
	got := calcEnginePartsSum(input)
	if got != want {
		t.Errorf("wrong sum: got=%d, want=%d", got, want)
	}
}

func TestGetSumOfEngineParts_Variant7(t *testing.T) {
	input := "" +
		"12.......*..\n" + // 34
		"+.........34" // 12

	want := 46
	got := calcEnginePartsSum(input)
	if got != want {
		t.Errorf("wrong sum: got=%d, want=%d", got, want)
	}
}

// part 2

func TestGetSumOfEngineParts_Part2(t *testing.T) {
	input := "" +
		"467..114..\n" +
		"...*......\n" + // 467 * 35 = 16345
		"..35..633.\n" +
		"......#...\n" + // 633
		"617*......\n" + // 617
		".....+.58.\n" + // 592
		"..592.....\n" +
		"......755.\n" +
		"...$.*....\n" + // 664  755 * 598 = 451490
		".664.598.."
	want := 467835
	got := calcEnginePartsSum_Part2(input)
	if got != want {
		t.Errorf("wrong sum: got=%d, want=%d", got, want)
	}
}
