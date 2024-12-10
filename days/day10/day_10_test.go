package day10

import "testing"

const input string = `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`

func TestPart1(t *testing.T) {
	result := ProcessPart1(input)
	expected := "36"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestPart22(t *testing.T) {
	result := ProcessPart2(input)
	expected := "81"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
