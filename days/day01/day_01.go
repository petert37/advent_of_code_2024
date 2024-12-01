package day01

import (
	"math"
	"slices"
	"strconv"
	"strings"
)

func ProcessPart1(input string) string {
	lines := strings.Split(input, "\n")

	list1 := make([]int, 0)
	list2 := make([]int, 0)

	for _, line := range lines {
		fields := strings.Fields(line)
		item1, _ := strconv.Atoi(fields[0])
		item2, _ := strconv.Atoi(fields[1])
		list1 = append(list1, item1)
		list2 = append(list2, item2)
	}

	slices.Sort(list1)
	slices.Sort(list2)

	sumDiff := 0

	for i := 0; i < len(list1); i++ {
		item1 := list1[i]
		item2 := list2[i]
		sumDiff += int(math.Abs(float64(item1 - item2)))
	}

	return strconv.Itoa(sumDiff)
}

func ProcessPart2(input string) string {
	lines := strings.Split(input, "\n")

	list1 := make([]int, 0)
	list2 := make(map[int]int)

	for _, line := range lines {
		fields := strings.Fields(line)
		item1, _ := strconv.Atoi(fields[0])
		item2, _ := strconv.Atoi(fields[1])
		list1 = append(list1, item1)
		if currentValue, exists := list2[item2]; exists {
			list2[item2] = currentValue + 1
		} else {
			list2[item2] = 1
		}
	}

	sum := 0
	for i := 0; i < len(list1); i++ {
		item1 := list1[i]
		if item2, exists := list2[item1]; exists {
			sum += item1 * item2
		}
	}

	return strconv.Itoa(sum)
}
