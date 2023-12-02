package main

type Game struct {
	id     int
	rounds []Round
}

func (g Game) isPossible(input Round) bool {
	for _, r := range g.rounds {
		if !r.isWithin(input) {
			return false
		}
	}
	return true
}

func (g Game) findLowestPossible() Round {
	result := Round{}
	for _, r := range g.rounds {
		if result.red < r.red {
			result.red = r.red
		}
		if result.green < r.green {
			result.green = r.green
		}
		if result.blue < r.blue {
			result.blue = r.blue
		}
	}
	return result
}

type Round struct {
	red   int
	green int
	blue  int
}

func (r Round) isWithin(input Round) bool {
	if r.red > input.red {
		return false
	}
	if r.blue > input.blue {
		return false
	}
	if r.green > input.green {
		return false
	}
	return true
}

func (r Round) getPower() int {
	return r.red * r.blue * r.green
}
