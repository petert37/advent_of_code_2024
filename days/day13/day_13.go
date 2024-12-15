package day13

import (
	"strconv"
	"strings"

	"github.com/draffensperger/golp"
)

func ProcessPart1(input string) string {
	machines := parseInput(input)
	sumTokens := 0
	for _, machine := range machines {
		var minTokens *int = nil
		for a := 0; a < 100; a++ {
			for b := 0; b < 100; b++ {
				if (machine.buttonA.x*a+machine.buttonB.x*b) == machine.prizeX && (machine.buttonA.y*a+machine.buttonB.y*b) == machine.prizeY {
					tokens := a*3 + b
					if minTokens == nil || tokens < *minTokens {
						minTokens = &tokens
					}
				}
			}
		}
		if minTokens != nil {
			sumTokens += *minTokens
		}
	}
	return strconv.Itoa(sumTokens)
}

func ProcessPart2(input string) string {
	machines := parseInput(input)
	for i, machine := range machines {
		machine.prizeX += 10000000000000
		machine.prizeY += 10000000000000
		machines[i] = machine
	}

	sumTokens := 0

	for _, machine := range machines {
		lp := golp.NewLP(0, 2)
		lp.AddConstraint([]float64{float64(machine.buttonA.x), float64(machine.buttonB.x)}, golp.EQ, float64(machine.prizeX))
		lp.AddConstraint([]float64{float64(machine.buttonA.y), float64(machine.buttonB.y)}, golp.EQ, float64(machine.prizeY))
		lp.SetObjFn([]float64{3.0, 1.0})
		lp.SetInt(0, true)
		lp.SetInt(1, true)

		solution := lp.Solve()
		vars := lp.Variables()

		if solution == golp.OPTIMAL {
			buttonA := int(vars[0])
			buttonB := int(vars[1])
			if buttonA*machine.buttonA.x+buttonB*machine.buttonB.x == machine.prizeX && buttonA*machine.buttonA.y+buttonB*machine.buttonB.y == machine.prizeY {
				tokens := buttonA*3 + buttonB
				sumTokens += tokens
			}
		}
	}

	return strconv.Itoa(sumTokens)
}

type button struct {
	x, y int
}

type clawMachine struct {
	buttonA, buttonB button
	prizeX, prizeY   int
}

func parseInput(input string) []clawMachine {
	result := make([]clawMachine, 0)
	lines := strings.Split(input, "\n")
	buttonA := button{}
	buttonB := button{}
	prizeX := 0
	prizeY := 0
	for _, line := range lines {
		trimmedLine := strings.TrimRight(line, "\r")
		if strings.Contains(trimmedLine, "Button A") {
			buttonA = parseButton(trimmedLine)
		} else if strings.Contains(trimmedLine, "Button B") {
			buttonB = parseButton(trimmedLine)
		} else if strings.Contains(trimmedLine, "Prize") {
			prizeX, prizeY = parsePrize(trimmedLine)
			result = append(result, clawMachine{buttonA, buttonB, prizeX, prizeY})
		}
	}
	return result
}

func parseButton(input string) button {
	parts := strings.Split(strings.Split(input, ": ")[1], ", ")
	x, err := strconv.Atoi(strings.Split(parts[0], "+")[1])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(strings.Split(parts[1], "+")[1])
	if err != nil {
		panic(err)
	}
	return button{x, y}
}

func parsePrize(input string) (int, int) {
	parts := strings.Split(strings.Split(input, ": ")[1], ", ")
	x, err := strconv.Atoi(strings.Split(parts[0], "=")[1])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(strings.Split(parts[1], "=")[1])
	if err != nil {
		panic(err)
	}
	return x, y
}
