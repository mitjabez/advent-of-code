package main

import (
	"bufio"
	"fmt"
	"os"
)

var words = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func solve(line []byte, handleLetters bool) int {
	first := 0
	last := 0

	for i, c := range line {
		num := 0
		if c >= '0' && c <= '9' {
			num = int(c - '0')
		} else if handleLetters {
			for j := i; j <= min(i+j, len(line)); j++ {
				num = words[string(line[i:j])]
				if num > 0 {
					break
				}
			}
		}

		if num > 0 {
			if first == 0 {
				first = num
			}
			last = num
		}
	}
	return first*10 + last
}

func input() []string {
	scanner := bufio.NewScanner(os.Stdin)

	var input = []string{}
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintf("Error reading standard input: %v", err))
	}

	return input
}

func main() {
	input := input()

	totalPart1 := 0
	totalPart2 := 0
	for _, line := range input {
		totalPart1 += solve([]byte(line), false)
		totalPart2 += solve([]byte(line), true)
	}

	fmt.Println("Part1:", totalPart1)
	fmt.Println("Part2:", totalPart2)
}
