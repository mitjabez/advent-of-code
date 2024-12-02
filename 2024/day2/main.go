package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func getSign(a int) int {
	if a < 0 {
		return -1
	}
	return 1
}

func delete(slice []int, i int) []int {
	newSlice := make([]int, len(slice))
	copy(newSlice, slice)
	removed := append(newSlice[:i], newSlice[i+1:]...)
	return removed
}

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

func isSafe(report []int) (bool, int) {
	sign := getSign(report[1] - report[0])
	for i := 1; i < len(report); i++ {
		delta := report[i] - report[i-1]
		if getSign(delta) != sign || abs(delta) < 1 || abs(delta) > 3 {
			return false, i
		}
	}
	return true, 0
}

func part1(input [][]int) int {
	total := 0
	for _, report := range input {
		isSafeReport, _ := isSafe(report)
		if isSafeReport {
			total++
		}
	}
	return total
}

func part2(input [][]int) int {
	total := 0
	for _, report := range input {
		isSafeReport, unsafeNum := isSafe(report)
		if !isSafeReport {
			for i := max(unsafeNum-2, 0); i <= unsafeNum; i++ {
				cleanedReport := delete(report, i)
				isSafeReport, _ = isSafe(cleanedReport)
				if isSafeReport {
					break
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
