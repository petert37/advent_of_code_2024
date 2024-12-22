package day22

import (
	"strconv"
	"strings"
)

func ProcessPart1(input string) string {
	initialSecretNumbers := parseInput(input)
	result := 0
	for _, secretNumber := range initialSecretNumbers {
		pseudoRandomized := secretNumber
		for i := 0; i < 2000; i++ {
			pseudoRandomized = nextSecretNumber(pseudoRandomized)
		}
		result += pseudoRandomized
	}
	return strconv.Itoa(result)
}

func ProcessPart2(input string) string {
	initialSecretNumbers := parseInput(input)
	prices := make([][]int, 0)
	priceChanges := make([][]int, 0)
	for _, secretNumber := range initialSecretNumbers {
		monkeyPrices := make([]int, 0)
		monkeyPriceChanges := make([]int, 0)
		pseudoRandomized := secretNumber
		price := pseudoRandomized % 10
		monkeyPrices = append(monkeyPrices, price)
		for i := 0; i < 2000; i++ {
			pseudoRandomized = nextSecretNumber(pseudoRandomized)
			nextPrice := pseudoRandomized % 10
			monkeyPrices = append(monkeyPrices, nextPrice)
			priceChange := nextPrice - price
			price = nextPrice
			monkeyPriceChanges = append(monkeyPriceChanges, priceChange)
		}
		prices = append(prices, monkeyPrices)
		priceChanges = append(priceChanges, monkeyPriceChanges)
	}
	sum := 0
	count := 0
	results := make(chan int)
	for c0 := -10; c0 <= 10; c0++ {
		for c1 := -10; c1 <= 10; c1++ {
			for c2 := -10; c2 <= 10; c2++ {
				for c3 := -10; c3 <= 10; c3++ {
					count++
					go getComboPrice(c0, c1, c2, c3, &priceChanges, &prices, results)
				}
			}
		}
	}
	for i := 0; i < count; i++ {
		comboSum := <-results
		if comboSum > sum {
			sum = comboSum
		}
	}
	return strconv.Itoa(sum)
}

func parseInput(input string) []int {
	result := make([]int, 0)
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		trimmedLine := strings.TrimRight(line, "\r")
		parsed, err := strconv.Atoi(trimmedLine)
		if err != nil {
			panic(err)
		}
		result = append(result, parsed)
	}
	return result
}

func nextSecretNumber(secretNumber int) int {
	mul := secretNumber * 64
	result := mul ^ secretNumber
	result = result % 16777216
	div := result / 32
	result = div ^ result
	result = result % 16777216
	mul = result * 2048
	result = mul ^ result
	result = result % 16777216
	return result
}

func getComboPrice(c0, c1, c2, c3 int, priceChanges *([][]int), prices *([][]int), result chan int) {
	comboSum := 0
	for i, monkeyPriceChanges := range *priceChanges {
		for j := 0; j < len(monkeyPriceChanges)-3; j++ {
			if monkeyPriceChanges[j] == c0 && monkeyPriceChanges[j+1] == c1 && monkeyPriceChanges[j+2] == c2 && monkeyPriceChanges[j+3] == c3 {
				price := (*prices)[i][j+4]
				comboSum += price
				break
			}
		}
	}
	result <- comboSum
}
