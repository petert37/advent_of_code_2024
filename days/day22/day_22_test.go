package day22

import "testing"

func TestPart1(t *testing.T) {
	input := `1
10
100
2024`
	result := ProcessPart1(input)
	expected := "37327623"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestPart2(t *testing.T) {
	input := `1
2
3
2024`
	result := ProcessPart2(input)
	expected := "23"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
