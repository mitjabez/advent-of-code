package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func input() []string {
	scanner := bufio.NewScanner(os.Stdin)

	var input = []string{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		input = append(input, line)
	}

	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintf("Error reading standard input: %v", err))
	}

	return input
}

func maxRating(powerBank string, start int, end int) (int, byte) {
	maxRating := byte('0')
	pos := -1
	for i := start; i < end; i++ {
		rating := powerBank[i]
		if rating > maxRating {
			maxRating = rating
			pos = i
		}
		if rating == '9' {
			break
		}
	}
	return pos, maxRating
}

func solvePowerBank(line string, startPos int, result string, depth int, maxDepth int) string {
	if depth == maxDepth {
		return result

	}

	end := len(line) - (maxDepth - depth - 1)
	pos, rating := maxRating(line, startPos, end)
	return solvePowerBank(line, pos+1, result+string(rating), depth+1, maxDepth)
}

func solve(input []string) (int, int) {
	p1 := 0
	p2 := 0

	for _, line := range input {
		joltage1, _ := strconv.Atoi(solvePowerBank(line, 0, "", 0, 2))
		joltage2, _ := strconv.Atoi(solvePowerBank(line, 0, "", 0, 12))
		p1 += joltage1
		p2 += joltage2
	}
	return p1, p2
}

func main() {
	input := input()
	p1, p2 := solve(input)
	fmt.Println(p1)
	fmt.Println(p2)
}
