package day12

import (
	"sort"
	"strconv"
	"strings"
)

func ProcessPart1(input string) string {
	garden := parseInput(input)
	regions := make([]map[position]bool, 0)
	for y, row := range garden {
		for x := range row {
			regionExists := false
			for _, region := range regions {
				if _, ok := region[position{x, y}]; ok {
					regionExists = true
					break
				}
			}
			if !regionExists {
				region := makeRegion(&garden, x, y)
				regions = append(regions, region)
			}
		}
	}
	price := 0
	for _, region := range regions {
		perimeter := getRegionPerimeter(&region)
		area := len(region)
		price += area * perimeter
	}
	return strconv.Itoa(price)
}

func ProcessPart2(input string) string {
	garden := parseInput(input)
	regions := make([]map[position]bool, 0)
	for y, row := range garden {
		for x := range row {
			regionExists := false
			for _, region := range regions {
				if _, ok := region[position{x, y}]; ok {
					regionExists = true
					break
				}
			}
			if !regionExists {
				region := makeRegion(&garden, x, y)
				regions = append(regions, region)
			}
		}
	}

	price := 0

	for _, region := range regions {
		lineParts := getRegionLineParts(&region)

		sort.Slice(lineParts, func(i, j int) bool {
			if lineParts[i].x == lineParts[j].x {
				if lineParts[i].y == lineParts[j].y {
					return lineParts[i].side < lineParts[j].side
				}
				return lineParts[i].y < lineParts[j].y
			}
			return lineParts[i].x < lineParts[j].x
		})

		horizontalFences := make(map[int][][]linePart)
		verticalFences := make(map[int][][]linePart)
		for _, lp := range lineParts {
			if lp.side == top || lp.side == bottom {
				if _, ok := horizontalFences[lp.y]; !ok {
					horizontalFences[lp.y] = make([][]linePart, 0)
				}
				found := false
				for i, fence := range horizontalFences[lp.y] {
					last := fence[len(fence)-1]
					if last.x == lp.x-1 && last.side == lp.side {
						horizontalFences[lp.y][i] = append(horizontalFences[lp.y][i], lp)
						found = true
						break
					}
				}
				if !found {
					horizontalFences[lp.y] = append(horizontalFences[lp.y], []linePart{lp})
				}
			} else {
				if _, ok := verticalFences[lp.x]; !ok {
					verticalFences[lp.x] = make([][]linePart, 0)
				}
				found := false
				for i, fence := range verticalFences[lp.x] {
					last := fence[len(fence)-1]
					if last.y == lp.y-1 && last.side == lp.side {
						verticalFences[lp.x][i] = append(verticalFences[lp.x][i], lp)
						found = true
						break
					}
				}
				if !found {
					verticalFences[lp.x] = append(verticalFences[lp.x], []linePart{lp})
				}
			}
		}

		area := len(region)
		numSides := 0
		for _, fence := range horizontalFences {
			numSides += len(fence)
		}
		for _, fence := range verticalFences {
			numSides += len(fence)
		}

		price += area * numSides
	}
	return strconv.Itoa(price)
}

type position struct {
	x, y int
}

func parseInput(input string) [][]rune {
	result := make([][]rune, 0)
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		trimmedLine := strings.TrimRight(line, "\r")
		runes := make([]rune, len(trimmedLine))
		for i, r := range trimmedLine {
			runes[i] = r
		}
		result = append(result, runes)
	}
	return result
}

var directions = [][]int{{0, 1}, {0, -1}, {1, 0}, {-1, 0}}

type side int

const (
	bottom side = iota
	top
	right
	left
)

func makeRegion(garden *[][]rune, x, y int) map[position]bool {
	region := make(map[position]bool)
	region[position{x, y}] = true
	stack := make([]position, 0)
	stack = append(stack, position{x, y})
	startCell := (*garden)[y][x]
	for {
		if len(stack) == 0 {
			break
		}
		current := stack[0]
		stack = stack[1:]
		for direction := range directions {
			newX := current.x + directions[direction][0]
			newY := current.y + directions[direction][1]
			if newY >= 0 && newY < len(*garden) {
				row := (*garden)[newY]
				if newX >= 0 && newX < len(row) {
					cell := row[newX]
					if cell == startCell {
						newPosition := position{newX, newY}
						if _, ok := region[newPosition]; ok {
							continue
						}
						stack = append(stack, newPosition)
						region[newPosition] = true
					}
				}
			}
		}
	}
	return region
}

func getRegionPerimeter(region *map[position]bool) int {
	area := 0
	for p := range *region {
		area += 4
		for direction := range directions {
			newX := p.x + directions[direction][0]
			newY := p.y + directions[direction][1]
			if _, ok := (*region)[position{newX, newY}]; ok {
				area--
			}
		}
	}
	return area
}

type linePart struct {
	x, y int
	side side
}

func getRegionLineParts(region *map[position]bool) []linePart {
	lineParts := make([]linePart, 0)
	for p := range *region {
		for direction := range directions {
			newX := p.x + directions[direction][0]
			newY := p.y + directions[direction][1]
			if _, ok := (*region)[position{newX, newY}]; !ok {
				lineParts = append(lineParts, linePart{p.x, p.y, side(direction)})
			}
		}
	}
	return lineParts
}
