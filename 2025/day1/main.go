package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func abs(n int) int {
	if n < 0 {
		return n * -1
	}
	return n
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

func sign(n int) int {
	if n == 0 {
		return 0
	} else if n > 0 {
		return 1
	}
	return -1
}

func solve(input []string) (int, int) {
	p1 := 0
	p2 := 0
	pos := 50
	for _, line := range input {
		direction := line[0]
		num, err := strconv.Atoi(line[1:])
		if err != nil {
			log.Fatalf("Cannot parse rotation: '%s'\n", line)
		}

		prevSign := sign(pos)
		prevPos := pos
		switch direction {
		case 'L':
			pos -= num
		case 'R':
			pos += num
		}
		newSign := sign(pos)
		if prevPos != 0 && prevSign != newSign {
			p2++
		}
		p2 += abs(pos) / 100
		pos %= 100

		if pos == 0 {
			p1++
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
