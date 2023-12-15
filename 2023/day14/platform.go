package main

import (
	"advent-of-code/pkg/util"
)

const RoundRock = 'O'
const FlatRock = '#'
const Space = '.'

type Platform struct {
	grid [][]Element
}

func NewPlatform(input string) Platform {
	p := Platform{}
	lines := util.SplitByLine(input)

	p.grid = make([][]Element, len(lines))
	for x, l := range lines {
		p.grid[x] = make([]Element, len(l))
		for y, r := range l {
			p.grid[x][y] = NewElement(r, x, y)
		}
	}
	return p
}

func (p *Platform) tiltNorth() {
	rocks := p.getAllRocks()

	for _, r := range rocks {
		newX := p.findOpenSpotNorth(r)
		if newX != r.x {
			p.set(p.get(newX, r.y), r.x, r.y)
			p.set(r, newX, r.y)
		}
	}
}

func (p *Platform) tiltSouth() {
	rocks := p.getAllRocksReversed()

	for _, r := range rocks {
		newX := p.findOpenSpotSouth(r)
		if newX != r.x {
			p.set(p.get(newX, r.y), r.x, r.y)
			p.set(r, newX, r.y)
		}
	}
}

func (p *Platform) tiltWest() {
	rocks := p.getAllRocks()

	for _, r := range rocks {
		newY := p.findOpenSpotWest(r)
		if newY != r.y {
			p.set(p.get(r.x, newY), r.x, r.y)
			p.set(r, r.x, newY)
		}
	}
}

func (p *Platform) tiltEast() {
	rocks := p.getAllRocksReversed()

	for _, r := range rocks {
		newY := p.findOpenSpotEast(r)
		if newY != r.y {
			p.set(p.get(r.x, newY), r.x, r.y)
			p.set(r, r.x, newY)
		}
	}
}

func (p *Platform) findOpenSpotNorth(r Element) int {
	x := r.x
	for i := x - 1; i >= 0; i-- {
		if p.get(i, r.y).blocking {
			break
		}
		x = i
	}
	return x
}

func (p *Platform) findOpenSpotSouth(r Element) int {
	x := r.x
	for i := x + 1; i < len(p.grid); i++ {
		if p.get(i, r.y).blocking {
			break
		}
		x = i
	}
	return x
}

func (p *Platform) findOpenSpotWest(r Element) int {
	y := r.y
	for i := y - 1; i >= 0; i-- {
		if p.get(r.x, i).blocking {
			break
		}
		y = i
	}
	return y
}

func (p *Platform) findOpenSpotEast(r Element) int {
	y := r.y
	for i := y + 1; i < len(p.grid[0]); i++ {
		if p.get(r.x, i).blocking {
			break
		}
		y = i
	}
	return y
}

func (p *Platform) getAllRocks() []Element {
	var elements []Element
	for _, l := range p.grid {
		for _, e := range l {
			if e.value == RoundRock {
				elements = append(elements, e)
			}
		}
	}
	return elements
}

func (p *Platform) getAllRocksReversed() []Element {
	var elements []Element
	for i := len(p.grid) - 1; i >= 0; i-- {
		for j := len(p.grid[i]) - 1; j >= 0; j-- {
			if p.grid[i][j].value == RoundRock {
				elements = append(elements, p.grid[i][j])
			}
		}
	}
	return elements
}

func (p *Platform) get(x, y int) Element {
	return p.grid[x][y]
}

func (p *Platform) set(e Element, x, y int) {
	e.x = x
	e.y = y
	p.grid[x][y] = e
}

func (p *Platform) String() string {
	result := ""
	for _, l := range p.grid {
		for _, e := range l {
			result += string(e.value)
		}
		result += util.GetLineBreak()
	}
	return result
}

func (p *Platform) Step() {
	p.tiltNorth()
	p.tiltWest()
	p.tiltSouth()
	p.tiltEast()
}

// Element

func NewElement(val rune, x, y int) Element {
	movable := false
	blocking := false
	if val == RoundRock {
		movable = true
		blocking = true
	}
	if val == FlatRock {
		blocking = true
	}

	return Element{
		movable:  movable,
		blocking: blocking,
		value:    val,
		x:        x,
		y:        y,
	}
}

type Element struct {
	movable  bool
	blocking bool
	value    rune
	x        int
	y        int
}

func (e Element) String() string {
	return string(e.value)
}
