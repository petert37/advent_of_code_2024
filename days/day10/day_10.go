package day10

import (
	"slices"
	"strconv"
	"strings"
)

func ProcessPart1(input string) string {
	topoMap := parseInput(input)
	trailheadCount := 0
	resultChannel := make(chan int)
	sum := 0
	for y, row := range topoMap {
		for x, num := range row {
			if num == 0 {
				trailheadCount++
				go getTrailheadScore(&topoMap, x, y, true, &resultChannel)
			}
		}
	}
	for i := 0; i < trailheadCount; i++ {
		sum += <-resultChannel
	}
	return strconv.Itoa(sum)
}

func ProcessPart2(input string) string {
	topoMap := parseInput(input)
	trailheadCount := 0
	resultChannel := make(chan int)
	sum := 0
	for y, row := range topoMap {
		for x, num := range row {
			if num == 0 {
				trailheadCount++
				go getTrailheadScore(&topoMap, x, y, false, &resultChannel)
			}
		}
	}
	for i := 0; i < trailheadCount; i++ {
		sum += <-resultChannel
	}
	return strconv.Itoa(sum)
}

type position struct {
	x, y, num int
}

func parseInput(input string) [][]int {
	result := make([][]int, 0)
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		trimmedLine := strings.TrimRight(line, "\r")
		row := make([]int, len(trimmedLine))
		for x, num := range trimmedLine {
			row[x], _ = strconv.Atoi(string(num))
		}
		result = append(result, row)
	}
	return result
}

var directions = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

func getTrailheadScore(topoMap *([][]int), x, y int, distinct bool, result *chan int) {
	fifo := make([]position, 0)
	fifo = append(fifo, position{x, y, 0})
	for {
		if len(fifo) == 0 || fifo[0].num == 9 {
			break
		}
		current := fifo[0]
		fifo = fifo[1:]
		for direction := range directions {
			newX := current.x + directions[direction][0]
			newY := current.y + directions[direction][1]
			if newY >= 0 && newY < len(*topoMap) {
				row := (*topoMap)[newY]
				if newX >= 0 && newX < len(row) {
					num := row[newX]
					if num == current.num+1 {
						newPosition := position{newX, newY, num}
						if distinct && slices.Contains(fifo, newPosition) {
							continue
						}
						fifo = append(fifo, newPosition)
					}
				}
			}
		}
	}
	*result <- len(fifo)
}
