package main

import (
	"fmt"

	"github.com/petert37/advent_of_code_2024/common"
	"github.com/petert37/advent_of_code_2024/days/day10"
)

func main() {
	input := common.ReadInput("days/day10/input.txt")
	result := day10.ProcessPart2(input)
	fmt.Println(result)
}
