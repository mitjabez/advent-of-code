package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"

	// "math"
	"os"
	"strconv"
	"strings"
)

type Pos struct {
	x, y int
}

type Machine struct {
	a, b, prize Pos
}

func buttonToInt(pos int, line string) int {
	n, _ := strconv.Atoi(strings.Replace(strings.Replace(strings.Fields(line)[pos][1:], ",", "", -1), "=", "", -1))
	return n
}

func input() []Machine {
	scanner := bufio.NewScanner(os.Stdin)

	machines := []Machine{}
	var machine Machine
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		if strings.Index(line, "Button A") == 0 {
			machine = Machine{a: Pos{x: buttonToInt(2, line), y: buttonToInt(3, line)}}
		} else if strings.Index(line, "Button B") == 0 {
			machine.b = Pos{x: buttonToInt(2, line), y: buttonToInt(3, line)}
		} else if strings.Index(line, "Prize") == 0 {
			machine.prize = Pos{x: buttonToInt(1, line), y: buttonToInt(2, line)}
			machines = append(machines, machine)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintf("Error reading standard input: %v", err))
	}

	return machines
}

// func ss(p Pos, machine Machine, cost int, visited map[Pos]int) int {
// 	visited[p] = cost
//
// 	if p == machine.prize {
// 		return cost
// 	}
// 	return min(ss(Pos{p.x + machine.a.x, p.y + machine.a.y}, machine, cost+3, visited),
// 		ss(Pos{p.x + machine.b.x, p.y + machine.b.y}, machine, cost+1, visited))
// }

type Item struct {
	value    Pos // The value of the item; arbitrary.
	priority int // The priority of the item in the queue.
	btns     Pos
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	// return pq[i].priority < pq[j].priority
	return pq[i].priority < pq[j].priority //&& (pq[i].btns.x+pq[i].btns.y < pq[j].btns.x+pq[j].btns.y)
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (pq *PriorityQueue) update(item *Item, value Pos, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}

func minDistance(dists map[Pos]int, spts map[Pos]bool) Pos {
	minCost := math.MaxInt

	var minPos Pos
	for d, cost := range dists {
		if cost < minCost && !spts[d] {
			minCost = cost
			minPos = d
		}
	}
	return minPos
}
func push(pq *PriorityQueue, pos Pos, priority int, btns Pos) {
	item := &Item{
		value:    pos,
		priority: priority,
		btns:     btns,
	}
	heap.Push(pq, item)
}

func djikstra(machine Machine) int {
	start := Pos{0, 0}
	dists := make(map[Pos]int)
	dists[start] = 0
	btns := 0

	dpq := make(PriorityQueue, 0)
	heap.Init(&dpq)
	push(&dpq, start, 0, Pos{})

	spts := make(map[Pos]bool)

	// maxRep := 10
	// i := 0
	for {
		btns++
		item := heap.Pop(&dpq).(*Item)
		minPos := item.value

		// fmt.Println(minPos)
		// fmt.Println("spts", spts)
		// fmt.Println("dists", dists)
		// fmt.Println("--------------------------")
		// i++
		// if i == maxRep {
		// 	return 0
		// }

		if minPos == machine.prize {
			return item.priority
		}

		if item.btns.x > 100 && item.btns.y > 100 {
			// Not found
			return 0
		}

		// if minPos.x > machine.prize.x+1000 && minPos.y > machine.prize.y+1000 {
		// 	return 0
		// }

		spts[minPos] = true
		da := Pos{minPos.x + machine.a.x, minPos.y + machine.a.y}
		db := Pos{minPos.x + machine.b.x, minPos.y + machine.b.y}

		// Update dist value of the adjacent vertices
		// of the picked vertex only if the current
		// distance is greater than new distance and
		// the vertex in not in the shortest path tree
		curDist, ok := dists[da]
		if !spts[da] && (!ok || curDist > dists[minPos]+3) {
			dists[da] = dists[minPos] + 3
			push(&dpq, da, dists[minPos]+3, Pos{item.btns.x + 1, item.btns.y})
		}
		curDist, ok = dists[db]
		if !spts[db] && (!ok || curDist > dists[minPos]+1) {
			dists[db] = dists[minPos] + 1
			push(&dpq, db, dists[minPos]+1, Pos{item.btns.x, item.btns.y + 1})
		}
	}
}

func solve(machines []Machine) int {
	p1 := 0
	for i, m := range machines {
		fmt.Printf("%d. Searching ... ", i+1)
		r := djikstra(m)
		fmt.Println(r)
		p1 += r
	}
	return p1
}

func main() {
	machines := input()
	fmt.Println(solve(machines))
	// fmt.Println(p1)
	// fmt.Println(p2)
}
