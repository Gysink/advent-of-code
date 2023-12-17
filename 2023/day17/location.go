package main

type Direction int

const (
	Up Direction = iota
	Down
	Left
	Right
)

// Turn turns left or right.
func (d Direction) Turn(turn Direction) Direction {
	if turn != Left && turn != Right {
		panic("should be left or right")
	}

	switch d {
	case Up:
		return turn
	case Down:
		switch turn {
		case Left:
			return Right
		case Right:
			return Left
		}
	case Left:
		switch turn {
		case Left:
			return Down
		case Right:
			return Up
		}
	case Right:
		switch turn {
		case Left:
			return Up
		case Right:
			return Down
		}
	}

	panic("not handled")
}

type Position struct {
	X, Y int
}

func NewPosition(x, y int) Position {
	return Position{X: x, Y: y}
}

// Delta returns a new position from a row delta and col delta.
func (p Position) Delta(row, col int) Position {
	return Position{
		X: p.X + row,
		Y: p.Y + col,
	}
}

// Move moves into a given direction and a certain number of times.
func (p Position) Move(direction Direction, moves int) Position {
	switch direction {
	case Up:
		return p.Delta(-moves, 0)
	case Down:
		return p.Delta(moves, 0)
	case Left:
		return p.Delta(0, -moves)
	case Right:
		return p.Delta(0, moves)
	}

	panic("not handled")
}

type Location struct {
	Pos Position
	Dir Direction
}

func NewLocation(row, col int, dir Direction) Location {
	return Location{
		Pos: NewPosition(row, col),
		Dir: dir,
	}
}

// Turn turns left or right.
func (l Location) Turn(d Direction, moves int) Location {
	dir := l.Dir.Turn(d)
	pos := l.Pos.Move(dir, moves)
	return Location{Pos: pos, Dir: dir}
}

// Straight moves in the current direction.
func (l Location) Straight(moves int) Location {
	pos := l.Pos.Move(l.Dir, moves)
	return Location{Pos: pos, Dir: l.Dir}
}
