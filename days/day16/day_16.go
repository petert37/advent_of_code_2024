package day16

import (
	"strconv"
	"strings"
)

func ProcessPart1(input string) string {
	tiles, start, end := parseInput(input)
	var minCost *int = nil
	stack := make([]path, 0)
	visited := make(map[step]int)
	stack = append(stack, path{step{start, east}, 0})
	visited[step{start, east}] = 0

	for len(stack) > 0 {
		current := stack[0]
		stack = stack[1:]

		if current.s.pos == end {
			if minCost == nil || current.cost < *minCost {
				minCost = &current.cost
			}
			continue
		}

		movePosition := move(current.s.pos, current.s.dir)
		if movePosition.y >= 0 && movePosition.y < len(tiles) {
			row := tiles[movePosition.y]
			if movePosition.x >= 0 && movePosition.x < len(row) {
				if row[movePosition.x] == empty {
					movePath := path{step{movePosition, current.s.dir}, current.cost + 1}
					visitedCost, ok := visited[movePath.s]
					if !ok || movePath.cost < visitedCost {
						visited[movePath.s] = movePath.cost
						stack = append(stack, movePath)
					}
				}
			}
		}

		turnLeftPath := path{step{current.s.pos, turn(current.s.dir, true)}, current.cost + 1000}
		visitedCost, ok := visited[turnLeftPath.s]
		if !ok || turnLeftPath.cost < visitedCost {
			visited[turnLeftPath.s] = turnLeftPath.cost
			stack = append(stack, turnLeftPath)
		}

		turnRightPath := path{step{current.s.pos, turn(current.s.dir, false)}, current.cost + 1000}
		visitedCost, ok = visited[turnRightPath.s]
		if !ok || turnRightPath.cost < visitedCost {
			visited[turnRightPath.s] = turnRightPath.cost
			stack = append(stack, turnRightPath)
		}

	}

	return strconv.Itoa(*minCost)
}

func ProcessPart2(input string) string {
	tiles, start, end := parseInput(input)
	var minCost *int = nil
	minPaths := make([][]position, 0)
	stack := make([]path2, 0)
	visited := make(map[step]int)
	stack = append(stack, path2{step{start, east}, []position{start}, 0})
	visited[step{start, east}] = 0

	for len(stack) > 0 {
		current := stack[0]
		stack = stack[1:]

		if current.s.pos == end {
			if minCost == nil || current.cost < *minCost {
				minCost = &current.cost
				minPaths = [][]position{current.p}
			} else if current.cost == *minCost {
				minPaths = append(minPaths, current.p)
			}
			continue
		}

		movePosition := move(current.s.pos, current.s.dir)
		if movePosition.y >= 0 && movePosition.y < len(tiles) {
			row := tiles[movePosition.y]
			if movePosition.x >= 0 && movePosition.x < len(row) {
				if row[movePosition.x] == empty {
					p := make([]position, len(current.p))
					copy(p, current.p)
					p = append(p, movePosition)
					movePath := path2{step{movePosition, current.s.dir}, p, current.cost + 1}
					visitedCost, ok := visited[movePath.s]
					if !ok || movePath.cost <= visitedCost {
						visited[movePath.s] = movePath.cost
						stack = append(stack, movePath)
					}
				}
			}
		}

		turnLeftPath := path2{step{current.s.pos, turn(current.s.dir, true)}, current.p, current.cost + 1000}
		visitedCost, ok := visited[turnLeftPath.s]
		if !ok || turnLeftPath.cost <= visitedCost {
			visited[turnLeftPath.s] = turnLeftPath.cost
			stack = append(stack, turnLeftPath)
		}

		turnRightPath := path2{step{current.s.pos, turn(current.s.dir, false)}, current.p, current.cost + 1000}
		visitedCost, ok = visited[turnRightPath.s]
		if !ok || turnRightPath.cost <= visitedCost {
			visited[turnRightPath.s] = turnRightPath.cost
			stack = append(stack, turnRightPath)
		}

	}

	positions := make(map[position]bool)
	for _, minPath := range minPaths {
		for _, pos := range minPath {
			positions[pos] = true
		}
	}

	return strconv.Itoa(len(positions))
}

type tile int

const (
	empty tile = iota
	wall
)

type direction int

const (
	east direction = iota
	south
	west
	north
)

type position struct {
	x, y int
}

type step struct {
	pos position
	dir direction
}

type path struct {
	s    step
	cost int
}

type path2 struct {
	s    step
	p    []position
	cost int
}

func parseInput(input string) (tiles [][]tile, start, end position) {
	tiles = make([][]tile, 0)
	lines := strings.Split(input, "\n")
	for y, line := range lines {
		trimmedLine := strings.TrimRight(line, "\r")
		row := make([]tile, len(trimmedLine))
		for x, r := range trimmedLine {
			switch r {
			case '#':
				row[x] = wall
			case 'S':
				start = position{x, y}
				row[x] = empty
			case 'E':
				end = position{x, y}
				row[x] = empty
			case '.':
				row[x] = empty
			}
		}
		tiles = append(tiles, row)
	}
	return
}

func turn(dir direction, left bool) direction {
	if left {
		return (dir + 3) % 4
	} else {
		return (dir + 1) % 4
	}
}

func move(pos position, dir direction) position {
	switch dir {
	case east:
		return position{pos.x + 1, pos.y}
	case south:
		return position{pos.x, pos.y + 1}
	case west:
		return position{pos.x - 1, pos.y}
	case north:
		return position{pos.x, pos.y - 1}
	}
	return pos
}
