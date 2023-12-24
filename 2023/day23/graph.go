package main

type Graph struct {
	nodes map[Point]map[Point]int
}

func NewGraph(points []Point) Graph {
	g := Graph{make(map[Point]map[Point]int)}
	for _, p := range points {
		g.set(p, make(map[Point]int))
	}
	return g
}

func (g Graph) get(point Point) map[Point]int {
	return g.nodes[point]
}

func (g Graph) has(point Point) bool {
	_, ok := g.nodes[point]
	return ok
}

func (g Graph) set(point Point, targets map[Point]int) {
	g.nodes[point] = targets
}

func (g Graph) len() int {
	return len(g.nodes)
}

func (g Graph) setSteps(from Point, to Point, steps int) {
	g.nodes[from][to] = steps
}
