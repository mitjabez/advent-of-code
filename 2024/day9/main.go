package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

type block struct {
	id         int
	size, free int
	moved      bool
}

func input() []block {
	scanner := bufio.NewScanner(os.Stdin)

	blocks := []block{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		for i := 0; i < len(line); i += 2 {
			free := 0
			if i+1 < len(line) {
				free = int(line[i+1] - '0')
			}
			b := block{id: i / 2, size: int(line[i] - '0'), free: free}
			blocks = append(blocks, b)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintf("Error reading standard input: %v", err))
	}

	return blocks
}

func part1(blocks []block) int {
	checksum := 0

	left := 0
	right := len(blocks) - 1
	pos := 0
	processLeft := true
	for {
		if processLeft {
			for i := 0; i < blocks[left].size; i++ {
				checksum += left * pos
				pos++
			}
			processLeft = false
		}

		for i := 0; i < blocks[right].size; i++ {
			if left == right {
				break
			}
			if blocks[left].free == 0 {
				break
			}

			checksum += right * pos
			pos++
			blocks[right].size--
			blocks[left].free--
			if blocks[right].size == 0 {
				right--
				i = 0
				continue
			}
		}

		if left == right {
			break
		}

		if blocks[left].free == 0 {
			left++
			processLeft = true
		}
	}

	return checksum
}

func part2(blocks []block) int {
	checksum := 0

	l := list.New()
	for _, b := range blocks {
		l.PushBack(b)
	}

	eR := l.Back()
	for {
		if eR.Prev() == nil {
			break
		}
		next := eR.Prev()

		blockRight := eR.Value.(block)
		if blockRight.moved {
			eR = next
			continue
		}

		for eL := l.Front(); eL != nil && eL != eR; eL = eL.Next() {
			blockLeft := eL.Value.(block)
			if blockLeft.free >= blockRight.size {
				if eR.Prev() != nil {
					bPrev := eR.Prev().Value.(block)
					bPrev.free += blockRight.size + blockRight.free
					eR.Prev().Value = bPrev
				}

				bFound := eL.Value.(block)

				blockRight.free = bFound.free - blockRight.size
				blockRight.moved = true
				eR.Value = blockRight

				bFound.free = 0
				eL.Value = bFound
				l.MoveAfter(eR, eL)
				break
			}
		}

		eR = next
	}

	pos := 0
	for e := l.Front(); e != nil; e = e.Next() {
		b := e.Value.(block)
		for i := 0; i < b.size; i++ {
			checksum += b.id * pos
			pos++
		}
		pos += b.free
	}

	return checksum

}
func main() {
	blocks := input()

	b1 := make([]block, len(blocks))
	copy(b1, blocks)
	fmt.Println(part1(b1))

	fmt.Println(part2(blocks))
}
