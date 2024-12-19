package main

import (
	"fmt"

	"github.com/petert37/advent_of_code_2024/common"
	"github.com/petert37/advent_of_code_2024/days/day19"
)

func main() {
	input := common.ReadInput("days/day19/input.txt")
	result := day19.ProcessPart2(input)
	fmt.Println(result)
}
