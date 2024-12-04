package main

import (
	"bufio"
	"bytes"
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
			var ver bytes.Buffer
			var diagR bytes.Buffer
			var diagL bytes.Buffer
			for yy := y; yy < min(len(input), y+4); yy++ {
				ver.WriteByte(input[yy][x])
			}
			for pos := 0; pos < 4; pos++ {
				if y+pos >= len(input) {
					break
				}
				if x+pos < len(line) {
					diagR.WriteByte(input[y+pos][x+pos])
				}
				if x-pos >= 0 {
					diagL.WriteByte(input[y+pos][x-pos])
				}
			}

			if y < len(input)-2 && len(hor) >= 3 && diagR.Len() >= 3 {
				diag1 := diagR.String()[:3]
				diag2 := string([]byte{input[y][x+2], input[y+1][x+1], input[y+2][x]})
				if isMas(diag1) && isMas(diag2) {
					total2++
				}
			}

			total1 += xmasScore(hor, ver.String(), diagR.String(), diagL.String())
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
