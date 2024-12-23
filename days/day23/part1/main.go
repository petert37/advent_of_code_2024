package main

import (
	"fmt"

	"github.com/petert37/advent_of_code_2024/common"
	"github.com/petert37/advent_of_code_2024/days/day23"
)

func main() {
	input := common.ReadInput("days/day23/input.txt")
	result := day23.ProcessPart1(input)
	fmt.Println(result)
}
