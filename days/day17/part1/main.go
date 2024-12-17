package main

import (
	"fmt"

	"github.com/petert37/advent_of_code_2024/common"
	"github.com/petert37/advent_of_code_2024/days/day17"
)

func main() {
	input := common.ReadInput("days/day17/input.txt")
	result := day17.ProcessPart1(input)
	fmt.Println(result)
}
