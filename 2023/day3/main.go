package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Result struct {
	Part1 int
	Part2 int
}

func clamp(n int, maxn int, minn int) int {
	return min(max(n, minn), maxn)
}

func isDigit(c byte) bool {
	return c >= '0' && c <= '9'
}

func isSymbol(c byte) bool {
	return c != '.' && !isDigit(c)
}

func extractNumber(line string, xstart int) (int, int) {
	digit := ""
	nums := []int{0, 0}
	n := 0
	lastDigit := false
	for x := xstart; x < min(xstart+3, len(line)); x++ {
		if isDigit(line[x]) && !lastDigit {
			lastDigit = true
			xx := x
			// Find first digit
			for {
				xx--
				if xx == -1 || !isDigit(line[xx]) {
					xx++
					break
				}

			}

			for {
				if xx > len(line)-1 || !isDigit(line[xx]) {
					break
				}

				digit += string(line[xx])
				xx++
			}
			nums[n], _ = strconv.Atoi(digit)
			n++
		}

		if !isDigit(line[x]) {
			lastDigit = false
		}
	}

	return nums[0], nums[1]
}

func solve(input []string) (int, int) {
	part1 := 0
	part2 := 0

	for y, line := range input {
		for x, c := range []byte(line) {
			if isSymbol(c) {
				isGear := c == '*'
				gears := []int{}

				if y > 0 {
					// below
					num1, num2 := extractNumber(input[y-1], x-1)
					part1 += num1 + num2
					// if isGear && num > 0 {
					// 	gears = append(gears, num)
					// }
				}
				if y < len(input)-1 {
					// above
					num1, num2 := extractNumber(input[y+1], x-1)
					part1 += num1 + num2
					// if isGear && num > 0 {
					// 	gears = append(gears, num)
					// }
				}
				if x > 0 && isDigit(input[y][x-1]) {
					num, _ := extractNumber(input[y], x-1)
					part1 += num
					if isGear && num > 0 {
						gears = append(gears, num)
					}
				}
				if x < len(line)-2 && isDigit(input[y][x+1]) {
					num, _ := extractNumber(input[y], x+1)
					part1 += num
					if isGear && num > 0 {
						gears = append(gears, num)
					}
				}

				if len(gears) == 2 {
					part2 += gears[0] * gears[1]
				}
			}
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
