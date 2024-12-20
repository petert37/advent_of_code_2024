package day20

import "testing"

var input = `###############
#...#...#.....#
#.#.#.#.#.###.#
#S#...#.#.#...#
#######.#.#.###
#######.#.#...#
#######.#.###.#
###..E#...#...#
###.#######.###
#...###...#...#
#.#####.#.###.#
#.#...#.#.#...#
#.#.#.#.#.#.###
#...#...#...###
###############`

func TestPart1(t *testing.T) {
	result := ProcessPart1(input, 1)
	expected := "44"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestPart2(t *testing.T) {
	result := ProcessPart2(input, 50)
	expected := "285"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
