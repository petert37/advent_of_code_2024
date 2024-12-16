package day15

import (
	"strconv"
	"strings"
)

func ProcessPart1(input string) string {
	tiles, moves := parseInput(input)
	robotX, robotY := 0, 0
	for y, row := range tiles {
		for x, t := range row {
			if t == robot {
				robotX = x
				robotY = y
			}
		}
	}
	for _, m := range moves {
		if move(&tiles, m, robotX, robotY) {
			dx, dy := getDirection(m)
			robotX += dx
			robotY += dy
		}
	}
	sumGps := 0
	for y, row := range tiles {
		for x, t := range row {
			if t == box {
				sumGps += x + y*100
			}
		}
	}
	return strconv.Itoa(sumGps)
}

func ProcessPart2(input string) string {
	tiles, moves := parseInput2(input)
	robotX, robotY := 0, 0
	for y, row := range tiles {
		for x, t := range row {
			if t == robot {
				robotX = x
				robotY = y
			}
		}
	}
	for _, m := range moves {
		if canMove2(&tiles, m, robotX, robotY) {
			doMove2(&tiles, m, robotX, robotY)
			dx, dy := getDirection(m)
			robotX += dx
			robotY += dy
		}
	}
	sumGps := 0
	for y, row := range tiles {
		for x, t := range row {
			if t == leftBox {
				sumGps += x + y*100
			}
		}
	}
	return strconv.Itoa(sumGps)
}

type tile int

const (
	empty tile = iota
	robot
	box
	wall
	leftBox
	rightBox
)

type direction int

const (
	up direction = iota
	down
	left
	right
)

func parseInput(input string) (tiles [][]tile, moves []direction) {
	tiles = make([][]tile, 0)
	moves = make([]direction, 0)
	lines := strings.Split(input, "\n")

	isMap := true

	for _, line := range lines {
		trimmedLine := strings.TrimRight(line, "\r")
		if len(trimmedLine) == 0 {
			isMap = false
		}
		if isMap {
			lineTiles := make([]tile, len(trimmedLine))
			for i, t := range trimmedLine {
				switch t {
				case '#':
					lineTiles[i] = wall
				case 'O':
					lineTiles[i] = box
				case '@':
					lineTiles[i] = robot
				default:
					lineTiles[i] = empty
				}
			}
			tiles = append(tiles, lineTiles)
		} else {
			for _, m := range trimmedLine {
				switch m {
				case '^':
					moves = append(moves, up)
				case 'v':
					moves = append(moves, down)
				case '<':
					moves = append(moves, left)
				case '>':
					moves = append(moves, right)
				}
			}
		}
	}
	return
}

func parseInput2(input string) (tiles [][]tile, moves []direction) {
	tiles = make([][]tile, 0)
	moves = make([]direction, 0)
	lines := strings.Split(input, "\n")

	isMap := true

	for _, line := range lines {
		trimmedLine := strings.TrimRight(line, "\r")
		if len(trimmedLine) == 0 {
			isMap = false
		}
		if isMap {
			lineTiles := make([]tile, 0)
			for _, t := range trimmedLine {
				switch t {
				case '#':
					lineTiles = append(lineTiles, wall)
					lineTiles = append(lineTiles, wall)
				case 'O':
					lineTiles = append(lineTiles, leftBox)
					lineTiles = append(lineTiles, rightBox)
				case '@':
					lineTiles = append(lineTiles, robot)
					lineTiles = append(lineTiles, empty)
				default:
					lineTiles = append(lineTiles, empty)
					lineTiles = append(lineTiles, empty)
				}
			}
			tiles = append(tiles, lineTiles)
		} else {
			for _, m := range trimmedLine {
				switch m {
				case '^':
					moves = append(moves, up)
				case 'v':
					moves = append(moves, down)
				case '<':
					moves = append(moves, left)
				case '>':
					moves = append(moves, right)
				}
			}
		}
	}
	return
}

func getDirection(d direction) (int, int) {
	switch d {
	case up:
		return 0, -1
	case down:
		return 0, 1
	case left:
		return -1, 0
	case right:
		return 1, 0
	}
	return 0, 0
}

