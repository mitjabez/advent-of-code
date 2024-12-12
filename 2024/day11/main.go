package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func input() []int {
	scanner := bufio.NewScanner(os.Stdin)

	var stones = []int{}
	for scanner.Scan() {
		line := scanner.Text()
		for _, stone := range strings.Fields(line) {
			s, _ := strconv.Atoi(stone)
			stones = append(stones, s)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintf("Error reading standard input: %v", err))
	}

	return stones
}

func solve(stones []int, times int) int {
	for i := 0; i < times; i++ {
		after := []int{}
		for _, s := range stones {
			if s == 0 {
				after = append(after, 1)
				continue
			}
			digits := int(math.Log10(float64(s)) + 1)
			if digits%2 == 0 {
				after = append(after, s/int(math.Pow10(digits/2)))
				after = append(after, s%int(math.Pow10(digits/2)))
				continue
			}
			after = append(after, s*2024)
		}
		stones = after
	}
	return len(stones)
}

func main() {
	stones := input()
	fmt.Println(solve(stones, 25))
	fmt.Println(solve(stones, 75))
}
