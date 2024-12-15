package main

import (
	"fmt"

	"github.com/draffensperger/golp"
	"github.com/petert37/advent_of_code_2024/common"
	"github.com/petert37/advent_of_code_2024/days/day13"
)

func main() {
	input := common.ReadInput("days/day13/input.txt")
	result := day13.ProcessPart2(input)
	fmt.Println(result)
}

func main2() {
	lp := golp.NewLP(0, 2)
	lp.AddConstraint([]float64{94.0, 22.0}, golp.EQ, 8400.0)
	lp.AddConstraint([]float64{34.0, 67.0}, golp.EQ, 5400.0)
	lp.SetObjFn([]float64{3.0, 1.0})
	lp.SetInt(0, true)
	lp.SetInt(1, true)
	// lp.SetMaximize()

	solution := lp.Solve()
	vars := lp.Variables()
	fmt.Printf("Solution: %v\n", solution)
	fmt.Printf("Button A: %.3f\n", vars[0])
	fmt.Printf("Button B: %.3f\n", vars[1])
	fmt.Printf("Tokens: %.3f\n", lp.Objective())

	// No need to explicitly free underlying C structure as golp.LP finalizer will
}

func main3() {
	lp := golp.NewLP(0, 2)
	lp.AddConstraint([]float64{26.0, 67.0}, golp.EQ, 12748.0)
	lp.AddConstraint([]float64{66.0, 21.0}, golp.EQ, 12176.0)
	lp.SetObjFn([]float64{3.0, 1.0})
	lp.SetInt(0, true)
	lp.SetInt(1, true)
	// lp.SetMaximize()

	solution := lp.Solve()
	vars := lp.Variables()
	fmt.Printf("Solution: %v\n", solution)
	fmt.Printf("Plant %.3f acres of barley\n", vars[0])
	fmt.Printf("And  %.3f acres of wheat\n", vars[1])
	fmt.Printf("For optimal profit of $%.2f\n", lp.Objective())

	// No need to explicitly free underlying C structure as golp.LP finalizer will
}
