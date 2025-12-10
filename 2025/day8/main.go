package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"sort"
	"strconv"
	"strings"
)

type Pos struct {
	x, y, z int
}

type Distance struct {
	a, b     Pos
	distance int
}

func pow2(a int) int {
	return a * a
}

func distance(pos1 Pos, pos2 Pos) int {
	return pow2(pos2.x-pos1.x) + pow2(pos2.y-pos1.y) + pow2(pos2.z-pos1.z)
}

type ByDistance []Distance

func (a ByDistance) Len() int           { return len(a) }
func (a ByDistance) Less(i, j int) bool { return a[i].distance < a[j].distance }
func (a ByDistance) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func input() []Distance {
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

	distances := []Distance{}
	for i, a := range boxes {
		for j, b := range boxes {
			if i >= j {
				continue
			}
			distances = append(distances, Distance{a, b, distance(a, b)})
		}
	}

	sort.Sort(ByDistance(distances))

	return distances
}

func circuitLen(circuits []map[Pos]bool) int {
	lens := []int{}
	for _, c := range circuits {
		lens = append(lens, len(c))
	}
	slices.Sort(lens)

	return lens[len(lens)-1] * lens[len(lens)-2] * lens[len(lens)-3]
}

func solve(distances []Distance) (int, int) {
	p1 := 0
	p2 := 0

	posMap := map[Pos]bool{}
	for _, p := range distances {
		posMap[p.a] = true
		posMap[p.b] = true
	}
	posCount := len(posMap)

	circuits := []map[Pos]bool{}
	for i := range 10000 {
		d := distances[i]
		found := []int{}
		for j, c := range circuits {
			if len(c) == posCount && p2 == 0 {
				p2 = distances[i-1].a.x * distances[i-1].b.x
			}
			if c[d.a] || c[d.b] {
				found = append(found, j)
				c[d.a] = true
				c[d.b] = true
			}
		}

		if p2 > 0 {
			break
		}

		if len(found) == 0 {
			m := map[Pos]bool{d.a: true, d.b: true}
			circuits = append(circuits, m)
		} else {
			// Merge
			for k := 1; k < len(found); k++ {
				for c := range circuits[found[k]] {
					circuits[found[0]][c] = true
				}
				circuits[found[k]] = make(map[Pos]bool)
			}
		}
		if i == 1000 {
			p1 = circuitLen(circuits)
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
