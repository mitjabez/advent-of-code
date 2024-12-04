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

func xmasScore(words ...string) int {
	s := 0
	for _, w := range words {
		if w == "XMAS" || w == "SAMX" {
			s++
		}
	}
	return s
}

func isMas(w string) bool {
	return w == "MAS" || w == "SAM"
}

func solve(input []string) (int, int) {
	total1 := 0
	total2 := 0
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

			if y < len(input)-2 && len(hor) >= 3 && len(diagR) >= 3 {
				diag1 := string(input[y][x]) + string(input[y+1][x+1]) + string(input[y+2][x+2])
				diag2 := string(input[y][x+2]) + string(input[y+1][x+1]) + string(input[y+2][x])
				if isMas(diag1) && isMas(diag2) {
					total2++
				}
			}

			total1 += xmasScore(hor, ver, diagR, diagL)
		}
	}
	return total1, total2
}

func main() {
	input := input()
	part1, part2 := solve(input)
	fmt.Println(part1)
	fmt.Println(part2)
}
