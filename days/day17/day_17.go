package day17

import (
	"math"
	"strconv"
	"strings"
)

func ProcessPart1(input string) string {
	registerA, registerB, registerC, program := parseInput(input)
	output := runProgram(registerA, registerB, registerC, program)
	outputStrs := make([]string, len(output))
	for i, outputInt := range output {
		outputStrs[i] = strconv.Itoa(outputInt)
	}
	return strings.Join(outputStrs, ",")
}

func ProcessPart2(input string) string {
	_, registerB, registerC, program := parseInput(input)
	result := backTrack([]int{}, registerB, registerC, &program, 0)
	return strconv.Itoa(result)
}

func parseInput(input string) (registerA, registerB, registerC int, program []int) {
	registerA = 0
	registerB = 0
	registerC = 0
	program = make([]int, 0)
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		trimmedLine := strings.TrimRight(line, "\r")
		if strings.HasPrefix(trimmedLine, "Register A: ") {
			value, err := strconv.Atoi(trimmedLine[12:])
			if err != nil {
				panic(err)
			}
			registerA = value
		} else if strings.HasPrefix(trimmedLine, "Register B: ") {
			value, err := strconv.Atoi(trimmedLine[12:])
			if err != nil {
				panic(err)
			}
			registerB = value
		} else if strings.HasPrefix(trimmedLine, "Register C: ") {
			value, err := strconv.Atoi(trimmedLine[12:])
			if err != nil {
				panic(err)
			}
			registerC = value
		} else if strings.HasPrefix(trimmedLine, "Program: ") {
			programStr := trimmedLine[9:]
			programStrs := strings.Split(programStr, ",")
			for _, programStr := range programStrs {
				programInt, err := strconv.Atoi(programStr)
				if err != nil {
					panic(err)
				}
				program = append(program, programInt)
			}
		}
	}
	return
}

func runProgram(registerA, registerB, registerC int, program []int) []int {
	output := make([]int, 0)
	programCounter := 0
	for programCounter < len(program) {
		instruction := program[programCounter]
		operand := program[programCounter+1]
		switch instruction {
		case 0:
			registerA = registerA / int(math.Pow(2, float64(getComboOperandValue(registerA, registerB, registerC, operand))))
			programCounter += 2
		case 1:
			registerB = registerB ^ operand
			programCounter += 2
		case 2:
			registerB = getComboOperandValue(registerA, registerB, registerC, operand) % 8
			programCounter += 2
		case 3:
			if registerA == 0 {
				programCounter += 2
			} else {
				programCounter = operand
			}
		case 4:
			registerB = registerB ^ registerC
			programCounter += 2
		case 5:
			output = append(output, getComboOperandValue(registerA, registerB, registerC, operand)%8)
			programCounter += 2
		case 6:
			registerB = registerA / int(math.Pow(2, float64(getComboOperandValue(registerA, registerB, registerC, operand))))
			programCounter += 2
		case 7:
			registerC = registerA / int(math.Pow(2, float64(getComboOperandValue(registerA, registerB, registerC, operand))))
			programCounter += 2
		}
	}
	return output
}

func getComboOperandValue(registerA, registerB, registerC int, operand int) int {
	if operand <= 3 {
		return operand
	} else if operand == 4 {
		return registerA
	} else if operand == 5 {
		return registerB
	} else if operand == 6 {
		return registerC
	}
	return 0
}

func backTrack(registerADigits []int, registerB, registerC int, program *[]int, depth int) int {
	registerA := assembleRegister(registerADigits)
	output := runProgram(registerA, registerB, registerC, *program)
	if len(output) == depth || depth == 0 {
		match := true
		for i := 0; i < depth; i++ {
			if output[i] != (*program)[len(*program)-depth+i] {
				match = false
				break
			}
		}
		if match {
			if depth == len(*program) {
				return registerA
			} else {
				for i := 0; i < 8; i++ {
					digits := make([]int, len(registerADigits))
					copy(digits, registerADigits)
					digits = append(digits, i)
					result := backTrack(digits, registerB, registerC, program, depth+1)
					if result != -1 {
						return result
					}
				}
			}

		}
	}
	return -1
}

func assembleRegister(digits []int) int {
	value := 0
	for _, digit := range digits {
		value = value*8 + digit
	}
	return value
}
