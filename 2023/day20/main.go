package main

import (
	"advent-of-code/pkg/util"
	"log"
	"os"
	"reflect"
	"strings"
)

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatalln(err)
	}

	res := solve1(string(data))
	log.Println("result for day20:", res)

	res2 := solve2(string(data))
	log.Println("result for day20 part 2:", res2)
}

func solve1(input string) int {
	lines := util.SplitByLine(input)
	modules := make(map[string]Module)

	for _, l := range lines {
		sp := strings.Split(l, " -> ")
		mod := parseModule(sp[0])
		modules[mod.GetName()] = mod
	}

	for _, l := range lines {
		sp := strings.Split(l, " -> ")
		name := parseName(sp[0])

		var targetModules []Module
		targets := strings.Split(sp[1], ", ")
		for _, t := range targets {
			if _, ok := modules[t]; ok {
				targetModules = append(targetModules, modules[t])
				if reflect.TypeOf(modules[t]).String() == "*main.Conjunction" {
					modules[t].SetOrigin(modules[name])
				}
			} else {
				modules[t] = NoopModule{name: t}
				targetModules = append(targetModules, modules[t])
			}
		}
		modules[name].SetTargets(targetModules)
	}

	totalHighs := 0
	totalLows := 0
	buttonPressed := 1000
	for i := 0; i < buttonPressed; i++ {
		queue := NewQueue()
		broadCaster := modules["broadcaster"]
		queue.push(NewQueueItem(nil, broadCaster, Low))

		for queue.len() > 0 {
			item := queue.pop()
			if item.value {
				totalHighs++
			} else {
				totalLows++
			}
			item.target.Send(item.origin, item.value, queue)
		}
	}
	return totalHighs * totalLows
}

func parseModule(input string) Module {
	if strings.HasPrefix(input, "%") {
		return &FlipFlop{name: input[1:]}
	}
	if strings.HasPrefix(input, "&") {
		return &Conjunction{name: input[1:]}
	}

	if input == "broadcaster" {
		return &BroadcasterModule{}
	}
	return &NoopModule{name: input}
}

func parseName(input string) string {
	switch input[0] {
	case '%':
		return input[1:]
	case '&':
		return input[1:]
	}
	return input
}

// Part 2

func solve2(input string) int {
	lines := util.SplitByLine(input)
	modules := make(map[string]Module)

	for _, l := range lines {
		sp := strings.Split(l, " -> ")
		mod := parseModule(sp[0])
		modules[mod.GetName()] = mod
	}

	for _, l := range lines {
		sp := strings.Split(l, " -> ")
		name := parseName(sp[0])

		var targetModules []Module
		targets := strings.Split(sp[1], ", ")
		for _, t := range targets {
			if _, ok := modules[t]; ok {
				targetModules = append(targetModules, modules[t])
				if reflect.TypeOf(modules[t]).String() == "*main.Conjunction" {
					modules[t].SetOrigin(modules[name])
				}
			} else {
				modules[t] = NoopModule{name: t}
				targetModules = append(targetModules, modules[t])
			}
		}
		modules[name].SetTargets(targetModules)
	}

	feed := modules["gf"]
	cycleLengths := make(map[string]int)
	seen := make(map[string]int)
	buttonPressed := 0
	for {
		buttonPressed++
		queue := NewQueue()
		broadCaster := modules["broadcaster"]
		queue.push(NewQueueItem(nil, broadCaster, Low))

		for queue.len() > 0 {
			item := queue.pop()
			item.target.Send(item.origin, item.value, queue)

			if item.target.GetName() == feed.GetName() && item.value {
				seen[item.origin.GetName()] += 1
				if _, ok := cycleLengths[item.origin.GetName()]; !ok {
					cycleLengths[item.origin.GetName()] = buttonPressed
				} else {
					if buttonPressed != seen[item.origin.GetName()]*cycleLengths[item.origin.GetName()] {
						log.Fatalln("assertion failed")
					}
				}
				if allSeenXTimes(seen, 10) {
					return calcLcm(cycleLengths)
				}
			}
		}
	}
	return buttonPressed
}

func calcLcm(cycleLengths map[string]int) int {
	var lengths []int
	for _, v := range cycleLengths {
		lengths = append(lengths, v)
	}

	sum := float64(lengths[0])
	for i := 1; i < len(lengths); i = i + 1 {
		sum = util.CalculateLCM(sum, float64(lengths[i]))
	}
	return int(sum)
}

func allSeenXTimes(seen map[string]int, i int) bool {
	for _, v := range seen {
		if v < i {
			return false
		}
	}
	return true
}
