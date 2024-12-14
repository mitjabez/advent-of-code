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

func compute(m Machine) (int, int) {
	b := (m.prize.y*m.a.x - m.prize.x*m.a.y) / (m.b.y*m.a.x - m.b.x*m.a.y)
	a := (m.prize.x - (b * m.b.x)) / m.a.x
	if a*m.a.x+b*m.b.x == m.prize.x && a*m.a.y+b*m.b.y == m.prize.y {
		return a, b
	} else {
		return 0, 0
	}

}

func solve(machines []Machine) (int, int) {
	p1 := 0
	p2 := 0
	for _, m := range machines {
		a, b := compute(m)
		p1 += a*3 + b*1

		m.prize.x += 10000000000000
		m.prize.y += 10000000000000
		a, b = compute(m)
		p2 += a*3 + b*1
	}
	return p1, p2
}

func main() {
	machines := input()
	p1, p2 := solve(machines)
	fmt.Println(p1)
	fmt.Println(p2)
}
