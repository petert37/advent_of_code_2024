package main

import (
	"fmt"

	"github.com/petert37/advent_of_code_2024/common"
	"github.com/petert37/advent_of_code_2024/days/day04"
)

func main() {
	input := common.ReadInput("days/day04/input.txt")
	result := day04.ProcessPart1(input)
	fmt.Println(result)
}
