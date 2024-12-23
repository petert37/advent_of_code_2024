package day23

import (
	"slices"
	"strconv"
	"strings"
)

func ProcessPart1(input string) string {
	parsedInput := parseInput(input)
	edges := make(map[string][]string)
	for _, parts := range parsedInput {
		current, ok := edges[parts[0]]
		if !ok {
			edges[parts[0]] = []string{parts[1]}
		} else {
			edges[parts[0]] = append(current, parts[1])
		}
		current, ok = edges[parts[1]]
		if !ok {
			edges[parts[1]] = []string{parts[0]}
		} else {
			edges[parts[1]] = append(current, parts[0])
		}
	}
	triplets := make(map[string]bool)
	for v1, e1 := range edges { //v1 = aq
		for _, v2 := range e1 { //v2 = cg
			e2 := edges[v2]
			for _, v3 := range e2 { // v3 = yn
				for _, v0 := range e1 { // v0 = yn
					if v0 == v3 {
						vert := []string{v1, v2, v3}
						slices.Sort(vert)
						triplet := strings.Join(vert, "_")
						triplets[triplet] = true
					}
				}
			}
		}
	}
	count := 0
	for key := range triplets {
		if key[0] == 't' || strings.Contains(key, "_t") {
			count++
		}
	}
	return strconv.Itoa(count)
}

func ProcessPart2(input string) string {
	parsedInput := parseInput(input)
	edges := make(map[string][]string)
	for _, parts := range parsedInput {
		current, ok := edges[parts[0]]
		if !ok {
			edges[parts[0]] = []string{parts[1]}
		} else {
			edges[parts[0]] = append(current, parts[1])
		}
		current, ok = edges[parts[1]]
		if !ok {
			edges[parts[1]] = []string{parts[0]}
		} else {
			edges[parts[1]] = append(current, parts[0])
		}
	}
	c := maxClique(&edges)
	slices.Sort(c)
	return strings.Join(c, ",")
}

func parseInput(input string) [][]string {
	result := make([][]string, 0)
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		trimmedLine := strings.TrimRight(line, "\r")
		parts := strings.Split(trimmedLine, "-")
		result = append(result, parts)
	}
	return result
}

func maxClique(edges *map[string][]string) []string {
	p := make([]string, 0)
	for key := range *edges {
		p = append(p, key)
	}
	results := make([][]string, 0)
	bronKerbosch([]string{}, p, []string{}, edges, &results)
	maxLength := 0
	var result []string
	for _, r := range results {
		if len(r) > maxLength {
			maxLength = len(r)
			result = r
		}
	}
	return result
}

func bronKerbosch(r, p, x []string, edges *map[string][]string, results *[][]string) {
	if len(p) == 0 && len(x) == 0 {
		*results = append(*results, r)
		return
	}
	currentP := make([]string, len(p))
	copy(currentP, p)
	currentX := make([]string, len(x))
	copy(currentX, x)
	for _, v := range p {
		nextR := make([]string, len(r))
		copy(nextR, r)
		nextR = append(nextR, v)
		nextP := make([]string, 0)
		for _, pV := range currentP {
			if slices.Contains((*edges)[v], pV) {
				nextP = append(nextP, pV)
			}
		}
		nextX := make([]string, 0)
		for _, xV := range currentX {
			if slices.Contains((*edges)[v], xV) {
				nextX = append(nextX, xV)
			}
		}
		bronKerbosch(nextR, nextP, nextX, edges, results)
		currentPUpdated := make([]string, 0)
		for _, pV := range currentP {
			if pV != v {
				currentPUpdated = append(currentPUpdated, pV)
			}
		}
		currentP = currentPUpdated
		currentX = append(currentX, v)
	}
}
