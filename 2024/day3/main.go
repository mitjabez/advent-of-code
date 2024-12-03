package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var reDigit = regexp.MustCompile(`\d+`)

func deque(queue [][]int) []int {
	if len(queue) == 0 {
		panic("No more items in queue")
	}

	n := queue[0]
	queue = queue[1:]
	return n
}

func peek(queue [][]int) []int {
	if len(queue) == 0 {
		return nil
	}

	return queue[0]
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

func compute(mul string) int {
	digits := reDigit.FindAllString(mul, -1)
	n1, _ := strconv.Atoi(digits[0])
	n2, _ := strconv.Atoi(digits[1])
	return n1 * n2
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

func isEnabled(mulStart int, dos [][]int, donts [][]int) bool {
	doPos := -1
	for i, do := range dos {
		if do[0] > mulStart {
			doPos = i - 1
			break
		}
	}
	dontPos := -1
	for i, dont := range donts {
		if dont[0] > mulStart {
			dontPos = i - 1
			break
		}
	}

	if dontPos < 0 {
		return true
	}

	if doPos < 0 {
		return false
	}

	return dos[doPos][0] >= donts[dontPos][0]
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
		doi := 0
		donti := 0
		muli := 0

		for i := 0; i < len(line); i++ {
			if doi < len(dos) && dos[doi][0] == i {
				enabled = true
				doi++
			}
			if donti < len(donts) && donts[donti][0] == i {
				enabled = false
				donti++
			}
			if muli < len(muls) && muls[muli][0] == i {
				if enabled {
					mul := line[muls[muli][0]:muls[muli][1]]
					total += compute(mul)
				}
				muli++
			}
		}

	}
	return total
}

func main() {
	input := input()

	// fmt.Println(part1(input))
	fmt.Println(part2(input))
}
