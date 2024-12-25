package day25

import "testing"

var input = `#####
.####
.####
.####
.#.#.
.#...
.....

#####
##.##
.#.##
...##
...#.
...#.
.....

.....
#....
#....
#...#
#.#.#
#.###
#####

.....
.....
#.#..
###..
###.#
###.#
#####

.....
.....
.....
#....
#.#..
#.#.#
#####`

func TestPart11(t *testing.T) {
	result := ProcessPart1(input)
	expected := "3"

	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}
