package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func partA(input []string) int {
	total := 0

	for gameId, line := range input {
		var isPossible = true
		for _, gameSet := range strings.Split(strings.Split(line, ":")[1], ";") {
			var cubes = map[string]int{
				"red":   12,
				"green": 13,
				"blue":  14,
			}

			for _, turn := range strings.Split(gameSet, ",") {
				split := strings.Split(strings.Trim(turn, " "), " ")
				numCubes, _ := strconv.Atoi(split[0])
				color := split[1]
				cubes[color] -= numCubes
				if cubes[color] < 0 {
					isPossible = false
					break
				}
			}

			if !isPossible {
				break
			}
		}

		if isPossible {
			total += gameId + 1
		}
	}

	return total
}

func partB(input []string) int {
	total := 0

	for _, line := range input {
		var cubes = map[string]int{
			"red":   0,
			"green": 0,
			"blue":  0,
		}
		for _, gameSet := range strings.Split(strings.Split(line, ":")[1], ";") {

			for _, turn := range strings.Split(gameSet, ",") {
				split := strings.Split(strings.Trim(turn, " "), " ")
				numCubes, _ := strconv.Atoi(split[0])
				color := split[1]
				cubes[color] = max(numCubes, cubes[color])
			}

		}
		total += cubes["red"] * cubes["green"] * cubes["blue"]
	}

	return total
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

	fmt.Println("Part 1:", partA(input))
	fmt.Println("Part 2:", partB(input))
}
