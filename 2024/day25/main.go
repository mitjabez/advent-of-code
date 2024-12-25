package main

import (
	"bufio"
	"fmt"
	"os"
)

func input() ([][]int, [][]int) {
	scanner := bufio.NewScanner(os.Stdin)

	locks := [][]int{}
	keys := [][]int{}

	lines := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			// lock
			var ignoreLine int
			if lines[0] == "#####" {
				// lock
				ignoreLine = 0
			} else {
				ignoreLine = len(lines) - 1
			}
			cnt := []int{0, 0, 0, 0, 0}
			for y := 0; y < len(lines); y++ {
				for x := 0; x < len(lines[0]); x++ {
					if y != ignoreLine && lines[y][x] == '#' {
						cnt[x]++
					}
				}
			}
			if lines[0] == "#####" {
				locks = append(locks, cnt)
			} else {
				keys = append(keys, cnt)
			}

			lines = []string{}
			continue
		}
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintf("Error reading standard input: %v", err))
	}

	return locks, keys
}

func isOverlap(lock, key []int) bool {
	for i := 0; i < len(lock); i++ {
		if key[i]+lock[i] >= 6 {
			return true
		}
	}
	return false
}

func solve(locks, keys [][]int) int {
	total := 0
	for _, lock := range locks {
		for _, key := range keys {
			if !isOverlap(lock, key) {
				total++
			}
		}
	}
	return total
}

func main() {
	locks, keys := input()
	fmt.Println(solve(locks, keys))
}
