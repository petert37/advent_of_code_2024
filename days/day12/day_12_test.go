package day12

import "testing"

func TestPart11(t *testing.T) {
	input := `AAAA
BBCD
BBCC
EEEC`
	result := ProcessPart1(input)
	expected := "140"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestPart12(t *testing.T) {
	input := `OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`
	result := ProcessPart1(input)
	expected := "772"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestPart13(t *testing.T) {
	input := `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`
	result := ProcessPart1(input)
	expected := "1930"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestPart21(t *testing.T) {
	input := `AAAA
BBCD
BBCC
EEEC`
	result := ProcessPart2(input)
	expected := "80"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestPart22(t *testing.T) {
	input := `OOOOO
OXOXO
OOOOO
OXOXO
OOOOO`
	result := ProcessPart2(input)
	expected := "436"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestPart23(t *testing.T) {
	input := `EEEEE
EXXXX
EEEEE
EXXXX
EEEEE`
	result := ProcessPart2(input)
	expected := "236"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestPart24(t *testing.T) {
	input := `AAAAAA
AAABBA
AAABBA
ABBAAA
ABBAAA
AAAAAA`
	result := ProcessPart2(input)
	expected := "368"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}

func TestPart25(t *testing.T) {
	input := `RRRRIICCFF
RRRRIICCCF
VVRRRCCFFF
VVRCCCJFFF
VVVVCJJCFE
VVIVCCJJEE
VVIIICJJEE
MIIIIIJJEE
MIIISIJEEE
MMMISSJEEE`
	result := ProcessPart2(input)
	expected := "1206"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
