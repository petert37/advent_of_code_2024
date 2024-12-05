package day05

import (
	"slices"
	"strconv"
	"strings"

	"github.com/dominikbraun/graph"
)

func ProcessPart1(input string) string {
	rules, updates := parseInput(input)

	sum := 0

	for _, update := range updates {
		if isUpdateValid(&update, &rules) {
			midIndex := len(update) / 2
			sum += update[midIndex]
		}
	}

	return strconv.Itoa(sum)
}

func parseInput(input string) (map[int][]int, [][]int) {
	lines := strings.Split(input, "\n")

	rules := make(map[int][]int)
	updates := make([][]int, 0)

	updatesStart := 0
	for i, line := range lines {
		trimmedLine := strings.TrimRight(line, "\r")
		if trimmedLine == "" {
			updatesStart = i + 1
			break
		}
		parts := strings.Split(trimmedLine, "|")
		first, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		second, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		currentRules, exists := rules[first]
		if exists {
			rules[first] = append(currentRules, second)
		} else {
			rules[first] = []int{second}
		}
	}

	for i := updatesStart; i < len(lines); i++ {
		trimmedLine := strings.TrimRight(lines[i], "\r")
		parts := strings.Split(trimmedLine, ",")
		update := make([]int, len(parts))
		for j, part := range parts {
			parsed, err := strconv.Atoi(part)
			if err != nil {
				panic(err)
			}
			update[j] = parsed
		}
		updates = append(updates, update)
	}

	return rules, updates
}

func isUpdateValid(update *[]int, rules *map[int][]int) bool {
	for i, num := range *update {
		invalidNumbers := (*rules)[num]
		for j := range i {
			if slices.Contains(invalidNumbers, (*update)[j]) {
				return false
			}
		}
	}
	return true
}

func ProcessPart2(input string) string {
	rules, updates := parseInput(input)

	sum := 0

	for _, update := range updates {
		if !isUpdateValid(&update, &rules) {
			reordered := reorderUpdate(&update, &rules)
			midIndex := len(reordered) / 2
			sum += reordered[midIndex]
		}
	}

	return strconv.Itoa(sum)
}

func reorderUpdate(update *[]int, rules *map[int][]int) []int {
	g := graph.New(graph.IntHash, graph.Directed(), graph.Acyclic())
	for _, v := range *update {
		g.AddVertex(v)
	}
	for key, values := range *rules {
		_, err := g.Vertex(key)
		if err != nil {
			continue
		}
		for _, value := range values {
			_, err := g.Vertex(value)
			if err != nil {
				continue
			}
			g.AddEdge(key, value)
		}
	}
	ordered, _ := graph.TopologicalSort(g)

	return ordered
}
