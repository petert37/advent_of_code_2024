package main

import (
	"fmt"

	"github.com/petert37/advent_of_code_2024/common"
	"github.com/petert37/advent_of_code_2024/days/day09"
)

func main() {
	input := common.ReadInput("days/day09/input.txt")
	result := day09.ProcessPart1(input)
	fmt.Println(result)
}
