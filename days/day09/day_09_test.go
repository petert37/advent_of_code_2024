package day09

import "testing"

const input string = `2333133121414131402`

func TestPart1(t *testing.T) {
	result := ProcessPart1(input)
	expected := "1928"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestPart22(t *testing.T) {
	result := ProcessPart2(input)
	expected := "2858"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
