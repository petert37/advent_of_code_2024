package day06

import (
	"strconv"
	"strings"

	"k8s.io/apimachinery/pkg/util/sets"
)

func ProcessPart1(input string) string {
	tiles, currentPosition := parseInput(input)

	visited := sets.New(currentPosition)
	direction := North

	for {
		nextPostion := getNextPosition(currentPosition, direction)
		if nextPostion.x < 0 || nextPostion.y < 0 || nextPostion.y >= len(tiles) || nextPostion.x >= len(tiles[nextPostion.y]) {
			break
		}
		if tiles[nextPostion.y][nextPostion.x] == Obstruction {
			direction = direction.next()
			continue
		}
		currentPosition = nextPostion
		visited.Insert(currentPosition)
	}

	return strconv.Itoa(visited.Len())
}

func ProcessPart2(input string) string {
	tiles, startPosition := parseInput(input)

	result := make(chan bool)
	check := 0
	for y := range tiles {
		for x := range tiles[y] {
			if tiles[y][x] == Empty && !(x == startPosition.x && y == startPosition.y) {
				go isLoop(&tiles, startPosition, position{x, y}, result)
				check++
			}
		}
	}

	count := 0
	for range check {
		if <-result {
			count++
		}
	}

	return strconv.Itoa(count)
}

type tile int

const (
	Empty tile = iota
	Obstruction
)

type direction int

const (
	North direction = iota
	East
	South
	West
)

type position struct {
	x, y int
}

func (d direction) next() direction {
	return (d + 1) % 4
}

type positionWithDirection struct {
	x, y int
	d    direction
}

func parseInput(input string) (tiles [][]tile, startPosition position) {
	startPosition = position{0, 0}
	tiles = make([][]tile, 0)

	lines := strings.Split(input, "\n")

	for y, line := range lines {
		trimmedLine := strings.TrimRight(line, "\r")
		row := make([]tile, len(trimmedLine))
		for x, tile := range trimmedLine {
			parsedTile := Empty
			if tile == '#' {
				parsedTile = Obstruction
			}
			if tile == '^' {
				startPosition = position{x, y}
			}
			row[x] = parsedTile
		}
		tiles = append(tiles, row)
	}

	return
}

func getNextPosition(current position, direction direction) position {
	switch direction {
	case North:
		return position{current.x, current.y - 1}
	case East:
		return position{current.x + 1, current.y}
	case South:
		return position{current.x, current.y + 1}
	case West:
		return position{current.x - 1, current.y}
	}
	return position{}
}

func isLoop(tiles *[][]tile, startPosition position, extraObstruction position, result chan<- bool) {
	turns := sets.New[positionWithDirection]()
	currentPosition := startPosition
	direction := North
	for {
		nextPostion := getNextPosition(currentPosition, direction)
		if nextPostion.x < 0 || nextPostion.y < 0 || nextPostion.y >= len(*tiles) || nextPostion.x >= len((*tiles)[nextPostion.y]) {
			result <- false
			return
		}
		if nextPostion.x == extraObstruction.x && nextPostion.y == extraObstruction.y || (*tiles)[nextPostion.y][nextPostion.x] == Obstruction {
			turn := positionWithDirection{currentPosition.x, currentPosition.y, direction}
			if turns.Has(turn) {
				result <- true
				return
			}
			turns.Insert(turn)
			direction = direction.next()
			continue
		}
		currentPosition = nextPostion
	}
}
