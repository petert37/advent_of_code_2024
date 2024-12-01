package main

import (
	"fmt"

	"github.com/petert37/advent_of_code_2024/common"
	"github.com/petert37/advent_of_code_2024/days/day01"
)

func main() {
	input := common.ReadInput("days/day01/input.txt")
	result := day01.ProcessPart2(input)
	fmt.Println(result)
}
