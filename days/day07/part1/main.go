package main

import (
	"fmt"

	"github.com/petert37/advent_of_code_2024/common"
	"github.com/petert37/advent_of_code_2024/days/day07"
)

func main() {
	input := common.ReadInput("days/day07/input.txt")
	result := day07.ProcessPart1(input)
	fmt.Println(result)
}
