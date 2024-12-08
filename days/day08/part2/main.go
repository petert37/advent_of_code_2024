package main

import (
	"fmt"

	"github.com/petert37/advent_of_code_2024/common"
	"github.com/petert37/advent_of_code_2024/days/day08"
)

func main() {
	input := common.ReadInput("days/day08/input.txt")
	result := day08.ProcessPart2(input)
	fmt.Println(result)
}
