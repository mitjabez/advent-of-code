package main

import (
	"bufio"
	"fmt"
	"os"
)

func input() [][]int {
	scanner := bufio.NewScanner(os.Stdin)

	var input = [][]int{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		row := []int{}
		for _, c := range line {
			if c != '.' {
				row = append(row, int(c)-'0')
			} else {
				row = append(row, 100)
			}
		}
		input = append(input, row)
	}

	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintf("Error reading standard input: %v", err))
	}

	return input
}

func path(maze [][]int, visited map[int]map[int]bool, x, y int, prev int, p2 bool) int {
	if x < 0 || x >= len(maze[0]) || y < 0 || y >= len(maze) {
		return 0
	}
	c := maze[y][x]
	if c-prev != 1 {
		return 0
	}
	if c == 9 {
		if !p2 {
			if visited[x] == nil {
				visited[x] = make(map[int]bool)
			}
			if visited[x][y] {
				return 0
			}
		}
		visited[x][y] = true
		return 1
	}
	return path(maze, visited, x+1, y, c, p2) + path(maze, visited, x-1, y, c, p2) + path(maze, visited, x, y+1, c, p2) + path(maze, visited, x, y-1, c, p2)
}

func solve(maze [][]int) (int, int) {
	p1 := 0
	p2 := 0
	for y, row := range maze {
		for x, c := range row {
			if c == 0 {
				visited := make(map[int]map[int]bool)
				p1 += path(maze, visited, x, y, -1, false)
				p2 += path(maze, visited, x, y, -1, true)
			}
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
