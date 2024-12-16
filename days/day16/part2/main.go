package main

import (
	"fmt"

	"github.com/petert37/advent_of_code_2024/common"
	"github.com/petert37/advent_of_code_2024/days/day16"
)

func main() {
	input := common.ReadInput("days/day16/input.txt")
	result := day16.ProcessPart2(input)
	fmt.Println(result)
}
