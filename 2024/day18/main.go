package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var Dst Pos = Pos{70, 70}
var DstTest Pos = Pos{6, 6}

const Steps int = 1024
const StepsTest int = 12

type Pos struct {
	x, y int
}

type Node struct {
	x, y     int
	cost     int
	parents  []*Node
	children []*Node
}

func input() []Pos {
	scanner := bufio.NewScanner(os.Stdin)

	var input = []Pos{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		nums := strings.Split(line, ",")
		x, _ := strconv.Atoi(nums[0])
		y, _ := strconv.Atoi(nums[1])
		p := Pos{x, y}
		input = append(input, p)
	}

	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintf("Error reading standard input: %v", err))
	}

	return input
}

func enqueue(queue []*Node, p *Node) []*Node {
	return append(queue, p)
}

func deque(queue []*Node) ([]*Node, *Node, bool) {
	if len(queue) == 0 {
		invalid := Node{x: -1, y: -1}
		return queue, &invalid, false
	}

	n := queue[0]
	queue = queue[1:]
	return queue, n, true
}

func (n Node) Pos() Pos {
	return Pos{n.x, n.y}
}

func (n Node) Move(direction Pos) Node {
	return Node{
		x:        n.x + direction.x,
		y:        n.y + direction.y,
		cost:     n.cost,
		children: n.children,
		parents:  n.parents,
	}
}

func draw(dst Pos, fallen map[Pos]bool) {
	for y := range dst.y + 1 {
		for x := range dst.x + 1 {
			p := Pos{x, y}
			if fallen[p] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}

}

func bfs(start Node, dst Pos, fallen map[Pos]bool) *Node {
	visited := map[Pos]bool{}
	queue := []*Node{}
	queue = enqueue(queue, &start)
	start.cost = 0
	directions := []Pos{
		{1, 0},
		{-1, 0},
		{0, 1},
		{0, -1},
	}

	for {
		var ok bool
		var node *Node
		queue, node, ok = deque(queue)
		if !ok {
			return nil
		}

		for _, dir := range directions {
			n := node.Move(dir)
			if n.x < 0 || n.x > dst.x || n.y < 0 || n.y > dst.y || fallen[n.Pos()] || visited[n.Pos()] {
				continue
			}

			visited[n.Pos()] = true
			n.parents = append(n.parents, node)
			n.cost = node.cost + 1

			if n.Pos() == dst {
				return &n
			}

			queue = enqueue(queue, &n)
		}
	}
}

func part1(allFallen []Pos, dst Pos, steps int) int {
	fallen := make(map[Pos]bool)
	for i := 0; i < steps; i++ {
		fallen[allFallen[i]] = true
	}
	node := bfs(Node{x: 0, y: 0}, dst, fallen)
	if node == nil {
		panic("Didn't find dst")
	}
	return node.cost
}

func part2(allFallen []Pos, dst Pos, steps int) string {
	fallen := make(map[Pos]bool)
	for i := 0; i < steps; i++ {
		fallen[allFallen[i]] = true
	}

	for i := steps; i < len(allFallen); i++ {
		fallen[allFallen[i]] = true

		node := bfs(Node{x: 0, y: 0}, dst, fallen)
		if node == nil {
			return fmt.Sprintf("%d,%d", allFallen[i].x, allFallen[i].y)
		}
	}
	panic("Cannot solve part2")
}

func main() {
	// steps := StepsTest
	// dst := DstTest
	steps := Steps
	dst := Dst

	fallen := input()
	fmt.Println(part1(fallen, dst, steps))
	fmt.Println(part2(fallen, dst, steps))
}
