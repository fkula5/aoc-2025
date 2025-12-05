package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	data, err := os.Open("day1.txt")
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}
	defer data.Close()

	dial := 50
	ticks := 0

	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		line := scanner.Text()

		if len(line) < 2 {
			continue
		}

		s := line[:1]
		n, _ := strconv.Atoi(line[1:])

		startDial := dial

		if s == "L" {
			for i := 0; i < n; i++ {
				dial--
				if dial < 0 {
					dial = 99
				}
				if dial == 0 {
					ticks++
				}
			}
		}

		if s == "R" {
			for i := 0; i < n; i++ {
				dial++
				if dial > 99 {
					dial = 0
				}
				if dial == 0 {
					ticks++
				}
			}
		}

		resultLine := fmt.Sprintf("sign: %s, num: %d, Start: %d, End: %d, Ticks: %d\n", s, n, startDial, dial, ticks)

		fmt.Print(resultLine)
	}

	fmt.Printf("Total ticks: %d\n", ticks)
}
