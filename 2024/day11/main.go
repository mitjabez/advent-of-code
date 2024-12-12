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

	nums := make(map[int]int)
	for _, s := range stones {
		nums[s] = 1
	}
	for i := 0; i < times; i++ {
		newNums := make(map[int]int)

		for s, v := range nums {
			newNums[s] -= v
			if s == 0 {
				newNums[1] += v
				continue
			}
			digits := int(math.Log10(float64(s)) + 1)
			if digits%2 == 0 {
				n1 := s / int(math.Pow10(digits/2))
				n2 := s % int(math.Pow10(digits/2))
				newNums[n1] += v
				newNums[n2] += v
				continue
			}
			newNums[s*2024] += v
		}

		for k, v := range newNums {
			nums[k] += v
			if nums[k] <= 0 {
				delete(nums, k)
			}
		}
	}

	total := 0
	for _, v := range nums {
		total += v
	}
	return total
}

func main() {
	input := input()
	fmt.Println(solve(input, 25))
	fmt.Println(solve(input, 75))
}
