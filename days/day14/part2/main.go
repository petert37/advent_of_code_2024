package main

import (
	"fmt"

	"github.com/petert37/advent_of_code_2024/common"
	"github.com/petert37/advent_of_code_2024/days/day14"
)

func main() {
	input := common.ReadInput("days/day14/input.txt")
	result := day14.ProcessPart2(input, 101, 103)
	fmt.Println(result)
}
