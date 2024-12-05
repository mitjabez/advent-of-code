package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func input() (map[int]map[int]bool, [][]int) {
	scanner := bufio.NewScanner(os.Stdin)

	rules := make(map[int]map[int]bool)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		rt := strings.Split(line, "|")
		n1, _ := strconv.Atoi(rt[0])
		n2, _ := strconv.Atoi(rt[1])
		if rules[n1] == nil {
			rules[n1] = make(map[int]bool)
		}
		rules[n1][n2] = true
	}

	updates := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		strUpdates := strings.Split(line, ",")
		update := []int{}
		for _, ut := range strUpdates {
			n, _ := strconv.Atoi(ut)
			update = append(update, n)
		}
		updates = append(updates, update)

	}
	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintf("Error reading standard input: %v", err))
	}

	return rules, updates
}

func isValid(rules map[int]map[int]bool, update []int) bool {
	for i, u := range update {
		for j := i + 1; j < len(update); j++ {
			if !rules[u][update[j]] {
				return false
			}
		}
	}
	return true
}

func fix(rules map[int]map[int]bool, update []int) []int {
	slices.SortFunc(update, func(a, b int) int {
		for v := range rules[b] {
			if v == a {
				return 1
			}
		}
		return -1
	})

	return update
}

func solve(rules map[int]map[int]bool, updates [][]int) (int, int) {
	total1 := 0
	total2 := 0

	for _, update := range updates {
		ok := isValid(rules, update)
		if !ok {
			fixed := fix(rules, update)
			total2 += fixed[len(fixed)/2]
		} else {
			total1 += update[len(update)/2]
		}
	}
	return total1, total2
}

func main() {
	rules, updates := input()
	total1, total2 := solve(rules, updates)
	fmt.Println(total1)
	fmt.Println(total2)
}
