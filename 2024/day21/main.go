package main

import (
	"bufio"
	"fmt"
	"math"
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

func sign(n int) int {
	if n == 0 {
		return 0
	} else if n > 0 {
		return 1
	} else {
		return -1
	}
}

func direction(src, dst Pos) Pos {
	return Pos{sign(dst.x - src.x), sign(dst.y - src.y)}
}

func (a Pos) move(b Pos) Pos {
	return Pos{a.x + b.x, a.y + b.y}
}

func getKeyMap(robotNo int) map[byte]Pos {
	if robotNo == 0 {
		return numKeys
	} else {
		return dirKeys
	}
}

var WrongBytes []byte = []byte{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}

func getNum(line string) int {
	n, _ := strconv.Atoi(strings.Replace(line, "A", "", -1))
	return n
}

func pathSS(p, dst, dir Pos, keyMap map[byte]Pos, path []byte, checkNext bool) []byte {
	if p == dst {
		return append(path, 'A')
	}

	maxY := 2
	if len(keyMap) == len(numKeys) {
		maxY = 3
	}

	// Forbidden
	if keyMap['X'] == p || p.x < 0 || p.x > 2 || p.y < 0 || p.y > maxY {
		return WrongBytes
	}

	px := Pos{p.x + dir.x, p.y}
	py := Pos{p.x, p.y + dir.y}

	pathx := append([]byte{}, path...)
	pathy := append([]byte{}, path...)
	if dir.x > 0 {
		pathx = append(pathx, '>')
	} else if dir.x < 0 {
		pathx = append(pathx, '<')
	}

	if dir.y > 0 {
		pathy = append(pathy, 'v')
	} else if dir.y < 0 {
		pathy = append(pathy, '^')
	}

	xbytes := WrongBytes
	ybytes := WrongBytes
	lenNextX := math.MaxInt
	lenNextY := math.MaxInt

	// fmt.Println("p:", p, "dst:", dst, "dir:", dir)
	if dir.x != 0 {
		xbytes = pathSS(px, dst, dir, keyMap, pathx, checkNext)
		checkNext = checkNext && len(xbytes) != len(WrongBytes)
		if checkNext {
			lenNextX = ss(string(xbytes), 2, false, dirKeys)
		}
	}
	if dir.y != 0 {
		ybytes = pathSS(py, dst, dir, keyMap, pathy, checkNext)
		checkNext = checkNext && len(ybytes) != len(WrongBytes)
		if checkNext {
			lenNextY = ss(string(ybytes), 2, false, dirKeys)
		}
	}

	if len(xbytes) == len(WrongBytes) && len(ybytes) == len(WrongBytes) {
		return WrongBytes
	}

	if checkNext {
		if lenNextX < lenNextY {
			return xbytes
		} else {
			return ybytes
		}
	}

	if len(xbytes) < len(ybytes) {
		return xbytes

	}
	return ybytes
}

func ss(origLine string, depth int, checkNext bool, firstKeyMap map[byte]Pos) int {
	var keyMap map[byte]Pos

	moves := append([]byte{}, origLine...)

	for i := 0; i < depth; i++ {
		if i == 0 {
			keyMap = firstKeyMap
		} else {
			keyMap = dirKeys
		}

		newMoves := []byte{}

		var prev byte = 'A'
		for _, c := range []byte(moves) {
			dir := direction(keyMap[prev], keyMap[c])
			pth := pathSS(keyMap[prev], keyMap[c], dir, keyMap, []byte{}, checkNext)
			newMoves = append(newMoves, pth...)
			prev = c
		}
		moves = []byte{}
		moves = append(moves, newMoves...)
		if checkNext {
			fmt.Printf("%d. %v\n", i+1, string(moves))
		}
	}
	// fmt.Printf("Result for %s: %s len: %d\n", origLine, string(moves), len(moves))
	return len(moves)
}

func solveSS(input []string) int {
	total := 0

	for _, line := range input {
		fmt.Printf("%s:\n", line)
		res := ss(line, 3, true, numKeys)
		fmt.Printf("%s: %d*%d=%d\n", line, res, getNum(line), res*getNum(line))
		total += res * getNum(line)
		// fmt.Println("line:", line)
		// fmt.Printf("len(result) * getNum(line)=%d*%d=%d\n\n", len(result), getNum(line), len(result)*getNum(line))
		// fmt.Println(string(result))
	}
	return total
}

func main() {
	// fmt.Println(string(path('A', '2', numKeys)))

	input := input()
	// too high 159558
	fmt.Println(solveSS(input))
}
