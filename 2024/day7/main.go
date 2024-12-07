package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func input() [][]int {
	scanner := bufio.NewScanner(os.Stdin)

	equations := [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			eq := []int{}
			for _, s := range strings.Fields(strings.Replace(line, ":", "", 1)) {
				n, _ := strconv.Atoi(s)
				eq = append(eq, n)
			}
			equations = append(equations, eq)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintf("Error reading standard input: %v", err))
	}

	return equations
}

// Go through all operator permutations of fixed size
func isValid(allOps []byte, previousOp []byte, opPos int, valPos int, r int, numbers []int, expectedResult int) bool {
	if opPos >= len(previousOp) || valPos >= len(allOps) {
		return false
	}

	currentOp := make([]byte, len(previousOp))
	copy(currentOp, previousOp)
	currentOp[opPos] = allOps[valPos]

	if opPos == r-1 {
		total := numbers[0]
		for i := 1; i < len(numbers); i++ {
			op := currentOp[i-1]
			if op == '+' {
				total = total + numbers[i]
			} else if op == '*' {
				total = total * numbers[i]
			} else if op == '|' {
				total, _ = strconv.Atoi(strconv.Itoa(total) + strconv.Itoa(numbers[i]))
			} else {
				panic("Unexpected operator: " + string(op))
			}
		}
		if total == expectedResult {
			return true
		}
	}

	return isValid(allOps, currentOp, opPos+1, 0, r, numbers, expectedResult) || isValid(allOps, currentOp, opPos, valPos+1, r, numbers, expectedResult)
}

func solve(equations [][]int) (int, int) {
	p1 := 0
	p2 := 0
	for _, eq := range equations {
		expected := eq[0]
		equation := eq[1:]
		if isValid([]byte{'+', '*'}, make([]byte, len(equation)-1), 0, 0, len(equation)-1, equation, expected) {
			p1 += expected
		}
		if isValid([]byte{'+', '*', '|'}, make([]byte, len(equation)-1), 0, 0, len(equation)-1, equation, expected) {
			p2 += expected
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
