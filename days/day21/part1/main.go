package main

import (
	"fmt"

	"github.com/petert37/advent_of_code_2024/common"
	"github.com/petert37/advent_of_code_2024/days/day21"
)

func main() {
	input := common.ReadInput("days/day21/input.txt")
	result := day21.ProcessPart1(input)
	fmt.Println(result)
}
