package day03

import "testing"

func TestPart1(t *testing.T) {
	input := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"
	result := ProcessPart1(input)
	expected := "161"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestPart2(t *testing.T) {
	input := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"
	result := ProcessPart2(input)
	expected := "48"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
