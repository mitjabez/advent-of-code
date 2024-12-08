package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct{ y, x int }

func input() (map[rune][]Pos, Pos) {
	scanner := bufio.NewScanner(os.Stdin)

	y := 0
	var width int
	antennas := make(map[rune][]Pos)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		for x, c := range line {
			if c != '.' {
				pos := Pos{y: y, x: x}
				antennas[c] = append(antennas[c], pos)
			}
		}
		y++
		width = len(line)
	}

	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintf("Error reading standard input: %v", err))
	}

	return antennas, Pos{y: y, x: width}
}

func draw(antennas map[rune][]Pos, antinodes map[Pos]bool, size Pos) {
	antennasPerPos := make(map[Pos]rune)

	for antennaId, positions := range antennas {
		for _, pos := range positions {
			antennasPerPos[pos] = antennaId
		}
	}

	fmt.Println()
	for y := 0; y < size.y; y++ {
		for x := 0; x < size.x; x++ {
			pos := Pos{y: y, x: x}
			antennaId, isAntena := antennasPerPos[pos]
			isAntinode := antinodes[pos]
			if isAntena {
				fmt.Print(string(antennaId))
			} else if isAntinode {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}

		}
		fmt.Println()
	}
}

func getAntinode(p1 Pos, p2 Pos, size Pos) (Pos, bool) {
	y := p1.y + (p1.y - p2.y)
	x := p1.x + (p1.x - p2.x)
	pos := Pos{y: y, x: x}
	withinBounds := x >= 0 && x < size.x && y >= 0 && y < size.y
	return pos, withinBounds
}

func solve(antennas map[rune][]Pos, size Pos) (int, int) {
	antinodesP1 := make(map[Pos]bool)
	antinodesP2 := make(map[Pos]bool)
	for _, positions := range antennas {
		for i, p1 := range positions {
			for j, p2 := range positions {
				if i == j {
					continue
				}

				antinodesP2[p2] = true
				antinode, ok := getAntinode(p1, p2, size)
				if ok {
					antinodesP1[antinode] = true
					antinodesP2[antinode] = true

					comparePos := p1
					for {
						tmp := antinode
						antinode, ok = getAntinode(antinode, comparePos, size)
						comparePos = tmp
						if ok {
							antinodesP2[antinode] = true
						} else {
							break
						}
					}
				}
			}
		}
	}

	return len(antinodesP1), len(antinodesP2)
}

func main() {
	antennas, size := input()
	p1, p2 := solve(antennas, size)
	fmt.Println(p1)
	fmt.Println(p2)
}
