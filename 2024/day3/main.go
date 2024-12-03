package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var reDigit = regexp.MustCompile(`\d+`)

func deque(queue [][]int) ([]int, [][]int) {
	if len(queue) == 0 {
		panic("No more items in queue")
	}

	n := queue[0]
	return n, queue[1:]
}

func peek(queue [][]int) []int {
	return queue[0]
}

func compute(mul string) int {
	digits := reDigit.FindAllString(mul, -1)
	n1, _ := strconv.Atoi(digits[0])
	n2, _ := strconv.Atoi(digits[1])
	return n1 * n2
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
func part1(input []string) int {
	total := 0
	reMul := regexp.MustCompile(`mul\(\d+,\d+\)`)
	for _, line := range input {
		for _, mul := range reMul.FindAllString(line, -1) {
			total += compute(mul)
		}
	}
	return total
}

func part2(input []string) int {
	total := 0
	reMul := regexp.MustCompile(`mul\(\d+,\d+\)`)
	reDo := regexp.MustCompile(`do\(\)`)
	reDont := regexp.MustCompile(`don't\(\)`)

	enabled := true
	for _, line := range input {
		dos := reDo.FindAllStringIndex(line, -1)
		donts := reDont.FindAllStringIndex(line, -1)
		muls := reMul.FindAllStringIndex(line, -1)
		var curMul []int

		for i := 0; i < len(line); i++ {
			if len(dos) > 0 && peek(dos)[0] == i {
				enabled = true
				_, dos = deque(dos)
			}
			if len(donts) > 0 && peek(donts)[0] == i {
				enabled = false
				_, donts = deque(donts)
			}
			if len(muls) > 0 && peek(muls)[0] == i {
				curMul, muls = deque(muls)
				if enabled {
					mul := line[curMul[0]:curMul[1]]
					total += compute(mul)
				}
			}
		}

	}
	return total
}

func main() {
	input := input()

	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
