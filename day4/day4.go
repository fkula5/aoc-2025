package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	data, _ := os.Open("data.txt")
	defer data.Close()

	var lines []string
	scanner := bufio.NewScanner(data)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	grid := make([][]byte, len(lines))
	for i := range lines {
		grid[i] = []byte(lines[i])
	}

	firstAccessible := countAccessibleRolls(grid)
	totalAccessible := 0

	for {
		accessible := countAccessibleRolls(grid)
		if accessible == 0 {
			break
		}

		totalAccessible += accessible
	}

	fmt.Printf("First accessible rolls: %d\n", firstAccessible)
	fmt.Printf("Total accessible rolls: %d\n", totalAccessible)
}

func countAccessibleRolls(grid [][]byte) int {
	if len(grid) == 0 {
		return 0
	}

	rows := len(grid)
	cols := len(grid[0])
	count := 0

	directions := [][2]int{
		{-1, 0},
		{-1, 1},
		{0, 1},
		{1, 1},
		{1, 0},
		{1, -1},
		{0, -1},
		{-1, -1},
	}

	var accessible [][2]int

	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			if grid[r][c] == '@' {
				adjacentRolls := 0
				for _, dir := range directions {
					nr, nc := r+dir[0], c+dir[1]
					if nr >= 0 && nr < rows && nc >= 0 && nc < cols {
						if grid[nr][nc] == '@' {
							adjacentRolls++
						}
					}
				}

				if adjacentRolls < 4 {
					count++
					accessible = append(accessible, [2]int{r, c})
				}
			}
		}
	}

	for _, pos := range accessible {
		grid[pos[0]][pos[1]] = 'x'
	}

	return count
}
