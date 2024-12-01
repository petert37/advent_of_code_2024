package day01

import "testing"

const input string = `3   4
4   3
2   5
1   3
3   9
3   3`

func TestPart1(t *testing.T) {
	result := ProcessPart1(input)
	expected := "11"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestPart2(t *testing.T) {
	result := ProcessPart2(input)
	expected := "31"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
