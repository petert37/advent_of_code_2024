package day08

import (
	"strconv"
	"strings"

	"k8s.io/apimachinery/pkg/util/sets"
)

func ProcessPart1(input string) string {
	antennas, width, height := parseInput(input)
	antinodes := sets.New[position]()

	for _, positions := range antennas {
		for i, first_position := range positions {
			for j := i + 1; j < len(positions); j++ {
				second_position := positions[j]
				antinode_1_x := first_position.x + (first_position.x - second_position.x)
				antinode_1_y := first_position.y + (first_position.y - second_position.y)
				antinode_2_x := second_position.x + (second_position.x - first_position.x)
				antinode_2_y := second_position.y + (second_position.y - first_position.y)
				if antinode_1_x >= 0 && antinode_1_x < width && antinode_1_y >= 0 && antinode_1_y < height {
					antinodes.Insert(position{antinode_1_x, antinode_1_y})
				}
				if antinode_2_x >= 0 && antinode_2_x < width && antinode_2_y >= 0 && antinode_2_y < height {
					antinodes.Insert(position{antinode_2_x, antinode_2_y})
				}
			}
		}
	}

	return strconv.Itoa(antinodes.Len())
}

func ProcessPart2(input string) string {
	antennas, width, height := parseInput(input)
	antinodes := sets.New[position]()

	for _, positions := range antennas {
		for i, first_position := range positions {
			for j := i + 1; j < len(positions); j++ {
				second_position := positions[j]
				diff_x := first_position.x - second_position.x
				diff_y := first_position.y - second_position.y
				for n := 0; ; n++ {
					antinode_x := first_position.x + diff_x*n
					antinode_y := first_position.y + diff_y*n
					if antinode_x < 0 || antinode_x >= width || antinode_y < 0 || antinode_y >= height {
						break
					}
					antinodes.Insert(position{antinode_x, antinode_y})
				}
				for n := 0; ; n-- {
					antinode_x := first_position.x + diff_x*n
					antinode_y := first_position.y + diff_y*n
					if antinode_x < 0 || antinode_x >= width || antinode_y < 0 || antinode_y >= height {
						break
					}
					antinodes.Insert(position{antinode_x, antinode_y})
				}
			}
		}
	}

	return strconv.Itoa(antinodes.Len())
}

type position struct {
	x, y int
}

func parseInput(input string) (antennas map[rune][]position, width, height int) {
	antennas = make(map[rune][]position)
	width = 0
	height = 0

	lines := strings.Split(input, "\n")

	for y, line := range lines {
		if y+1 > height {
			height = y + 1
		}
		trimmedLine := strings.TrimRight(line, "\r")
		for x, char := range trimmedLine {
			if x+1 > width {
				width = x + 1
			}
			if char != '.' {
				current, exists := antennas[char]
				if exists {
					antennas[char] = append(current, position{x, y})
				} else {
					antennas[char] = []position{{x, y}}
				}
			}
		}
	}

	return
}
