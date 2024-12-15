package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pos struct {
	x, y int
}

type Robot struct {
	p, v Pos
}

func partToPos(s string) Pos {
	t := strings.Split(strings.Fields(s)[0][2:], ",")
	x, _ := strconv.Atoi(t[0])
	y, _ := strconv.Atoi(t[1])
	return Pos{x, y}
}

func input() []Robot {
	scanner := bufio.NewScanner(os.Stdin)

	var input = []Robot{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		r := Robot{
			p: partToPos(strings.Fields(line)[0]),
			v: partToPos(strings.Fields(line)[1]),
		}

		input = append(input, r)
	}

	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintf("Error reading standard input: %v", err))
	}

	return input
}

func draw(robots map[Pos]bool) {
	// cursorYX(0, 0)
	for y := 0; y < height; y += 1 {
		for x := 0; x < width; x += 1 {
			if robots[Pos{x, y}] {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
		fmt.Println()
	}
}

const width int = 101
const height int = 103

func move(r Robot, moves int) (Pos, Pos) {
	x := (r.p.x + (r.v.x * moves)) % width
	y := (r.p.y + (r.v.y * moves)) % height
	if x < 0 {
		x = width + x
	}
	if y < 0 {
		y = height + y
	}

	if x == width/2 || y == height/2 {
		return Pos{-1, -1}, Pos{x, y}
	}

	return Pos{(x) / (width/2 + 1), (y) / (height/2 + 1)}, Pos{x, y}
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func solve(robots []Robot) (int, int) {
	p2 := 0
	quads := make(map[Pos]int)
	for _, r := range robots {
		p, _ := move(r, 100)
		if p.x != -1 {
			quads[p]++
		}
	}
	p1 := 1
	for _, v := range quads {
		p1 *= max(v, 1)
	}

	for i := 0; ; i++ {
		insideBounds := 0
		for _, r := range robots {
			_, p := move(r, i)

			if abs(p.x-width/2) < p.y+1 {
				insideBounds++
			}
		}

		if float32(insideBounds)/float32(len(robots)) >= 0.92 {
			p2 = i
			// draw(rs)
			break
		}
	}

	return p1, p2
}

func main() {
	robots := input()
	p1, p2 := solve(robots)
	fmt.Println(p1)
	fmt.Println(p2)
}
