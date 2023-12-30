package main

import (
	"advent-of-code/pkg/util"
	"log"
	"os"
	"sort"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	res := solve1(string(data))
	log.Println("result for day25:", res)
}

func solve1(input string) int {
	graph := buildGraph(input)
	edges := findEdges(graph)

	first, second, third := getTop3Edges(edges)
	sp := strings.Split(first, "-")
	graph[sp[0]] = remove(graph[sp[0]], sp[1])
	graph[sp[1]] = remove(graph[sp[1]], sp[0])
	sp = strings.Split(second, "-")
	graph[sp[0]] = remove(graph[sp[0]], sp[1])
	graph[sp[1]] = remove(graph[sp[1]], sp[0])
	sp = strings.Split(third, "-")
	graph[sp[0]] = remove(graph[sp[0]], sp[1])
	graph[sp[1]] = remove(graph[sp[1]], sp[0])

	start := ""
	for k, _ := range graph {
		start = k
		break
	}
	queue := NewQueue()
	queue.push(start)
	seen := make(map[string]bool)
	seen[start] = true
	for queue.len() > 0 {
		node := queue.pop()
		for _, nx := range graph[node] {
			if _, ok := seen[nx]; !ok {
				queue.push(nx)
				seen[nx] = true
			}
		}
	}
	a := len(seen)
	b := len(graph) - a
	return a * b
}

func findEdges(graph Graph) map[string]int {
	edges := make(map[string]int)
	for start, _ := range graph {
		queue := NewQueue()
		queue.push(start)
		seen := make(map[string]bool)
		seen[start] = true
		prev := make(map[string]string)
		for queue.len() > 0 {
			node := queue.pop()
			for _, nx := range graph[node] {
				if _, ok := seen[nx]; !ok {
					seen[nx] = true
					queue.push(nx)
					prev[nx] = node
				}
			}
		}
		for node, _ := range graph {
			for node != start {
				nx := prev[node]
				edges[min(nx, node)+"-"+max(nx, node)] += 1
				node = nx
			}
		}
	}
	return edges
}

func buildGraph(input string) Graph {
	lines := util.SplitByLine(input)
	graph := make(map[string][]string)
	for _, l := range lines {
		sp := strings.Split(l, ":")
		links := strings.Split(sp[1][1:], " ")
		if _, ok := graph[sp[0]]; !ok {
			graph[sp[0]] = make([]string, 0)
		}
		for _, link := range links {
			graph[sp[0]] = append(graph[sp[0]], link)
			if _, ok := graph[link]; !ok {
				graph[link] = make([]string, 0)
			}
			graph[link] = append(graph[link], sp[0])
		}
	}
	return graph
}

func remove(arr []string, s string) []string {
	index := 0
	for i, v := range arr {
		if v == s {
			index = i
		}
	}
	res := arr[:index]
	return append(res, arr[index+1:]...)
}

type Tuple struct {
	val   string
	count int
}

func getTop3Edges(edges map[string]int) (string, string, string) {
	var all []Tuple
	for k, v := range edges {
		all = append(all, Tuple{k, v})
	}
	sort.Slice(all, func(i, j int) bool {
		return all[i].count > all[j].count
	})
	return all[0].val, all[1].val, all[2].val
}

func min(a, b string) string {
	if a < b {
		return a
	}
	return b
}

func max(a, b string) string {
	if a > b {
		return a
	}
	return b
}

type Graph map[string][]string
