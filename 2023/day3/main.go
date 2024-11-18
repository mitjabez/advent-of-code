package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Result struct {
	Part1 int
	Part2 int
}

func clamp(n int, maxn int, minn int) int {
	return min(max(n, minn), maxn)
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isSymbol(c byte) bool {
	return c != '.' && !isDigit(c)
}

func hasAdjacentSymbol(part []byte) bool {
	for _, c := range part {
		if isSymbol(c) {
			return true
		}
	}
	return false
}

type Pos struct {
	x int
	y int
}

func solve(input []string) Result {
	digits := []string{}

	for y, line := range input {
		lastDigit := ""
		hasAdjSymbol := false
		for x, c := range []byte(line) {
			if !isDigit(c) {
				if hasAdjSymbol {
					digits = append(digits, lastDigit)
				}
				lastDigit = ""
				hasAdjSymbol = false
				continue
			}

			lastDigit += string(c)

			if y > 0 {
				above := input[y-1][max(x-1, 0):min(x+2, len(line))]
				if hasAdjacentSymbol([]byte(above)) {
					hasAdjSymbol = true
				}
			}

			if y < len(input)-1 {
				below := input[y+1][max(x-1, 0):min(x+2, len(line))]
				if hasAdjacentSymbol([]byte(below)) {
					hasAdjSymbol = true
				}
			}

			if x > 0 && isSymbol(line[x-1]) {
				hasAdjSymbol = true
			}

			if x < len(line)-1 && isSymbol(line[x+1]) {
				hasAdjSymbol = true
			}
		}

		if hasAdjSymbol {
			digits = append(digits, lastDigit)
		}
	}

	part1 := 0
	for _, digit := range digits {
		num, _ := strconv.Atoi(digit)
		part1 += num
	}

	return Result{Part1: part1}
}

func input() []string {
	scanner := bufio.NewScanner(os.Stdin)

	var input = []string{}
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			input = append(input, line)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintf("Error reading standard input: %v", err))
	}

	return input
}

func main() {
	input := input()

	result := solve(input)
	fmt.Println("Part1:", result.Part1)
}
