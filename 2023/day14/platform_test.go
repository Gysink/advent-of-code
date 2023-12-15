package main

import (
	"advent-of-code/pkg/util"
	"reflect"
	"testing"
)

func TestNewPlatform(t *testing.T) {
	input := "" +
		"O.." + util.GetLineBreak() +
		"O.#"

	want := Platform{grid: [][]Element{
		{
			{value: 'O', movable: true, blocking: true, x: 0, y: 0},
			{value: '.', movable: false, blocking: false, x: 0, y: 1},
			{value: '.', movable: false, blocking: false, x: 0, y: 2},
		},
		{
			{value: 'O', movable: true, blocking: true, x: 1, y: 0},
			{value: '.', movable: false, blocking: false, x: 1, y: 1},
			{value: '#', movable: false, blocking: true, x: 1, y: 2},
		},
	}}
	got := NewPlatform(input)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("wron platform: want=%v, got=%v", want, got)
	}
}

func TestPlatform_TiltNorth(t *testing.T) {
	input := "" +
		"O....#...." + util.GetLineBreak() +
		"O.OO#....#" + util.GetLineBreak() +
		".....##..." + util.GetLineBreak() +
		"OO.#O....O" + util.GetLineBreak() +
		".O.....O#." + util.GetLineBreak() +
		"O.#..O.#.#" + util.GetLineBreak() +
		"..O..#O..O" + util.GetLineBreak() +
		".......O.." + util.GetLineBreak() +
		"#....###.." + util.GetLineBreak() +
		"#OO..#...."
	want := "" +
		"OOOO.#.O.." + util.GetLineBreak() +
		"OO..#....#" + util.GetLineBreak() +
		"OO..O##..O" + util.GetLineBreak() +
		"O..#.OO..." + util.GetLineBreak() +
		"........#." + util.GetLineBreak() +
		"..#....#.#" + util.GetLineBreak() +
		"..O..#.O.O" + util.GetLineBreak() +
		"..O......." + util.GetLineBreak() +
		"#....###.." + util.GetLineBreak() +
		"#....#...." + util.GetLineBreak()
	p := NewPlatform(input)
	p.tiltNorth()
	got := p.String()
	if want != got {
		t.Errorf("wrong result: want=\n%s, got=\n%s", want, got)
	}
}

func TestPlatform_TiltWest(t *testing.T) {
	input := "" +
		"OOOO.#.O.." + util.GetLineBreak() +
		"OO..#....#" + util.GetLineBreak() +
		"OO..O##..O" + util.GetLineBreak() +
		"O..#.OO..." + util.GetLineBreak() +
		"........#." + util.GetLineBreak() +
		"..#....#.#" + util.GetLineBreak() +
		"..O..#.O.O" + util.GetLineBreak() +
		"..O......." + util.GetLineBreak() +
		"#....###.." + util.GetLineBreak() +
		"#....#...."
	want := "" +
		"OOOO.#O..." + util.GetLineBreak() +
		"OO..#....#" + util.GetLineBreak() +
		"OOO..##O.." + util.GetLineBreak() +
		"O..#OO...." + util.GetLineBreak() +
		"........#." + util.GetLineBreak() +
		"..#....#.#" + util.GetLineBreak() +
		"O....#OO.." + util.GetLineBreak() +
		"O........." + util.GetLineBreak() +
		"#....###.." + util.GetLineBreak() +
		"#....#...." + util.GetLineBreak()
	p := NewPlatform(input)
	p.tiltWest()
	got := p.String()
	if want != got {
		t.Errorf("wrong result: want=\n[%s], \ngot=\n[%s]", want, got)
	}
}

func TestPlatform_TiltSouth(t *testing.T) {
	input := "" +
		"OOOO.#O..." + util.GetLineBreak() +
		"OO..#....#" + util.GetLineBreak() +
		"OOO..##O.." + util.GetLineBreak() +
		"O..#OO...." + util.GetLineBreak() +
		"........#." + util.GetLineBreak() +
		"..#....#.#" + util.GetLineBreak() +
		"O....#OO.." + util.GetLineBreak() +
		"O........." + util.GetLineBreak() +
		"#....###.." + util.GetLineBreak() +
		"#....#...."
	want := "" +
		".....#...." + util.GetLineBreak() +
		"....#.O..#" + util.GetLineBreak() +
		"O..O.##..." + util.GetLineBreak() +
		"O.O#......" + util.GetLineBreak() +
		"O.O....O#." + util.GetLineBreak() +
		"O.#..O.#.#" + util.GetLineBreak() +
		"O....#...." + util.GetLineBreak() +
		"OO....OO.." + util.GetLineBreak() +
		"#O...###.." + util.GetLineBreak() +
		"#O..O#...." + util.GetLineBreak()
	p := NewPlatform(input)
	p.tiltSouth()
	got := p.String()
	if want != got {
		t.Errorf("wrong result: want=\n[%s], \ngot=\n[%s]", want, got)
	}
}

func TestPlatform_TiltEast(t *testing.T) {
	input := "" +
		".....#...." + util.GetLineBreak() +
		"....#.O..#" + util.GetLineBreak() +
		"O..O.##..." + util.GetLineBreak() +
		"O.O#......" + util.GetLineBreak() +
		"O.O....O#." + util.GetLineBreak() +
		"O.#..O.#.#" + util.GetLineBreak() +
		"O....#...." + util.GetLineBreak() +
		"OO....OO.." + util.GetLineBreak() +
		"#O...###.." + util.GetLineBreak() +
		"#O..O#...."
	want := "" +
		".....#...." + util.GetLineBreak() +
		"....#...O#" + util.GetLineBreak() +
		"...OO##..." + util.GetLineBreak() +
		".OO#......" + util.GetLineBreak() +
		".....OOO#." + util.GetLineBreak() +
		".O#...O#.#" + util.GetLineBreak() +
		"....O#...." + util.GetLineBreak() +
		"......OOOO" + util.GetLineBreak() +
		"#...O###.." + util.GetLineBreak() +
		"#..OO#...." + util.GetLineBreak()
	p := NewPlatform(input)
	p.tiltEast()
	got := p.String()
	if want != got {
		t.Errorf("wrong result: want=\n[%s], \ngot=\n[%s]", want, got)
	}
}
