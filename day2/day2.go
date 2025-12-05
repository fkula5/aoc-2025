package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	data, _ := os.Open("data.txt")

	defer data.Close()

	dataStat, _ := data.Stat()

	dataBytes := make([]byte, dataStat.Size())
	_, _ = data.Read(dataBytes)

	formatedData := strings.Split(string(dataBytes), ",")

	answer1 := int64(0)
	answer2 := int64(0)

	for _, v := range formatedData {
		line := strings.Split(v, "-")
		l, _ := strconv.ParseInt(strings.TrimSpace(line[0]), 10, 64)
		r, _ := strconv.ParseInt(strings.TrimSpace(line[1]), 10, 64)

		for i := l; i <= r; i++ {
			s := strconv.FormatInt(i, 10)
			mid := len(s) / 2
			left := s[:mid]
			right := s[mid:]
			if left == "" {
				left = "0"
			}
			if left == right {
				answer1 += i
			}

			strLen := len(s)
			for pLen := 1; pLen < strLen; pLen++ {
				pattern := s[:pLen]
				repetitions := strLen / pLen
				repeated := strings.Repeat(pattern, repetitions)
				if repeated == s {
					answer2 += i
					break
				}
			}
		}
	}
	fmt.Printf("Answer1: %d\n", answer1)
	fmt.Printf("Answer2: %d\n", answer2)
}
