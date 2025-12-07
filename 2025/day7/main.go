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

func solve(input []string) (int, int) {
	p1 := 0
	p2 := 0

	timelinesPerPos := map[Pos]int{}
	for y, line := range input {
		for x, c := range line {
			switch c {
			case 'S':
				timelinesPerPos[Pos{x, y}] = 1
			case '.':
				if timelinesPerPos[Pos{x, y - 1}] > 0 {
					timelinesPerPos[Pos{x, y}] += timelinesPerPos[Pos{x, y - 1}]
				}
			case '^':
				if timelinesPerPos[Pos{x, y - 1}] > 0 {
					p1++
					timelinesPerPos[Pos{x - 1, y}] += timelinesPerPos[Pos{x, y - 1}]
					timelinesPerPos[Pos{x + 1, y}] += timelinesPerPos[Pos{x, y - 1}]
				}
			}
		}
	}

	for pos, count := range timelinesPerPos {
		if pos.y == len(input)-1 {
			p2 += count
		}
	}

	return p1, p2
}

func main() {
	input := input()
	p1, p2 := solve(input)
	fmt.Println(p1)
	fmt.Println(p2)
}
