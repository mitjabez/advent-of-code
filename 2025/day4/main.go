package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	x, y int
}

func input() []string {
	scanner := bufio.NewScanner(os.Stdin)

	var input = []string{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		input = append(input, line)
	}

	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintf("Error reading standard input: %v", err))
	}

	return input
}

func solveRound(input []string, removedRolls map[Pos]bool) (int, map[Pos]bool) {
	total := 0

	for y, row := range input {
		for x, c := range row {
			if c != '@' || removedRolls[Pos{x, y}] {
				continue
			}

			nRolls := 0
			for _, dy := range []int{-1, 0, 1} {
				for _, dx := range []int{-1, 0, 1} {
					if x+dx < 0 || x+dx >= len(row) ||
						y+dy < 0 || y+dy >= len(input) ||
						(x == dx && y == dy) {
						continue
					}
					if input[y+dy][x+dx] == '@' && !removedRolls[Pos{x + dx, y + dy}] {
						nRolls++
					}
				}
			}
			if nRolls <= 4 {
				removedRolls[Pos{x, y}] = true
				total++
			}
		}
	}

	return total, removedRolls
}

func solve(input []string) (int, int) {
	p1, _ := solveRound(input, make(map[Pos]bool))

	p2 := 0
	removedRolls := map[Pos]bool{}
	round := -1
	for round != 0 {
		round, removedRolls = solveRound(input, removedRolls)
		p2 += round
	}
	return p1, p2
}

func main() {
	input := input()
	p1, p2 := solve(input)
	fmt.Println(p1)
	fmt.Println(p2)
}
