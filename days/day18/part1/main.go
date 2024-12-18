package main

import (
	"fmt"

	"github.com/petert37/advent_of_code_2024/common"
	"github.com/petert37/advent_of_code_2024/days/day18"
)

func main() {
	input := common.ReadInput("days/day18/input.txt")
	result := day18.ProcessPart1(input, 70, 1024)
	fmt.Println(result)
}
