package day21

import (
	"math"
	"strconv"
	"strings"
)

func ProcessPart1(input string) string {
	codes, numericCodes := parseInput(input)
	numericPaths := makeNumericPaths()
	directionalPaths := makeDirectionalPaths()
	sumComplexity := 0
	for i, code := range codes {
		currentPosition1 := 10
		sum := 0
		for _, pos := range code { //numeric keypad operated by robot 1
			paths1 := numericPaths[numericPathFromTo{currentPosition1, pos}]
			digitSum := math.MaxInt
			for _, path1 := range paths1 {
				path1A := make([]directionalInput, len(path1))
				for i, dir := range path1 {
					path1A[i] = directionalInput(dir)
				}
				path1A = append(path1A, directionalInputActivate) //intermediate robot 1
				currentDirection1 := directionalInputActivate
				sum1 := 0
				for _, dir1 := range path1A {
					paths2 := directionalPaths[directionalPathFromTo{currentDirection1, dir1}]
					minLen2 := math.MaxInt
					for _, path2 := range paths2 {
						path2A := make([]directionalInput, len(path2))
						for i, dir := range path2 {
							path2A[i] = directionalInput(dir)
						}
						path2A = append(path2A, directionalInputActivate) //intermediate robot 2
						currentDirection2 := directionalInputActivate
						len2 := 0
						for _, dir2 := range path2A {
							paths3 := directionalPaths[directionalPathFromTo{currentDirection2, dir2}]
							for ip3, path3 := range paths3 {
								path3A := make([]directionalInput, len(path3))
								for i, dir := range path3 {
									path3A[i] = directionalInput(dir)
								}
								path3A = append(path3A, directionalInputActivate) //manual input
								if ip3 == 0 {
									len2 += len(path3A)
								}
							}
							currentDirection2 = dir2
						}
						if len2 < minLen2 {
							minLen2 = len2
						}
					}
					currentDirection1 = dir1
					sum1 += minLen2
				}
				if sum1 < digitSum {
					digitSum = sum1
				}
			}
			currentPosition1 = pos
			sum += digitSum
		}
		numericCode := numericCodes[i]
		sumComplexity += sum * numericCode
	}
	return strconv.Itoa(sumComplexity)
}

func ProcessPart2(input string) string {
	codes, numericCodes := parseInput(input)
	numericPaths := makeNumericPaths()
	directionalPaths := makeDirectionalPaths()
	sumComplexity := 0
	cache := make(map[string]int)
	for i, code := range codes {
		currentPosition := 10
		sum := 0
		for _, pos := range code { //numeric keypad operated by robot 1
			paths := numericPaths[numericPathFromTo{currentPosition, pos}]

			digitSum := math.MaxInt
			for _, path := range paths {
				pathStepCount := calc(path, &directionalPaths, &cache, 25)
				if pathStepCount < digitSum {
					digitSum = pathStepCount
				}
			}
			currentPosition = pos
			sum += digitSum
		}
		numericCode := numericCodes[i]
		sumComplexity += sum * numericCode
	}
	return strconv.Itoa(sumComplexity)
}

type direction int

const (
	directionUp direction = iota
	directionDown
	directionLeft
	directionRight
)

func (d direction) String() string {
	switch d {
	case directionUp:
		return "^"
	case directionDown:
		return "v"
	case directionLeft:
		return "<"
	case directionRight:
		return ">"
	default:
		return "?"
	}
}

type directionalInput int

const (
	directionalInputUp directionalInput = iota
	directionalInputDown
	directionalInputLeft
	directionalInputRight
	directionalInputActivate
)

func (d directionalInput) String() string {
	switch d {
	case directionalInputUp:
		return "^"
	case directionalInputDown:
		return "v"
	case directionalInputLeft:
		return "<"
	case directionalInputRight:
		return ">"
	case directionalInputActivate:
		return "A"
	default:
		return "?"
	}
}

func parseInput(input string) (codes [][]int, numeric []int) {
	codes = make([][]int, 0)
	numeric = make([]int, 0)
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		trimmedLine := strings.TrimRight(line, "\r")
		row := make([]int, 0)
		for _, c := range trimmedLine {
			parsed, err := strconv.ParseInt(string(c), 16, 0)
			if err != nil {
				panic(err)
			}
			row = append(row, int(parsed))
		}
		codes = append(codes, row)
		numericTrimmedLine := strings.TrimRight(trimmedLine, "A")
		parsed, err := strconv.Atoi(numericTrimmedLine)
		if err != nil {
			panic(err)
		}
		numeric = append(numeric, parsed)
	}
	return
}

