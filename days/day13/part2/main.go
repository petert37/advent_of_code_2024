package main

import (
	"fmt"

	"github.com/petert37/advent_of_code_2024/common"
	"github.com/petert37/advent_of_code_2024/days/day13"
)

func main() {
	input := common.ReadInput("days/day13/input.txt")
	result := day13.ProcessPart2(input)
	fmt.Println(result)
}
