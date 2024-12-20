package day20

import (
	"iter"
	"math"
	"strconv"
	"strings"
)

func ProcessPart1(input string, minCheatDistance int) string {
	count := process(input, minCheatDistance, 2)
	return strconv.Itoa(count)
}

func ProcessPart2(input string, minCheatDistance int) string {
	count := process(input, minCheatDistance, 20)
	return strconv.Itoa(count)
}

func process(input string, minCheatSavings int, maxCheatDistance int) int {
	tiles, start, end := parseInput(input)
	distance := 0
	distances := make(map[position]int)
	track := make([]position, 0)
	track = append(track, start)
	currentPosition := start
	for currentPosition != end {
		distances[currentPosition] = distance
		distance++
		for _, direction := range directions {
			nextPosition := position{currentPosition.x + direction.x, currentPosition.y + direction.y}
			if nextPosition.y >= 0 && nextPosition.y < len(tiles) {
				row := tiles[nextPosition.y]
				if nextPosition.x >= 0 && nextPosition.x < len(row) {
					if row[nextPosition.x] == path {
						if _, ok := distances[nextPosition]; !ok {
							currentPosition = nextPosition
							track = append(track, currentPosition)
							break
						}
					}
				}
			}
		}
	}
	distances[end] = distance
	count := 0
	for _, currentPosition := range track {
		for direction := range cheatMoves(maxCheatDistance) {
			nextPosition := position{currentPosition.x + direction.x, currentPosition.y + direction.y}
			cheatDistance := int(math.Abs(float64(direction.x)) + math.Abs(float64(direction.y)))
			if nextPosition.y >= 0 && nextPosition.y < len(tiles) {
				row := tiles[nextPosition.y]
				if nextPosition.x >= 0 && nextPosition.x < len(row) {
					if row[nextPosition.x] == path {
						currentDistance := distances[currentPosition]
						nextDistance := distances[nextPosition]
						if nextDistance-currentDistance-cheatDistance >= minCheatSavings {
							count++
						}
					}
				}
			}
		}
	}
	return count
}

type tile int

const (
	path tile = iota
	wall
)

type position struct {
	x, y int
}

var directions = []position{{0, -1}, {1, 0}, {0, 1}, {-1, 0}}

func parseInput(input string) (tiles [][]tile, start, end position) {
	tiles = make([][]tile, 0)
	lines := strings.Split(input, "\n")
	for y, line := range lines {
		trimmedLine := strings.TrimRight(line, "\r")
		row := make([]tile, len(trimmedLine))
		for x, char := range trimmedLine {
			switch char {
			case '#':
				row[x] = wall
			case '.':
				row[x] = path
			case 'S':
				row[x] = path
				start = position{x, y}
			case 'E':
				row[x] = path
				end = position{x, y}
			}
		}
		tiles = append(tiles, row)
	}
	return
}

func cheatMoves(distance int) iter.Seq[position] {
	return func(yield func(position) bool) {
		for x := 0; x <= distance; x++ {
			for y := 0; y <= distance-x; y++ {
				if !yield(position{x, y}) {
					return
				}
				if x != 0 {
					if !yield(position{-x, y}) {
						return
					}
				}
				if y != 0 {
					if !yield(position{x, -y}) {
						return
					}
				}
				if x != 0 && y != 0 {
					if !yield(position{-x, -y}) {
						return
					}
				}
			}
		}
	}
}
