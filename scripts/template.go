package main

import (
	"bufio"
	"fmt"
	"os"
)

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

	totalPart1 := 0
	// totalPart2 := 0

	fmt.Println("Part1:", totalPart1)
	// fmt.Println("Part2:", totalPart2)
}
