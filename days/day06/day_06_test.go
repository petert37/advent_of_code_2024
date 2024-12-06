package day06

import "testing"

const input string = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

func TestPart1(t *testing.T) {
	result := ProcessPart1(input)
	expected := "41"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestPart22(t *testing.T) {
	result := ProcessPart2(input)
	expected := "6"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