type numericStep struct {
	path     []direction
	position int
}

type numericPathFromTo struct {
	p1, p2 int
}

type directionalPathFromTo struct {
	p1, p2 directionalInput
}

type directionalStep struct {
	path     []direction
	position directionalInput
}

func makeNumericPaths() map[numericPathFromTo][][]direction {
	numericPaths := make(map[numericPathFromTo][][]direction)
	for p1 := 0; p1 <= 10; p1++ {
		for p2 := 0; p2 <= 10; p2++ {
			fifo := make([]numericStep, 0)
			fifo = append(fifo, numericStep{path: []direction{}, position: p1})
			visited := make(map[int]int)
			visited[p1] = 0
			var shortestPathLength *int = nil
			shortestPaths := make([][]direction, 0)
			for len(fifo) > 0 {
				current := fifo[0]
				fifo = fifo[1:]
				if current.position == p2 {
					length := len(current.path)
					if shortestPathLength == nil || length < *shortestPathLength {
						shortestPathLength = &length
						shortestPaths = make([][]direction, 0)
						shortestPaths = append(shortestPaths, current.path)
					} else if length == *shortestPathLength {
						shortestPaths = append(shortestPaths, current.path)
					}
				} else {
					for _, dir := range []direction{directionUp, directionDown, directionLeft, directionRight} {
						nextNumeric, ok := getNextNumeric(current.position, dir)
						if ok {
							visitedLen, ok := visited[nextNumeric]
							if !ok || len(current.path)+1 <= visitedLen {
								visited[nextNumeric] = len(current.path) + 1
								newPath := make([]direction, len(current.path))
								copy(newPath, current.path)
								newPath = append(newPath, dir)
								fifo = append(fifo, numericStep{path: newPath, position: nextNumeric})
							}
						}
					}
				}
			}
			numericPaths[numericPathFromTo{p1, p2}] = shortestPaths
		}
	}
	return numericPaths
}

func makeDirectionalPaths() map[directionalPathFromTo][][]direction {
	directionalPaths := make(map[directionalPathFromTo][][]direction)
	for p1 := directionalInputUp; p1 <= directionalInputActivate; p1++ {
		for p2 := directionalInputUp; p2 <= directionalInputActivate; p2++ {
			fifo := make([]directionalStep, 0)
			fifo = append(fifo, directionalStep{path: []direction{}, position: p1})
			visited := make(map[directionalInput]int)
			visited[p1] = 0
			var shortestPathLength *int = nil
			shortestPaths := make([][]direction, 0)
			for len(fifo) > 0 {
				current := fifo[0]
				fifo = fifo[1:]
				if current.position == p2 {
					length := len(current.path)
					if shortestPathLength == nil || length < *shortestPathLength {
						shortestPathLength = &length
						shortestPaths = make([][]direction, 0)
						shortestPaths = append(shortestPaths, current.path)
					} else if length == *shortestPathLength {
						shortestPaths = append(shortestPaths, current.path)
					}
				} else {
					for _, dir := range []direction{directionUp, directionDown, directionLeft, directionRight} {
						nextNumeric, ok := getNextDirectional(current.position, dir)
						if ok {
							visitedLen, ok := visited[nextNumeric]
							if !ok || len(current.path)+1 <= visitedLen {
								visited[nextNumeric] = len(current.path) + 1
								newPath := make([]direction, len(current.path))
								copy(newPath, current.path)
								newPath = append(newPath, dir)
								fifo = append(fifo, directionalStep{path: newPath, position: nextNumeric})
							}
						}
					}
				}
			}
			directionalPaths[directionalPathFromTo{p1, p2}] = shortestPaths
		}
	}
	return directionalPaths
}

func calc(path []direction, directionalPaths *map[directionalPathFromTo][][]direction, cache *map[string]int, depth int) int {
	key := cacheKey(path, depth)
	if val, ok := (*cache)[key]; ok {
		return val
	}
	pathA := make([]directionalInput, len(path))
	for i, dir := range path {
		pathA[i] = directionalInput(dir)
	}
	pathA = append(pathA, directionalInputActivate)
	if depth == 0 {
		result := len(pathA)
		(*cache)[key] = result
		return result
	}
	currentDirection := directionalInputActivate
	stepCount := 0
	for _, dir := range pathA {
		newPaths := (*directionalPaths)[directionalPathFromTo{currentDirection, dir}]
		shortestNewPathLength := math.MaxInt
		for _, newPath := range newPaths {
			newPathStepCount := calc(newPath, directionalPaths, cache, depth-1)
			if newPathStepCount < shortestNewPathLength {
				shortestNewPathLength = newPathStepCount
			}
		}
		stepCount += shortestNewPathLength
		currentDirection = dir
	}
	(*cache)[key] = stepCount
	return stepCount
}

