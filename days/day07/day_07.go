package day07

import (
	"iter"
	"math"
	"strconv"
	"strings"
)

func ProcessPart1(input string) string {
	operators := []operator{Add, Multiply}
	return process(input, operators)
}

func ProcessPart2(input string) string {
	operators := []operator{Add, Multiply, Concatenate}
	return process(input, operators)
}

func process(input string, operators []operator) string {
	equations := parseInput(input)
	result := make(chan int64)
	var sum int64 = 0
	for _, equation := range equations {
		go getEquationValue(equation, operators, result)
	}
	for range equations {
		sum += <-result
	}
	return strconv.FormatInt(sum, 10)
}

type equation struct {
	result   int
	operands []int
}

type operator int

const (
	Add operator = iota
	Multiply
	Concatenate
)

func parseInput(input string) (equations []equation) {
	equations = make([]equation, 0)

	lines := strings.Split(input, "\n")

	for _, line := range lines {
		trimmedLine := strings.TrimRight(line, "\r")
		parts := strings.Split(trimmedLine, ": ")
		result, _ := strconv.Atoi(parts[0])
		operands := make([]int, 0)
		for _, operand := range strings.Split(parts[1], " ") {
			operandInt, _ := strconv.Atoi(operand)
			operands = append(operands, operandInt)
		}
		equations = append(equations, equation{result, operands})
	}

	return
}

func possibleOperators(options []operator, length int) iter.Seq[[]operator] {
	return func(yield func([]operator) bool) {
		lenOptions := len(options)
		lenOptionsf64 := float64(lenOptions)
		numOptions := int(math.Pow(lenOptionsf64, float64(length)))
		for i := 0; i < numOptions; i++ {
			sequence := make([]operator, length)
			for j := 0; j < length; j++ {
				digit := int(math.Pow(lenOptionsf64, float64(j)))
				sequence[j] = options[(i/digit)%lenOptions]
			}
			if !yield(sequence) {
				break
			}
		}
	}
}

func getEquationValue(equation equation, operators []operator, result chan<- int64) {
	for op := range possibleOperators(operators, len(equation.operands)-1) {
		if evaluate(equation.operands, op) == equation.result {
			result <- int64(equation.result)
			return
		}
	}
	result <- 0
}

func evaluate(operands []int, operators []operator) int {
	result := operands[0]
	for i := 1; i < len(operands); i++ {
		switch operators[i-1] {
		case Add:
			result += operands[i]
		case Multiply:
			result *= operands[i]
		case Concatenate:
			numDigits := len(strconv.Itoa(operands[i]))
			shift := int(math.Pow10(numDigits))
			result = result*shift + operands[i]
		}

	}
	return result
}
