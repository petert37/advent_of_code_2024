package day11

import (
	"strconv"
	"strings"
)

func ProcessPart1(input string) string {
	stones := parseInput(input)
	for i := 0; i < 25; i++ {
		stones = step(stones)
	}
	return strconv.Itoa(len(stones))
}

func ProcessPart2(input string) string {
	stones := parseInput(input)
	sum := 0
	cache := make(map[[2]int]int)
	for _, stone := range stones {
		sum += count(stone, 0, 75, &cache)
	}
	return strconv.Itoa(sum)
}

func parseInput(input string) []int {
	result := make([]int, 0)
	trimmedInput := strings.TrimRight(input, "\r\n")
	nums := strings.Split(trimmedInput, " ")
	for _, num := range nums {
		parsed, _ := strconv.Atoi(num)
		result = append(result, parsed)
	}
	return result
}

func step(stones []int) []int {
	result := make([]int, 0, len(stones))
	for _, stone := range stones {
		if stone == 0 {
			result = append(result, 1)
			continue
		}
		str := strconv.Itoa(stone)
		length := len(str)
		if length%2 == 0 {
			firstHalf, err := strconv.Atoi(str[:length/2])
			if err != nil {
				panic(err)
			}
			secondHalf, err := strconv.Atoi(str[length/2:])
			if err != nil {
				panic(err)
			}
			result = append(result, firstHalf, secondHalf)
			continue
		}
		result = append(result, stone*2024)
	}
	return result
}

func next(stone int) []int {
	if stone == 0 {
		return []int{1}
	}
	str := strconv.Itoa(stone)
	length := len(str)
	if length%2 == 0 {
		firstHalf, err := strconv.Atoi(str[:length/2])
		if err != nil {
			panic(err)
		}
		secondHalf, err := strconv.Atoi(str[length/2:])
		if err != nil {
			panic(err)
		}
		return []int{firstHalf, secondHalf}
	}
	return []int{stone * 2024}
}

func count(stone int, depth int, maxDepth int, cache *map[[2]int]int) int {
	if depth == maxDepth {
		return 1
	}
	key := [2]int{stone, depth}
	if val, ok := (*cache)[key]; ok {
		return val
	}
	sum := 0
	for _, nextStone := range next(stone) {
		sum += count(nextStone, depth+1, maxDepth, cache)
	}
	(*cache)[key] = sum
	return sum
}
