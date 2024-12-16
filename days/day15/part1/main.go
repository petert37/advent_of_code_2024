package main

import (
	"fmt"

	"github.com/petert37/advent_of_code_2024/common"
	"github.com/petert37/advent_of_code_2024/days/day15"
)

func main() {
	input := common.ReadInput("days/day15/input.txt")
	result := day15.ProcessPart1(input)
	fmt.Println(result)
}
