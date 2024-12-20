package main

import (
	"fmt"

	"github.com/petert37/advent_of_code_2024/common"
	"github.com/petert37/advent_of_code_2024/days/day20"
)

func main() {
	input := common.ReadInput("days/day20/input.txt")
	result := day20.ProcessPart2(input, 100)
	fmt.Println(result)
}
