package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Solved part 2 with help of https://observablehq.com/@jwolondon/advent-of-code-2024-day-21

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

func sign(a int) int {
	if a > 0 {
		return 1
	} else if a < 0 {
		return -1
	}
	return 0
}

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

func path(a byte, b byte, keyMap map[byte]Pos) []byte {
	pA := keyMap[a]
	pB := keyMap[b]
	distX := pB.x - pA.x
	distY := pB.y - pA.y
	xDir := map[int]byte{1: '>', -1: '<'}
	yDir := map[int]byte{1: 'v', -1: '^'}

	xpath := []byte{}
	ypath := []byte{}
	for i := 0; i < abs(distX); i++ {
		xpath = append(xpath, xDir[sign(distX)])
	}
	for i := 0; i < abs(distY); i++ {
		ypath = append(ypath, yDir[sign(distY)])
	}

	var path []byte
	if (Pos{pB.x, pA.y}) == keyMap['X'] {
		// move y first
		path = append(ypath, xpath...)
	} else if (Pos{pA.x, pB.y}) == keyMap['X'] {
		// move x first
		path = append(xpath, ypath...)
	} else {
		// moving right is wasteful
		if distX < 0 {
			path = append(xpath, ypath...)
		} else {
			path = append(ypath, xpath...)
		}
	}

	return append(path, 'A')
}

func getNum(line string) int {
	n, _ := strconv.Atoi(strings.Replace(line, "A", "", -1))
	return n
}

func sequence(line string, keyMap map[byte]Pos, cache map[string][]byte) map[string]int {
	prev := byte('A')
	dstCnt := map[string]int{}
	for _, c := range []byte(line) {
		cKey := string(prev) + string(c)
		pth, ok := cache[cKey]
		if !ok {
			pth = path(prev, c, keyMap)
			cache[cKey] = pth
		}
		dstCnt[string(pth)]++
		prev = c
	}
	return dstCnt
}

func solve(input []string, depth int) int {
	total := 0
	cache := map[string][]byte{}

	for _, line := range input {
		cnt := sequence(line, numKeys, map[string][]byte{})

		for i := 0; i < depth; i++ {
			stepCnt := map[string]int{}
			for pth, pthCount := range cnt {
				newCnt := sequence(pth, dirKeys, cache)
				for k, newCount := range newCnt {
					stepCnt[k] += newCount * pthCount
				}
			}
			cnt = stepCnt
		}

		totalLen := 0
		for pth, v := range cnt {
			totalLen += v * len(pth)
		}

		total += totalLen * getNum(line)
	}
	return total
}

func main() {
	input := input()
	fmt.Println(solve(input, 2))
	fmt.Println(solve(input, 25))
}
