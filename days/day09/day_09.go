package day09

import (
	"strconv"
)

func ProcessPart1(input string) string {
	diskMap := parseInput(input)
	start := 0
	end := len(diskMap) - 1

	for start < end {

		for ; ; start++ {
			if diskMap[start] == -1 {
				break
			}
		}

		for ; ; end-- {
			if diskMap[end] != -1 {
				break
			}
		}

		if start >= end {
			break
		}

		diskMap[start] = diskMap[end]
		diskMap[end] = -1
	}

	sum := 0

	for i, v := range diskMap {
		if v != -1 {
			sum += i * v
		}
	}

	return strconv.Itoa(sum)
}

func ProcessPart2(input string) string {
	diskMap := parseInput(input)
	fileId := 0

	for i := len(diskMap) - 1; i >= 0; i-- {
		if diskMap[i] != -1 {
			fileId = diskMap[i]
			break
		}
	}

	for ; fileId >= 0; fileId-- {
		sourceStart := 0
		soueceEnd := len(diskMap) - 1

		for i := 0; i < len(diskMap); i++ {
			if diskMap[i] == fileId {
				sourceStart = i
				for j := i + 1; j < len(diskMap); j++ {
					if diskMap[j] != fileId {
						soueceEnd = j - 1
						break
					}
				}
				break
			}
		}

		sourceLength := soueceEnd - sourceStart + 1

		for i := 0; i < sourceStart; i++ {
			targetStart := -1
			targetEnd := -1

			if diskMap[i] == -1 {
				targetStart = i

				for j := i + 1; j < len(diskMap); j++ {
					if diskMap[j] != -1 {
						targetEnd = j - 1
						break
					}
				}

				if targetEnd != -1 && targetEnd-targetStart+1 >= sourceLength {

					for j := 0; j < sourceLength; j++ {
						diskMap[targetStart+j] = fileId
						diskMap[sourceStart+j] = -1
					}

					break
				}
			}
		}
	}

	sum := 0
	for i, v := range diskMap {
		if v != -1 {
			sum += i * v
		}
	}

	return strconv.Itoa(sum)
}

func parseInput(input string) []int {
	result := make([]int, 0)
	for i, r := range input {
		num, _ := strconv.Atoi(string(r))
		var c int
		if i%2 == 0 {
			c = i / 2
		} else {
			c = -1
		}
		for j := 0; j < num; j++ {
			result = append(result, c)
		}
	}
	return result
}
