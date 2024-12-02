package day02

import (
	"log"
	"slices"
	"strconv"
	"strings"
)

func ProcessPart1(input string) string {
	lines := strings.Split(input, "\n")

	count := 0
	for _, line := range lines {
		report := strings.Fields(strings.TrimRight(line, "\r"))
		safe, err := isSafe(report)
		if err != nil {
			log.Fatal(err)
		}
		if safe {
			count++
		}
	}

	return strconv.Itoa(count)
}

func ProcessPart2(input string) string {
	lines := strings.Split(input, "\n")

	count := 0
	for _, line := range lines {
		report := strings.Fields(strings.TrimRight(line, "\r"))
		safe, err := isSafe(report)
		if err != nil {
			log.Fatal(err)
		}
		if safe {
			count++
			continue
		}
		for i := 0; i < len(report); i++ {
			tolerated_report := slices.Concat(report[:i], report[i+1:])
			safe, err := isSafe(tolerated_report)
			if err != nil {
				log.Fatal(err)
			}
			if safe {
				count++
				break
			}
		}
	}

	return strconv.Itoa(count)
}

func isSafe(report []string) (bool, error) {
	first, err := strconv.Atoi(report[0])
	if err != nil {
		return false, err
	}
	second, err := strconv.Atoi(report[1])
	if err != nil {
		return false, err
	}
	if first == second {
		return false, nil
	}
	increasing := first < second

	for i := 0; i < len(report)-1; i++ {
		current, err := strconv.Atoi(report[i])
		if err != nil {
			return false, err
		}
		next, err := strconv.Atoi(report[i+1])
		if err != nil {
			return false, err
		}
		diff := next - current
		if increasing {
			if diff < 1 || diff > 3 {
				return false, nil
			}
		} else {
			if diff > -1 || diff < -3 {
				return false, nil
			}
		}
	}

	return true, nil
}
