package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func input() ([]string, []string) {
	scanner := bufio.NewScanner(os.Stdin)

	isTowels := true
	towels := []string{}
	designs := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			isTowels = false
			continue
		}

		if isTowels {
			towels = strings.Split(line, ", ")
			continue
		}

		designs = append(designs, line)
	}

	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintf("Error reading standard input: %v", err))
	}

	return towels, designs
}

func ss(line string, leftover string, cache map[string]int, towels []string) int {
	if line == "" {
		return 1
	}

	var n int = 0
	for _, t := range towels {
		if strings.HasPrefix(line, t) {
			nCache, hasCache := cache[leftover+line[0:len(t)]]
			if hasCache {
				n += nCache
			} else {
				sum := ss(line[len(t):], leftover+t, cache, towels)
				n += sum
				cache[leftover+line[0:len(t)]] = sum
			}
		}
	}

	return n
}

func solve(towels, designs []string) (int, int) {
	towelMap := make(map[string]bool)
	for _, t := range towels {
		towelMap[t] = true
	}

	p1 := 0
	p2 := 0
	for _, d := range designs {
		ways := ss(d, "", make(map[string]int), towels)
		if ways > 0 {
			p1++
		}
		p2 += ways
	}
	return p1, p2
}

func main() {
	towels, designs := input()
	p1, p2 := solve(towels, designs)
	fmt.Println(p1)
	fmt.Println(p2)
}
