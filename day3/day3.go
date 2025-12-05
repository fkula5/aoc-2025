package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func findMaxNDigits(line string, n int) int {
	if len(line) < n {
		return 0
	}

	result := make([]byte, 0, n)
	rem := n
	start := 0

	for rem > 0 {
		maxDigit := byte('0')
		maxPos := start

		endPos := len(line) - rem + 1
		for i := start; i < endPos; i++ {
			if line[i] > maxDigit {
				maxDigit = line[i]
				maxPos = i
			}
		}

		result = append(result, maxDigit)
		start = maxPos + 1
		rem--
	}

	num, _ := strconv.Atoi(string(result))
	return num
}

func main() {
	data, _ := os.Open("data.txt")
	defer data.Close()

	sumPart1 := 0
	sumPart2 := 0

	scanner := bufio.NewScanner(data)

	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		max2 := findMaxNDigits(line, 2)
		sumPart1 += max2

		max12 := findMaxNDigits(line, 12)
		sumPart2 += max12
	}

	fmt.Printf("Part 1 : %d\n", sumPart1)
	fmt.Printf("Part 2 : %d\n", sumPart2)
}
