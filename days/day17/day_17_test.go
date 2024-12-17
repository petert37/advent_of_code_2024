package day17

import "testing"

func TestPart1(t *testing.T) {
	input := `Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0`

	result := ProcessPart1(input)
	expected := "4,6,3,5,6,3,5,2,1,0"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestPart2(t *testing.T) {
	input := `Register A: 2024
Register B: 0
Register C: 0

Program: 0,3,5,4,3,0`

	result := ProcessPart2(input)
	expected := "117440"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
