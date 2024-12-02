package main

import (
	"fmt"

	"github.com/petert37/advent_of_code_2024/common"
	"github.com/petert37/advent_of_code_2024/days/day02"
)

func main() {
	input := common.ReadInput("days/day02/input.txt")
	result := day02.ProcessPart2(input)
	fmt.Println(result)
}
