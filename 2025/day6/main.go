package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Problem struct {
	numbers  []int
	operator string
}

func input() ([]Problem, []Problem) {
	scanner := bufio.NewScanner(os.Stdin)

	problems1 := []Problem{}
	grid := []string{}

	// Part1 parse
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		problems := strings.Fields(line)
		grid = append(grid, line)

		// Part1
		if len(problems1) == 0 {
			problems1 = make([]Problem, len(problems))
		}

		for i, strProblem := range problems {
			if strProblem == "" {
				continue
			}
			if strProblem != "+" && strProblem != "*" {
				num, _ := strconv.Atoi(strProblem)
				problems1[i].numbers = append(problems1[i].numbers, num)
			} else {
				problems1[i].operator = problems[i]
			}
		}
	}

	// Part2 parse
	// "A man has to do what a man has to do ..."
	problems2 := make([]Problem, 1)
	problemId := 0
	for x := 0; x < len(grid[0]); x++ {
		strNum := ""
		hasNumber := false
		for y := 0; y < len(grid); y++ {
			c := grid[y][x]
			if c >= '0' && c <= '9' {
				hasNumber = true
				strNum += string(c)
			} else if c == '+' || c == '*' {
				problems2[problemId].operator = string(c)
			}
		}
		if hasNumber {
			num, _ := strconv.Atoi(strNum)
			problems2[problemId].numbers = append(problems2[problemId].numbers, num)
		} else {
			problemId++
			problems2 = append(problems2, Problem{})
		}
	}

	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintf("Error reading standard input: %v", err))
	}

	return problems1, problems2
}

func solve(problems []Problem) int {
	p1 := 0

	for _, p := range problems {
		total := 0
		for j, num := range p.numbers {
			if j == 0 {
				total = num
			} else {
				if p.operator == "+" {
					total += num
				} else {
					total *= num
				}
			}
		}
		p1 += total
	}

	return p1
}

func main() {
	inputPart1, inputPart2 := input()
	p1 := solve(inputPart1)
	p2 := solve(inputPart2)
	fmt.Println(p1)
	fmt.Println(p2)
}
