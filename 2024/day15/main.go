package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Pos struct {
	x, y int
}

// Initial position, boxes, moves
func input() ([][]byte, []byte, Pos) {
	scanner := bufio.NewScanner(os.Stdin)

	maze := [][]byte{}
	moves := []byte{}
	isMaze := true
	var start Pos
	y := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			isMaze = false
			continue
		}
		if isMaze {
			x := strings.Index(line, "@")
			if x > -1 {
				start = Pos{x, y}
			}
			maze = append(maze, []byte(line))
			y++
			continue
		}

		moves = append(moves, []byte(line)...)
	}

	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintf("Error reading standard input: %v", err))
	}

	return maze, moves, start
}

func popPos(queue []Pos) (Pos, []Pos) {
	if len(queue) == 0 {
		panic("No more items in queue")
	}

	p := queue[len(queue)-1]
	return p, queue[0 : len(queue)-1]
}
func popByte(queue []byte) (byte, []byte) {
	if len(queue) == 0 {
		panic("No more items in queue")
	}

	p := queue[len(queue)-1]
	return p, queue[0 : len(queue)-1]
}

func direction(c byte) Pos {
	if c == '>' {
		return Pos{1, 0}
	} else if c == '<' {
		return Pos{-1, 0}
	} else if c == '^' {
		return Pos{0, -1}
	} else if c == 'v' {
		return Pos{0, 1}
	}
	panic("Unknown direction")
}

func draw(maze [][]byte) {
	for y := range maze {
		for x := range maze[y] {
			fmt.Print(string(maze[y][x]))
		}
		fmt.Println()
	}
}

func part1(maze [][]byte, moves []byte, start Pos) int {
	p := start
	for _, d := range moves {
		dir := direction(d)

		pn := p
		qc := []byte{}
		qp := []Pos{}

		for {
			qc = append(qc, maze[pn.y][pn.x])
			pn = Pos{pn.x + dir.x, pn.y + dir.y}
			qp = append(qp, pn)

			c := maze[pn.y][pn.x]
			// We cannot move anything
			if c == '#' {
				break
			}

			// Can move through dot
			if c == '.' {
				for {
					c, qc = popByte(qc)
					pn, qp = popPos(qp)
					maze[pn.y][pn.x] = c
					if len(qc) == 0 {
						break
					}
				}

				maze[p.y][p.x] = '.'
				p = pn
				break
			} else if c == 'O' {
				// qc = append(qc, c)
			} else {
				panic("Unknown char " + string(c))
			}
		}

	}

	total := 0
	for y := range maze {
		for x, c := range maze[y] {
			if c == 'O' {
				total += y*100 + x
			}
		}
	}

	return total
}

func getBox(p Pos, boxes map[Pos]bool) (Pos, bool) {
	if boxes[p] {
		return p, true
	} else if boxes[Pos{p.x - 1, p.y}] {
		return Pos{p.x - 1, p.y}, true
	} else {
		return Pos{0, 0}, false
	}
}

func isWall(p Pos, walls map[Pos]bool) bool {
	return walls[p]
}