func getNextNumeric(position int, direction direction) (nextNumeric int, ok bool) {
	nextNumeric = -1
	ok = true
	switch position {
	case 0:
		switch direction {
		case directionUp:
			nextNumeric = 2
		case directionDown:
			ok = false
		case directionLeft:
			ok = false
		case directionRight:
			nextNumeric = 10

		}
	case 1:
		switch direction {
		case directionUp:
			nextNumeric = 4
		case directionDown:
			ok = false
		case directionLeft:
			ok = false
		case directionRight:
			nextNumeric = 2
		}
	case 2:
		switch direction {
		case directionUp:
			nextNumeric = 5
		case directionDown:
			nextNumeric = 0
		case directionLeft:
			nextNumeric = 1
		case directionRight:
			nextNumeric = 3
		}
	case 3:
		switch direction {
		case directionUp:
			nextNumeric = 6
		case directionDown:
			nextNumeric = 10
		case directionLeft:
			nextNumeric = 2
		case directionRight:
			ok = false
		}
	case 4:
		switch direction {
		case directionUp:
			nextNumeric = 7
		case directionDown:
			nextNumeric = 1
		case directionLeft:
			ok = false
		case directionRight:
			nextNumeric = 5
		}
	case 5:
		switch direction {
		case directionUp:
			nextNumeric = 8
		case directionDown:
			nextNumeric = 2
		case directionLeft:
			nextNumeric = 4
		case directionRight:
			nextNumeric = 6
		}
	case 6:
		switch direction {
		case directionUp:
			nextNumeric = 9
		case directionDown:
			nextNumeric = 3
		case directionLeft:
			nextNumeric = 5
		case directionRight:
			ok = false
		}
	case 7:
		switch direction {
		case directionUp:
			ok = false
		case directionDown:
			nextNumeric = 4
		case directionLeft:
			ok = false
		case directionRight:
			nextNumeric = 8
		}
	case 8:
		switch direction {
		case directionUp:
			ok = false
		case directionDown:
			nextNumeric = 5
		case directionLeft:
			nextNumeric = 7
		case directionRight:
			nextNumeric = 9
		}
	case 9:
		switch direction {
		case directionUp:
			ok = false
		case directionDown:
			nextNumeric = 6
		case directionLeft:
			nextNumeric = 8
		case directionRight:
			ok = false
		}
	case 10:
		switch direction {
		case directionUp:
			nextNumeric = 3
		case directionDown:
			ok = false
		case directionLeft:
			nextNumeric = 0
		case directionRight:
			ok = false
		}
	default:
		ok = false
	}
	return
}

func getNextDirectional(position directionalInput, direction direction) (nextDirectional directionalInput, ok bool) {
	nextDirectional = -1
	ok = true
	switch position {
	case directionalInputUp:
		switch direction {
		case directionUp:
			ok = false
		case directionDown:
			nextDirectional = directionalInputDown
		case directionLeft:
			ok = false
		case directionRight:
			nextDirectional = directionalInputActivate
		}
	case directionalInputDown:
		switch direction {
		case directionUp:
			nextDirectional = directionalInputUp
		case directionDown:
			ok = false
		case directionLeft:
			nextDirectional = directionalInputLeft
		case directionRight:
			nextDirectional = directionalInputRight
		}
	case directionalInputLeft:
		switch direction {
		case directionUp:
			ok = false
		case directionDown:
			ok = false
		case directionLeft:
			ok = false
		case directionRight:
			nextDirectional = directionalInputDown
		}
	case directionalInputRight:
		switch direction {
		case directionUp:
			nextDirectional = directionalInputActivate
		case directionDown:
			ok = false
		case directionLeft:
			nextDirectional = directionalInputDown
		case directionRight:
			ok = false
		}
	case directionalInputActivate:
		switch direction {
		case directionUp:
			ok = false
		case directionDown:
			nextDirectional = directionalInputRight
		case directionLeft:
			nextDirectional = directionalInputUp
		case directionRight:
			ok = false
		}
	default:
		ok = false
	}
	return
}

func cacheKey(path []direction, depth int) string {
	key := ""
	for _, dir := range path {
		key += dir.String()
	}
	key += "_"
	key += strconv.Itoa(depth)
	return key
}