func move(tiles *[][]tile, d direction, x, y int) bool {
	currentTile := (*tiles)[y][x]
	dx, dy := getDirection(d)
	newX := x + dx
	newY := y + dy
	if newY < 0 || newY >= len(*tiles) {
		return false
	}
	row := (*tiles)[newY]
	if newX < 0 || newX >= len(row) {
		return false
	}
	newTile := row[newX]
	if newTile == wall {
		return false
	}
	if newTile == empty {
		row[newX] = currentTile
		(*tiles)[y][x] = empty
		return true
	}
	if newTile == box && move(tiles, d, newX, newY) {
		row[newX] = currentTile
		(*tiles)[y][x] = empty
		return true
	}
	return false
}

func canMove2(tiles *[][]tile, d direction, x, y int) bool {
	currentTile := (*tiles)[y][x]
	dx, dy := getDirection(d)
	newX := x + dx
	newY := y + dy
	if newY < 0 || newY >= len(*tiles) {
		return false
	}
	row := (*tiles)[newY]
	if newX < 0 || newX >= len(row) {
		return false
	}
	newTile := row[newX]
	if newTile == wall {
		return false
	}
	if newTile == empty && currentTile == robot {
		return true
	}
	if currentTile == leftBox || currentTile == rightBox {
		otherBoxHalfX := x
		otherBoxHalfY := y
		if currentTile == leftBox {
			otherBoxHalfX += 1
		} else {
			otherBoxHalfX -= 1
		}
		if d == left {
			if currentTile == rightBox {
				return canMove2(tiles, d, otherBoxHalfX, otherBoxHalfY)
			}
			if currentTile == leftBox {
				if newTile == empty {
					return true
				}
				if newTile == leftBox || newTile == rightBox {
					return canMove2(tiles, d, newX, newY)
				}
			}
		}
		if d == right {
			if currentTile == leftBox {
				return canMove2(tiles, d, otherBoxHalfX, otherBoxHalfY)
			}
			if currentTile == rightBox {
				if newTile == empty {
					return true
				}
				if newTile == leftBox || newTile == rightBox {
					return canMove2(tiles, d, newX, newY)
				}
			}
		}
		if d == up || d == down {
			otherBoxHalfNewX := otherBoxHalfX + dx
			otherBoxHalfNewY := otherBoxHalfY + dy
			if otherBoxHalfNewY < 0 || otherBoxHalfNewY >= len(*tiles) {
				return false
			}
			otherHalfRow := (*tiles)[otherBoxHalfNewY]
			if otherBoxHalfNewX < 0 || otherBoxHalfNewX >= len(otherHalfRow) {
				return false
			}
			otherHalfNewTile := otherHalfRow[otherBoxHalfNewX]
			if otherHalfNewTile == wall {
				return false
			}
			thisHalfCanMove := newTile == empty || canMove2(tiles, d, newX, newY)
			otherHalfCanMove := otherHalfNewTile == empty || canMove2(tiles, d, otherBoxHalfNewX, otherBoxHalfNewY)
			return thisHalfCanMove && otherHalfCanMove
		}
	}
	if newTile == leftBox || newTile == rightBox {
		return canMove2(tiles, d, newX, newY)
	}

	return false
}

func doMove2(tiles *[][]tile, d direction, x, y int) {
	currentTile := (*tiles)[y][x]
	dx, dy := getDirection(d)
	newX := x + dx
	newY := y + dy
	newTile := (*tiles)[newY][newX]

	if newTile != empty {
		doMove2(tiles, d, newX, newY)
	}
	(*tiles)[newY][newX] = currentTile
	(*tiles)[y][x] = empty

	if (currentTile == leftBox || currentTile == rightBox) && (d == up || d == down) {
		otherBoxHalfX := x
		otherBoxHalfY := y
		if currentTile == leftBox {
			otherBoxHalfX += 1
		} else {
			otherBoxHalfX -= 1
		}
		otherBoxHalfNewX := otherBoxHalfX + dx
		otherBoxHalfNewY := otherBoxHalfY + dy
		otherBoxHalfCurrentTile := (*tiles)[otherBoxHalfY][otherBoxHalfX]
		otherBoxHalNewTile := (*tiles)[otherBoxHalfNewY][otherBoxHalfNewX]

		if otherBoxHalNewTile != empty {
			doMove2(tiles, d, otherBoxHalfNewX, otherBoxHalfNewY)
		}
		(*tiles)[otherBoxHalfNewY][otherBoxHalfNewX] = otherBoxHalfCurrentTile
		(*tiles)[otherBoxHalfY][otherBoxHalfX] = empty
	}
}
