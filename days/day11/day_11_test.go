package day11

import "testing"

const input string = `125 17`

func TestPart1(t *testing.T) {
	result := ProcessPart1(input)
	expected := "55312"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestPart2(t *testing.T) {
	result := ProcessPart2(input)
	expected := "65601038650482"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
