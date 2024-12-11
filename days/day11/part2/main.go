package main

import (
	"fmt"

	"github.com/petert37/advent_of_code_2024/common"
	"github.com/petert37/advent_of_code_2024/days/day11"
)

func main() {
	input := common.ReadInput("days/day11/input.txt")
	result := day11.ProcessPart2(input)
	fmt.Println(result)
}
