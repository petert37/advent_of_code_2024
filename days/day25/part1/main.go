package main

import (
	"fmt"

	"github.com/petert37/advent_of_code_2024/common"
	"github.com/petert37/advent_of_code_2024/days/day25"
)

func main() {
	input := common.ReadInput("days/day25/input.txt")
	result := day25.ProcessPart1(input)
	fmt.Println(result)
}
