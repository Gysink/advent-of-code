package main

import "log"

type Point struct {
	x      int
	y      int
	symbol string
	viewed bool
}

func NewPoint(c string, x, y int) Point {
	return Point{
		symbol: c,
		x:      x,
		y:      y,
		viewed: false,
	}
}

func (p *Point) next(points [][]Point, prevPoint Point) Point {
	if points[p.x][p.y].viewed {
		log.Println("this point was already viewed:", p.symbol, " at ", p.x, ",", p.y)
	}
	points[p.x][p.y].viewed = true

	switch p.symbol {
	case "|":
		if prevPoint.x > p.x {
			return points[p.x-1][p.y]
		} else {
			return points[p.x+1][p.y]
		}
	case "-":
		if prevPoint.y > p.y {
			return points[p.x][p.y-1]
		} else {
			return points[p.x][p.y+1]
		}
	case "L":
		if prevPoint.x == p.x {
			return points[p.x-1][p.y]
		} else {
			return points[p.x][p.y+1]
		}
	case "J":
		if prevPoint.x == p.x {
			return points[p.x-1][p.y]
		} else {
			return points[p.x][p.y-1]
		}
	case "7":
		if prevPoint.x == p.x {
			return points[p.x+1][p.y]
		} else {
			return points[p.x][p.y-1]
		}
	case "F":
		if prevPoint.x == p.x {
			return points[p.x+1][p.y]
		} else {
			return points[p.x][p.y+1]
		}
	case "S":
		nextRight := points[p.x][p.y+1]
		if nextRight.symbol == "-" || nextRight.symbol == "7" || nextRight.symbol == "J" {
			return nextRight
		}
		return points[p.x+1][p.y]
	default:
		log.Fatalln("should not reach an end")
		return Point{}
	}
}
