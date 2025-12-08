package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	file, _ := os.Open("data.txt")
	defer file.Close()

	var ranges [][]int64
	idsMap := make(map[int64]struct{})

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}

		if strings.Contains(line, "-") {
			parts := strings.Split(line, "-")
			l, _ := strconv.ParseInt(parts[0], 10, 64)
			r, _ := strconv.ParseInt(parts[1], 10, 64)
			ranges = append(ranges, []int64{l, r})
		} else {
			d, _ := strconv.ParseInt(line, 10, 64)
			idsMap[d] = struct{}{}
		}
	}

	sort.Slice(ranges, func(i, j int) bool {
		return ranges[i][0] < ranges[j][0]
	})

	var mergedRanges [][]int64
	if len(ranges) > 0 {
		current := ranges[0]
		for i := 1; i < len(ranges); i++ {
			next := ranges[i]

			if next[0] <= current[1]+1 {
				if next[1] > current[1] {
					current[1] = next[1]
				}
			} else {
				mergedRanges = append(mergedRanges, current)
				current = next
			}
		}
		mergedRanges = append(mergedRanges, current)
	}

	freshCount := 0
	for id := range idsMap {
		for _, r := range mergedRanges {
			if id >= r[0] && id <= r[1] {
				freshCount++
				break
			}
			if r[0] > id {
				break
			}
		}
	}

	var totalCoverage int64 = 0
	for _, r := range mergedRanges {
		count := r[1] - r[0] + 1
		totalCoverage += count
	}

	fmt.Printf("Fresh IDs (matches): %d\n", freshCount)
	fmt.Printf("Total Fresh IDs (coverage): %d\n", totalCoverage)
}
