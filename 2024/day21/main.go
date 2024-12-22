package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

type Pos struct {
	x, y int
}

var numKeys = map[byte]Pos{
	'7': {0, 0},
	'8': {1, 0},
	'9': {2, 0},
	'4': {0, 1},
	'5': {1, 1},
	'6': {2, 1},
	'1': {0, 2},
	'2': {1, 2},
	'3': {2, 2},
	'0': {1, 3},
	'A': {2, 3},
	// forbidden
	'X': {0, 3},
}

var dirKeys = map[byte]Pos{
	'^': {1, 0},
	'A': {2, 0},
	'<': {0, 1},
	'v': {1, 1},
	'>': {2, 1},
	// forbidden
	'X': {0, 0},
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func path(a byte, b byte, keyMap map[byte]Pos) []byte {
	pA := keyMap[a]
	pB := keyMap[b]
	x := pB.x - pA.x
	y := pB.y - pA.y
	path := []byte{}
	px := Pos{pA.x + x, pA.y}
	py := Pos{pA.x, pA.y + y}
	if px == keyMap['X'] {
		// move y first
		for i := 0; i < abs(y); i++ {
			if y > 0 {
				path = append(path, 'v')
			} else {
				path = append(path, '^')
			}
		}
		for i := 0; i < abs(x); i++ {
			if x > 0 {
				path = append(path, '>')
			} else {
				path = append(path, '<')
			}
		}
	} else if py == keyMap['X'] {
		// move x first
		for i := 0; i < abs(x); i++ {
			if x > 0 {
				path = append(path, '>')
			} else {
				path = append(path, '<')
			}
		}
		for i := 0; i < abs(y); i++ {
			if y > 0 {
				path = append(path, 'v')
			} else {
				path = append(path, '^')
			}
		}
	} else {
		for i := 0; i < abs(x); i++ {
			if x > 0 {
				path = append(path, '>')
			} else {
				path = append(path, '<')
			}
		}
		for i := 0; i < abs(y); i++ {
			if y > 0 {
				path = append(path, 'v')
			} else {
				path = append(path, '^')
			}
		}
	}
	return append(path, 'A')
}

func getNum(line string) int {
	n, _ := strconv.Atoi(strings.Replace(line, "A", "", -1))
	return n
}

func solve(input []string) int {
	total := 0
	prev := byte('A')
	prev1 := byte('A')
	prev2 := byte('A')

	for _, line := range input {
		// fmt.Print(line, ": ")
		// prev3 := byte('A')
		result := []byte{}
		result1 := []byte{}
		result2 := []byte{}
		for _, c := range []byte(line) {
			pth := path(prev, c, numKeys)
			result1 = append(result1, pth...)
			prev = c

			for _, d := range pth {
				pth2 := path(prev1, d, dirKeys)
				result2 = append(result2, pth2...)
				prev1 = d

				for _, e := range pth2 {
					pth3 := path(prev2, e, dirKeys)
					// result = append(result, []byte(fmt.Sprintf("%c-%c:", prev2, e))...)
					result = append(result, pth3...)
					// result = append(result, '|')
					prev2 = e

					// for _, f := range pth3 {
					// 	pth4 := path(prev3, e, dirKeys)
					// 	prev3 = f
					// }
				}
			}

		}
		// 159558 too high
		fmt.Println("line:", line)
		fmt.Printf("len(result) * getNum(line)=%d*%d=%d\n\n", len(result), getNum(line), len(result)*getNum(line))
		total += len(result) * getNum(line)
		// fmt.Println(string(result))
		// fmt.Println(string(result2))
		// fmt.Println(string(result1))
	}
	return total
}

func main() {
	// fmt.Println(string(path('A', '2', numKeys)))

	input := input()
	// too high 159558
	fmt.Println(solve(input))
}
