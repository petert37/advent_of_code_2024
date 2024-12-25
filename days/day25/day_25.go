package day25

import (
	"strconv"
	"strings"
)

func ProcessPart1(input string) string {
	locks, keys := parseInput(input)
	count := 0
	for _, key := range keys {
		for _, lock := range locks {
			match := true
			for i := 0; i < 5; i++ {
				if key[i]+lock[i] > 5 {
					match = false
					break
				}
			}
			if match {
				count++
			}
		}
	}
	return strconv.Itoa(count)
}

func parseInput(input string) (locks [][5]int, keys [][5]int) {
	locks = make([][5]int, 0)
	keys = make([][5]int, 0)
	lines := strings.Split(input, "\n")
	for i := 0; i < len(lines); i += 8 {
		isLock := true
		if lines[i][0] == '.' {
			isLock = false
		}
		start := 0
		end := 6
		if isLock {
			start = 1
			end = 7
		}
		sizes := [5]int{}
		for j := start; j < end; j++ {
			for k := 0; k < 5; k++ {
				if lines[i+j][k] == '#' {
					sizes[k]++
				}
			}
		}
		if isLock {
			locks = append(locks, sizes)
		} else {
			keys = append(keys, sizes)
		}
	}
	return
}
