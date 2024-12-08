package day08

import "testing"

const input string = `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`

func TestPart1(t *testing.T) {
	result := ProcessPart1(input)
	expected := "14"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestPart22(t *testing.T) {
	result := ProcessPart2(input)
	expected := "34"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
