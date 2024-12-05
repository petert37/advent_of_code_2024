package main

import (
	"fmt"

	"github.com/petert37/advent_of_code_2024/common"
	"github.com/petert37/advent_of_code_2024/days/day05"
)

func main() {
	input := common.ReadInput("days/day05/input.txt")
	result := day05.ProcessPart1(input)
	fmt.Println(result)
}
