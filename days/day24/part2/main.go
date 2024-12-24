package main

import (
	"fmt"

	"github.com/petert37/advent_of_code_2024/common"
	"github.com/petert37/advent_of_code_2024/days/day24"
)

func main() {
	input := common.ReadInput("days/day24/input.txt")
	result := day24.ProcessPart2(input)
	fmt.Println(result)
}
