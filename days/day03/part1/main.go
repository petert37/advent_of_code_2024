package main

import (
	"fmt"

	"github.com/petert37/advent_of_code_2024/common"
	"github.com/petert37/advent_of_code_2024/days/day03"
)

func main() {
	input := common.ReadInput("days/day03/input.txt")
	result := day03.ProcessPart1(input)
	fmt.Println(result)
}
