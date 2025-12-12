package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func calc(nums []int, op string) int {
	res := nums[0]
	for i := 1; i < len(nums); i++ {
		switch op {
		case "+":
			res += nums[i]
		case "-":
			res -= nums[i]
		case "*":
			res *= nums[i]
		case "/":
			res /= nums[i]
		}
	}
	return res
}

func main() {
	data, _ := os.Open("data.txt")
	defer data.Close()

	var lines []string
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var fields [][]string
	for _, line := range lines {
		fields = append(fields, strings.Fields(line))
	}

	rows := len(fields)
	cols := len(fields[0])

	problems := make([][]string, cols)
	for i := range problems {
		problems[i] = make([]string, rows)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			problems[j][i] = fields[i][j]
		}
	}

	sum1 := 0
	for _, p := range problems {
		var nums []int
		for i := 0; i < len(p)-1; i++ {
			n, _ := strconv.Atoi(p[i])
			nums = append(nums, n)
		}
		sum1 += calc(nums, p[len(p)-1])
	}

	sum2 := 0
	maxLen := 0
	for _, line := range lines {
		if len(line) > maxLen {
			maxLen = len(line)
		}
	}

	var nums []int
	for i := maxLen - 1; i >= 0; i-- {
		var col []rune
		op := ""
		empty := true

		for j := 0; j < len(lines); j++ {
			ch := ' '
			if i < len(lines[j]) {
				ch = rune(lines[j][i])
			}
			if ch != ' ' {
				empty = false
			}
			if ch == '+' || ch == '*' || ch == '-' || ch == '/' {
				op = string(ch)
			} else {
				col = append(col, ch)
			}
		}

		if !empty {
			s := strings.ReplaceAll(string(col), " ", "")
			if s != "" {
				n, _ := strconv.Atoi(s)
				nums = append(nums, n)
			}
		}

		if op != "" {
			sum2 += calc(nums, op)
			nums = []int{}
		}
	}

	fmt.Println("Part 1:", sum1)
	fmt.Println("Part 2:", sum2)
}
