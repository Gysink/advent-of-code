package main

const (
	space              = '.'
	mirror             = '/'
	mirror2            = '\\'
	verticalSplitter   = '|'
	horizontalSplitter = '-'
)

type Point struct {
	x, y        int
	value       rune
	isEnergized bool
}

func NewPoint(x, y int, value rune) Point {
	return Point{
		x:           x,
		y:           y,
		value:       value,
		isEnergized: false,
	}
}
