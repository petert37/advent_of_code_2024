package main

import (
	"fmt"

	"github.com/petert37/advent_of_code_2024/common"
	"github.com/petert37/advent_of_code_2024/days/day06"
)

func main() {
	input := common.ReadInput("days/day06/input.txt")
	result := day06.ProcessPart1(input)
	fmt.Println(result)
}