func draw2(current Pos, walls map[Pos]bool, boxes map[Pos]bool, size Pos) {
	fmt.Print("  ")
	for x := range size.x {
		sx := strconv.Itoa(x)
		fmt.Print(sx[0:1])
	}
	fmt.Println()
	for y := 0; y < size.y; y++ {
		fmt.Print(y, " ")
		for x := 0; x < size.x; x++ {
			p := Pos{x, y}
			if p == current {
				fmt.Print("@")
			} else if walls[p] {
				fmt.Print("#")
			} else if boxes[p] {
				fmt.Print("[]")
				x++
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Print("  ")
	for x := range size.x {
		sx := strconv.Itoa(x)
		fmt.Print(sx[0:1])
	}
	fmt.Println()
}

func part2(maze Maze) int {
	p := maze.start
	boxes := maze.boxes
	draw2(p, maze.walls, boxes, maze.size)
	for i, d := range maze.moves {
		dir := direction(d)
		// fmt.Println(dir)
		pn := p
		fmt.Printf("%d. %s\n", i+1, string(d))
		movedBoxes := make(map[Pos]bool)
		endMove := false

		bounds := []int{p.x + dir.x}
		// newBounds := []int{pn.x, pn.x}

		for {
			pn = Pos{pn.x + dir.x, pn.y + dir.y}
			if dir.y != 0 {
				// bounds[0] = min(bounds[0], pn.x)
				// bounds[1] = max(bounds[1], pn.x)
			} else {
				// bounds[0] = pn.x
				// bounds[1] = pn.x
				bounds = []int{pn.x}
			}

			newBounds := []int{}
			// copy(newBounds, bounds)
			isDot := true

			mb := make(map[int]bool)
			for _, x := range bounds {
				mb[x] = true
			}
			fmt.Println("B:", bounds, "mb:", mb, "y:", pn.y)
			fmt.Println(movedBoxes)
			for x, _ := range mb {
				px := Pos{x, pn.y}
				box, hasBox := getBox(px, boxes)
				if isWall(px, maze.walls) {
					endMove = true
					isDot = false
					break
				} else if hasBox {
					isDot = false
					movedBoxes[box] = true
					newBounds = append(newBounds, box.x)
					newBounds = append(newBounds, box.x+1)
					// newBounds[0] = min(box.x, newBounds[0])
					// newBounds[1] = max(box.x+1, newBounds[1])
				}
			}

			if endMove {
				break
			}

			// all dots in line so we can move
			if isDot {
				for box := range movedBoxes {
					delete(boxes, box)
				}
				for box := range movedBoxes {
					boxes[Pos{box.x + dir.x, box.y + dir.y}] = true
				}
				p = Pos{p.x + dir.x, p.y + dir.y}
				break
			}

			// fmt.Println("Copy bounds", newBounds)
			bounds = []int{}
			// clear(bounds)
			for _, b := range newBounds {
				bounds = append(bounds, b)
			}
			// copy(bounds, newBounds)
		}

		// draw2(p, maze.walls, boxes, maze.size)
		time.Sleep(time.Millisecond * 0)
	}

	// too low   1416239
	// too high  1469871
	// not right 1457765
	// not right 1457766
	// no right  1454116
	// not right 1439968
	total := 0
	for box := range boxes {
		total += box.y*100 + box.x
	}

	return total
}

func doubleMaze(maze [][]byte) ([][]byte, Pos) {
	m2 := [][]byte{}
	var s Pos
	for y := range maze {
		row := []byte{}
		for x, c := range maze[y] {
			if c == '@' {
				s = Pos{x * 2, y}
				row = append(row, '@', '.')
			} else if c == 'O' {
				row = append(row, '[', ']')
			} else {
				row = append(row, c, c)
			}
		}
		m2 = append(m2, row)
	}
	return m2, s
}

type Maze struct {
	start Pos
	walls map[Pos]bool
	boxes map[Pos]bool
	moves []byte
	size  Pos
}

func getMappedMaze(rawMaze [][]byte, moves []byte, start Pos) Maze {
	walls := make(map[Pos]bool)
	boxes := make(map[Pos]bool)

	for y := range rawMaze {
		for x, c := range rawMaze[y] {
			if c == 'O' {
				boxes[Pos{x * 2, y}] = true
			} else if c == '#' {
				walls[Pos{x * 2, y}] = true
				walls[Pos{x*2 + 1, y}] = true
			}
		}
	}
	return Maze{
		start: Pos{start.x * 2, start.y},
		walls: walls,
		boxes: boxes,
		moves: moves,
		size:  Pos{len(rawMaze[0]) * 2, len(rawMaze)},
	}
}

func main() {
	rawMaze, moves, start := input()
	maze := getMappedMaze(rawMaze, moves, start)
	fmt.Println(part2(maze))

	// fmt.Println(solve(maze, moves, start))
	// fmt.Println(part1(doubleMaze, moves, doubleStart))
	// fmt.Println(p2)
}
