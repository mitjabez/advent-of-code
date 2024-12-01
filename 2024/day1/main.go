package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
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

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func main() {
	input := input()
	left := make([]int, len(input))
	right := make([]int, len(input))
	rightLoc := make(map[int]int, len(input))

	for i, line := range input {
		tokens := strings.Split(line, "   ")
		left[i], _ = strconv.Atoi(tokens[0])
		right[i], _ = strconv.Atoi(tokens[1])
		rightLoc[right[i]]++
	}

	slices.Sort(left)
	slices.Sort(right)

	total1 := 0
	total2 := 0
	for i := 0; i < len(input); i++ {
		total1 += abs(right[i] - left[i])
		total2 += left[i] * rightLoc[left[i]]
	}

	fmt.Println(total1)
	fmt.Println(total2)
}
