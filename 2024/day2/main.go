package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func input() [][]int {
	scanner := bufio.NewScanner(os.Stdin)

	var input = [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			numLevels := []int{}
			for _, level := range strings.Split(line, " ") {
				nl, _ := strconv.Atoi(level)
				numLevels = append(numLevels, nl)
			}
			input = append(input, numLevels)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintf("Error reading standard input: %v", err))
	}

	return input
}
func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func isGradual(a int, b int) bool {
	diff := abs(a - b)
	return diff >= 1 && diff <= 3
}

func sign(a int) bool {
	return a < 0
}

func part1(input [][]int) int {
	total := 0
	for _, line := range input {
		isSafeReport := true
		var isNegative bool
		for n, curLevel := range line {
			if n == 0 {
				continue
			}

			lastLevel := line[n-1]
			if n == 1 {
				isNegative = sign(curLevel - lastLevel)
			}
			if !isGradual(curLevel, lastLevel) || sign(curLevel-lastLevel) != isNegative {
				isSafeReport = false
				break
			}
		}
		if isSafeReport {
			total++
		}
	}
	return total
}

func isSafe(report []int) (bool, int) {
	var isNegative bool
	for n, curLevel := range report {
		if n == 0 {
			continue
		}

		lastLevel := report[n-1]
		if n == 1 {
			isNegative = sign(curLevel - lastLevel)
		}
		if !isGradual(curLevel, lastLevel) || sign(curLevel-lastLevel) != isNegative {
			return false, n
		}
	}
	return true, 0
}

func delete(slice []int, i int) []int {
	newSlice := make([]int, len(slice))
	copy(newSlice, slice)
	removed := append(newSlice[:i], newSlice[i+1:]...)
	return removed
}

func part2(input [][]int) int {
	total := 0
	for _, report := range input {
		isSafeReport, unsafeNum := isSafe(report)
		// 314,339 wrong (343?)
		if !isSafeReport {
			cleanedReport := delete(report, unsafeNum-1)
			isSafeReport, _ = isSafe(cleanedReport)
			if !isSafeReport {
				fmt.Println(report)
				cleanedReport = delete(report, unsafeNum)
				isSafeReport, unsafeNum = isSafe(cleanedReport)
				if !isSafeReport && unsafeNum > 1 {
					cleanedReport = delete(report, unsafeNum-2)
					isSafeReport, _ = isSafe(cleanedReport)
				}
			}
		}

		if isSafeReport {
			total++
		}
	}
	return total
}

func main() {
	input := input()

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
