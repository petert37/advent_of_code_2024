package main

import (
	"fmt"

	"github.com/petert37/advent_of_code_2024/common"
	"github.com/petert37/advent_of_code_2024/days/day22"
)

func main() {
	input := common.ReadInput("days/day22/input.txt")
	result := day22.ProcessPart2(input)
	fmt.Println(result)
}
