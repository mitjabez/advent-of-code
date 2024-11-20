package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
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

func numMap(items []string) map[string]int {
	nm := make(map[string]int)
	for _, s := range items {
		if s == "" {
			continue
		}
		nm[s] = 0
	}
	return nm
}

func numList(items []string) []string {
	nm := []string{}
	for _, s := range items {
		if s == "" {
			continue
		}
		nm = append(nm, s)
	}
	return nm
}
func main() {
	input := input()

	part1 := 0
	scratchcards := make(map[int]int)

	for n, matchLine := range input {
		numberSets := strings.Split(strings.Split(matchLine, ":")[1], " | ")
		winningNumbers := numList(strings.Split(numberSets[0], " "))
		myNumbers := numMap(strings.Split(numberSets[1], " "))
		matchScore := 0
		matches := 0
		scratchcards[n]++

		for _, num := range winningNumbers {
			_, ok := myNumbers[num]
			if ok {
				matches++
				if matchScore == 0 {
					matchScore = 1
				} else {
					matchScore *= 2
				}
			}
		}

		if matches > 0 {
			for i := n + 1; i < min(n+matches+1, len(input)); i++ {
				scratchcards[i] += scratchcards[n]
			}
		}

		part1 += matchScore
	}

	part2 := 0
	for _, v := range scratchcards {
		part2 += v
	}

	fmt.Println(part1)
	fmt.Println(part2)
}
