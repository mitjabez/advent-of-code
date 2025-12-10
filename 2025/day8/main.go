package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

type Pos struct {
	x, y, z int
}

func (pos1 Pos) distance(pos2 Pos) int {
	dist := math.Sqrt(math.Pow(float64(pos2.x-pos1.x), 2) +
		math.Pow(float64(pos2.y-pos1.y), 2) +
		math.Pow(float64(pos2.z-pos1.z), 2))
	return int(dist)
}

func input() []Pos {
	scanner := bufio.NewScanner(os.Stdin)

	boxes := []Pos{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		tokens := strings.Split(line, ",")
		x, _ := strconv.Atoi(tokens[0])
		y, _ := strconv.Atoi(tokens[1])
		z, _ := strconv.Atoi(tokens[2])
		boxes = append(boxes, Pos{x, y, z})
	}

	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintf("Error reading standard input: %v", err))
	}

	return boxes
}

func closest(boxes map[Pos][]Pos) (Pos, Pos, int) {
	min := math.MaxInt
	var minPair [2]Pos
	nJoined := 0
	for box := range boxes {
		for destBox, destAdj := range boxes {
			if box == destBox || slices.Contains(destAdj, box) {
				continue
			}

			nJoined++
			d := box.distance(destBox)
			if d < min {
				minPair[0] = box
				minPair[1] = destBox
				min = d
			}
		}
	}
	return minPair[0], minPair[1], nJoined
}

func circuitSize(box Pos, boxes map[Pos][]Pos, visited map[Pos]bool, size int) int {
	visited[box] = true

	for _, adj := range boxes[box] {
		if visited[adj] {
			continue
		}

		// fmt.Printf("->%v", adj)
		visited[adj] = true
		size++
		if len(boxes[adj]) > 0 {
			size += circuitSize(adj, boxes, visited, 0)
		}
	}

	return size
}

func solve(arrBoxes []Pos) (int, int) {
	p1 := 0
	p2 := 0

	boxes := map[Pos][]Pos{}
	for _, b := range arrBoxes {
		boxes[b] = make([]Pos, 0)
	}

	// var lb1 Pos
	// var lb2 Pos
	for i := range 100000 {
		b1, b2, nJoined := closest(boxes)

		boxes[b1] = append(boxes[b1], b2)
		boxes[b2] = append(boxes[b2], b1)
		visited := map[Pos]bool{}
		cs := circuitSize(b1, boxes, visited, 1)
		fmt.Printf("#%d. joined=%d. size: %d\n", i, nJoined, cs)
		if cs == len(boxes) {
			fmt.Println(b1, b2)
			p2 = b1.x * b2.x
			break
		}
	}

	circuitSizes := []int{}
	visited := map[Pos]bool{}
	for box := range boxes {
		// fmt.Printf("box %v:\n", box)
		size := circuitSize(box, boxes, visited, 1)
		// fmt.Printf("\nsize: %d\n", size)
		fmt.Printf("------------------------------------------------\n")
		circuitSizes = append(circuitSizes, size)
	}
	slices.Sort(circuitSizes)
	p1 = circuitSizes[len(circuitSizes)-1] * circuitSizes[len(circuitSizes)-2] * circuitSizes[len(circuitSizes)-3]

	return p1, p2
}

func main() {
	input := input()
	p1, p2 := solve(input)
	fmt.Println(p1)
	fmt.Println(p2)
}
