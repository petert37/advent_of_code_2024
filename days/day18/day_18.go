package day18

import (
	"strconv"
	"strings"
)

func ProcessPart1(input string, size, bytes int) string {
	positions := parseInput(input)[:bytes]
	memory := make([][]bool, size+1)
	for i := 0; i < size+1; i++ {
		memory[i] = make([]bool, size+1)
	}
	for _, pos := range positions {
		memory[pos.y][pos.x] = true
	}

	visited := make(map[position]int)
	fifo := make([]step, 0)
	fifo = append(fifo, step{position{0, 0}, 0})
	visited[position{0, 0}] = 0
	var minDistance *int = nil

	for len(fifo) > 0 {
		current := fifo[0]
		fifo = fifo[1:]
		if current.position.x == size && current.position.y == size {
			if minDistance == nil || current.distance < *minDistance {
				minDistance = &current.distance
			}
		}
		for _, dir := range directions {
			nextX := current.position.x + dir[0]
			nextY := current.position.y + dir[1]
			nextPostion := position{nextX, nextY}
			if nextX >= 0 && nextX <= size && nextY >= 0 && nextY <= size && !memory[nextY][nextX] {
				prev, exists := visited[nextPostion]
				if !exists || prev > current.distance+1 {
					visited[nextPostion] = current.distance + 1
					fifo = append(fifo, step{nextPostion, current.distance + 1})
				}
			}
		}
	}

	return strconv.Itoa(*minDistance)
}

func ProcessPart2(input string, size, bytes int) string {
	positions := parseInput(input)
	for i := bytes + 1; i < len(positions); i++ {
		currentPositions := positions[:i]
		memory := make([][]bool, size+1)
		for i := 0; i < size+1; i++ {
			memory[i] = make([]bool, size+1)
		}
		for _, pos := range currentPositions {
			memory[pos.y][pos.x] = true
		}

		visited := make(map[position]int)
		fifo := make([]step, 0)
		fifo = append(fifo, step{position{0, 0}, 0})
		visited[position{0, 0}] = 0

		pathFound := false
		for len(fifo) > 0 {
			current := fifo[0]
			fifo = fifo[1:]
			if current.position.x == size && current.position.y == size {
				pathFound = true
				break
			}
			for _, dir := range directions {
				nextX := current.position.x + dir[0]
				nextY := current.position.y + dir[1]
				nextPostion := position{nextX, nextY}
				if nextX >= 0 && nextX <= size && nextY >= 0 && nextY <= size && !memory[nextY][nextX] {
					prev, exists := visited[nextPostion]
					if !exists || prev > current.distance+1 {
						visited[nextPostion] = current.distance + 1
						fifo = append(fifo, step{nextPostion, current.distance + 1})
					}
				}
			}
		}
		if !pathFound {
			return strconv.Itoa(currentPositions[len(currentPositions)-1].x) + "," + strconv.Itoa(currentPositions[len(currentPositions)-1].y)
		}
	}
	return "???"
}

type position struct {
	x, y int
}

type step struct {
	position position
	distance int
}

var directions = [][2]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}

func parseInput(input string) []position {
	result := make([]position, 0)
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		trimmedLine := strings.TrimRight(line, "\r")
		parts := strings.Split(trimmedLine, ",")
		x, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		y, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		result = append(result, position{x, y})
	}
	return result
}
