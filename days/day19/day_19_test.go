package day19

import "testing"

var input = `r, wr, b, g, bwu, rb, gb, br

brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb`

func TestPart1(t *testing.T) {
	result := ProcessPart1(input)
	expected := "6"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestPart2(t *testing.T) {
	result := ProcessPart2(input)
	expected := "16"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
