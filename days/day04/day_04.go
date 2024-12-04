package day04

import (
	"strconv"
	"strings"
)

func ProcessPart1(input string) string {
	lines := strings.Split(input, "\n")

	word_search := make([][]rune, 0)

	for _, line := range lines {
		line_runes := make([]rune, 0)
		for _, char := range strings.TrimRight(line, "\r") {
			line_runes = append(line_runes, char)
		}
		word_search = append(word_search, line_runes)
	}

	count := 0

	for y := range len(word_search) {
		for x := range len(word_search[y]) {
			count += countXmas(&word_search, x, y)
		}
	}

	return strconv.Itoa(count)
}

var xmas [4]rune = [4]rune{'X', 'M', 'A', 'S'}
var directions [8][2]int = [8][2]int{{-1, -1}, {0, -1}, {1, -1}, {-1, 0}, {1, 0}, {-1, 1}, {0, 1}, {1, 1}}

func countXmas(word_search *[][]rune, x int, y int) int {
	direction_states := make([]bool, 8)
	for i := range 8 {
		direction_states[i] = true
	}
	for distance := range 4 {
		for direction := range 8 {
			dx := directions[direction][0] * distance
			dy := directions[direction][1] * distance
			next_x := x + dx
			next_y := y + dy
			if next_x >= 0 && next_x < len((*word_search)[x]) && next_y >= 0 && next_y < len(*word_search) {
				if (*word_search)[next_y][next_x] != xmas[distance] {
					direction_states[direction] = false
				}
			} else {
				direction_states[direction] = false
			}
		}
	}
	count := 0
	for direction := range 8 {
		if direction_states[direction] {
			count++
		}
	}
	return count
}

func ProcessPart2(input string) string {
	lines := strings.Split(input, "\n")

	word_search := make([][]rune, 0)

	for _, line := range lines {
		line_runes := make([]rune, 0)
		for _, char := range strings.TrimRight(line, "\r") {
			line_runes = append(line_runes, char)
		}
		word_search = append(word_search, line_runes)
	}

	count := 0

	for y := range len(word_search) {
		for x := range len(word_search[y]) {
			if isCross(&word_search, x, y) {
				count++
			}
		}
	}

	return strconv.Itoa(count)
}

var leftCrossOffsets [3][2]int = [3][2]int{{-1, -1}, {0, 0}, {1, 1}}
var rightCrossOffsets [3][2]int = [3][2]int{{-1, 1}, {0, 0}, {1, -1}}

func isCross(word_search *[][]rune, x int, y int) bool {
	left := getHalfCross(word_search, x, y, &leftCrossOffsets)
	right := getHalfCross(word_search, x, y, &rightCrossOffsets)
	return (left == "MAS" || left == "SAM") && (right == "MAS" || right == "SAM")
}

func getHalfCross(word_search *[][]rune, x int, y int, offsets *[3][2]int) string {
	half_cross := ""
	for i := range 3 {
		next_x := x + (*offsets)[i][0]
		next_y := y + (*offsets)[i][1]
		if next_x >= 0 && next_x < len((*word_search)[x]) && next_y >= 0 && next_y < len(*word_search) {
			half_cross += string((*word_search)[next_y][next_x])
		}
	}
	return half_cross
}
