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

func sign(a int) int {
	if a > 0 {
		return 1
	} else if a < 0 {
		return -1
	}
	return 0
}

func distance(src, dst Pos) int {
	return abs(dst.x-src.x) + abs(dst.y-src.y)
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
		// Don't want to move right if possible
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

func sequence(line string, keyMap map[byte]Pos) string {
	prev := byte('A')
	result := []byte{}
	for _, c := range []byte(line) {
		pth := path(prev, c, keyMap)
		result = append(result, pth...)
		prev = c
	}
	return string(result)
}

func solve(input []string) int {
	total := 0

	for _, line := range input {
		seq := sequence(line, numKeys)
		seq1 := sequence(seq, dirKeys)
		seq2 := sequence(seq1, dirKeys)
		total += len(seq2) * getNum(line)
	}
	return total
}

func main() {
	fmt.Printf("Distance between A and v %d\n", distance(dirKeys['A'], dirKeys['v']))
	fmt.Printf("Distance between A and < %d\n", distance(dirKeys['A'], dirKeys['<']))

	input := input()
	// too high 159558
	fmt.Println(solve(input))
}
