package day19

import (
	"strconv"
	"strings"
)

func ProcessPart1(input string) string {
	towels, patterns := parseInput(input)
	count := 0
	impossiblePatterns := make(map[string]bool)
	for _, pattern := range patterns {
		if isPatternPossible(pattern, &towels, &impossiblePatterns) {
			count++
		} else {
		}
	}
	return strconv.Itoa(count)
}

func ProcessPart2(input string) string {
	towels, patterns := parseInput(input)
	count := 0
	possibilities := make(map[string]int)
	for _, pattern := range patterns {
		patternCount := countPossibilities(pattern, &towels, &possibilities)
		count += patternCount
	}
	return strconv.Itoa(count)
}

func parseInput(input string) (towels, patterns []string) {
	towels = make([]string, 0)
	lines := strings.Split(input, "\n")
	isPatterns := false
	for _, line := range lines {
		trimmedLine := strings.TrimRight(line, "\r")
		if isPatterns {
			patterns = append(patterns, trimmedLine)
		} else {
			if trimmedLine == "" {
				isPatterns = true
			} else {
				towels = strings.Split(trimmedLine, ", ")
			}
		}
	}
	return
}

func isPatternPossible(pattern string, towels *[]string, impossiblePatterns *map[string]bool) bool {
	if pattern == "" {
		return true
	}
	if (*impossiblePatterns)[pattern] {
		return false
	}
	possible := false
	for _, towel := range *towels {
		if strings.HasPrefix(pattern, towel) {
			next := pattern[len(towel):]
			nextPossible := isPatternPossible(next, towels, impossiblePatterns)
			if nextPossible {
				possible = true
				break
			}
		}
	}
	if !possible {
		(*impossiblePatterns)[pattern] = true
	}
	return possible
}

func countPossibilities(pattern string, towels *[]string, possibilities *map[string]int) int {
	if pattern == "" {
		return 1
	}
	if val, ok := (*possibilities)[pattern]; ok {
		return val
	}
	count := 0
	for _, towel := range *towels {
		if strings.HasPrefix(pattern, towel) {
			next := pattern[len(towel):]
			count += countPossibilities(next, towels, possibilities)
		}
	}
	(*possibilities)[pattern] = count
	return count
}
