package day03

import (
	"regexp"
	"strconv"
)

func ProcessPart1(input string) string {
	mulsPattern := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	muls := mulsPattern.FindAllStringSubmatch(input, -1)
	sum := 0
	for _, mul := range muls {
		first, _ := strconv.Atoi(mul[1])
		second, _ := strconv.Atoi(mul[2])
		sum += first * second
	}
	return strconv.Itoa(sum)
}

func ProcessPart2(input string) string {
	pattern := regexp.MustCompile(`(mul\((\d{1,3}),(\d{1,3})\))|(do\(\))|(don't\(\))`)
	matches := pattern.FindAllStringSubmatch(input, -1)
	enabled := true
	sum := 0
	for _, match := range matches {
		switch {
		case match[0] == "do()":
			enabled = true
		case match[0] == "don't()":
			enabled = false
		case enabled:
			first, _ := strconv.Atoi(match[2])
			second, _ := strconv.Atoi(match[3])
			sum += first * second
		}
	}
	return strconv.Itoa(sum)
}
