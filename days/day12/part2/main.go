package main

import (
	"fmt"

	"github.com/petert37/advent_of_code_2024/common"
	"github.com/petert37/advent_of_code_2024/days/day12"
)

func main() {
	input := common.ReadInput("days/day12/input.txt")
	result := day12.ProcessPart2(input)
	fmt.Println(result)
}
