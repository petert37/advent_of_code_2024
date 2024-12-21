package day21

import "testing"

var input = `029A
980A
179A
456A
379A`

func TestPart1(t *testing.T) {
	result := ProcessPart1(input)
	expected := "126384"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestPart2(t *testing.T) {
	result := ProcessPart2(input)
	expected := "154115708116294"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
