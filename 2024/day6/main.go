package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Pos struct {
	x, y int
}

func input() ([]string, Pos) {
	scanner := bufio.NewScanner(os.Stdin)
	var guardPos Pos

	var input = []string{}
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			input = append(input, line)
			guardX := strings.Index(line, "^")
			if guardX >= 0 {
				guardPos = Pos{x: guardX, y: len(input) - 1}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintf("Error reading standard input: %v", err))
	}

	return input, guardPos
}

func solve(input []string, guardPos Pos) (int, bool, map[Pos]bool) {
	directions := []Pos{
		{x: 0, y: -1},
		{x: 1, y: 0},
		{x: 0, y: 1},
		{x: -1, y: 0},
	}
	dir := 0
	visitedPos := make(map[Pos]bool)
	visitedDirections := make(map[Pos]int)

	for {
		newPos := Pos{x: guardPos.x + directions[dir].x, y: guardPos.y + directions[dir].y}
		if newPos.x >= len(input[0]) || newPos.y >= len(input) || newPos.x < 0 || newPos.y < 0 {
			break
		}

		vd, ok := visitedDirections[newPos]
		if visitedPos[newPos] && vd == dir && ok {
			return len(visitedDirections), false, visitedPos
		}

		if input[newPos.y][newPos.x] == '#' {
			dir = (dir + 1) % 4
			continue
		} else {
			visitedDirections[newPos] = dir
			visitedPos[newPos] = true
			guardPos = newPos
		}
	}
	return len(visitedPos) + 1, true, visitedPos
}

func part1(input []string, guardPos Pos) int {
	total, _, _ := solve(input, guardPos)
	return total
}

func replaceAt(str string, c rune, pos int) string {
	return str[:pos] + string(c) + str[pos+1:]
}

func part2(input []string, guardPos Pos) int {
	total := 0
	_, _, visitedPos := solve(input, guardPos)

	for p := range visitedPos {
		if input[p.y][p.x] != '.' {
			continue
		}
		input2 := make([]string, len(input))
		copy(input2, input)
		input2[p.y] = replaceAt(input2[p.y], '#', p.x)
		_, ok, _ := solve(input2, guardPos)
		if !ok {
			total++
		}
	}
	return total
}

func main() {
	input, guardPos := input()
	fmt.Println(part1(input, guardPos))
	fmt.Println(part2(input, guardPos))
}
