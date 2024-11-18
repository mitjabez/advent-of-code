package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isSymbol(c byte) bool {
	return c != '.' && !isDigit(c)
}

type Pos struct {
	x int
	y int
}

func solve(input []string) (int, int) {
	part1 := 0
	part2 := 0
	gears := map[Pos][]int{}

	for y, line := range input {
		for x := 0; x < len(line); x++ {
			digit := ""
			hasAdj := false
			hasStar := false
			var starPos Pos
			if isDigit(line[x]) {
				for {
					digit += string(line[x])

					for yy := max(0, y-1); yy < min(y+2, len(input)); yy++ {
						for xx := max(0, x-1); xx < min(x+2, len(line)); xx++ {
							if yy == y && xx == x {
								continue
							}
							c := input[yy][xx]
							if isSymbol(c) {
								hasAdj = true
								if c == '*' {
									hasStar = true
									starPos = Pos{x: xx, y: yy}
								}
							}
						}
					}

					x++

					if x > len(line)-1 || !isDigit(line[x]) {
						break
					}
				}

				if hasAdj {
					numDigit, _ := strconv.Atoi(digit)
					part1 += numDigit

					if hasStar {
						gears[starPos] = append(gears[starPos], numDigit)
					}
				}
			}
		}
	}

	for _, gearsForPos := range gears {
		if len(gearsForPos) == 2 {
			part2 += gearsForPos[0] * gearsForPos[1]
		}
	}

	return part1, part2
}

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

func main() {
	input := input()

	part1, part2 := solve(input)
	fmt.Println("Part1:", part1)
	fmt.Println("Part2:", part2)
}
