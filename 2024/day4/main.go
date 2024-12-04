package main

import (
	"bufio"
	"fmt"
	"os"
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

func score(words ...string) int {
	s := 0
	for _, w := range words {
		if w == "XMAS" || w == "SAMX" {
			s++
		}
	}
	return s
}

func part1(input []string) int {
	total := 0
	for y := 0; y < len(input); y++ {
		line := input[y]
		for x := 0; x < len(line); x++ {
			hor := line[x:min(len(line), x+4)]
			ver := ""
			diagR := ""
			diagL := ""
			for yy := y; yy < min(len(input), y+4); yy++ {
				ver += string(input[yy][x])
			}
			for pos := 0; pos < 4; pos++ {
				if y+pos >= len(input) {
					break
				}
				if x+pos < len(line) {
					diagR += string(input[y+pos][x+pos])
				}
				if x-pos >= 0 {
					diagL += string(input[y+pos][x-pos])
				}
			}
			total += score(hor, ver, diagR, diagL)
		}
	}
	return total
}

func main() {
	input := input()
	fmt.Println(part1(input))
}
