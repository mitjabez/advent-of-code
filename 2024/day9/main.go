package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type block struct {
	id             int
	size, free     int
	previous, next *block
}

func input() []block {
	scanner := bufio.NewScanner(os.Stdin)

	blocks := []block{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		var prev *block
		for i := 0; i < len(line); i += 2 {
			free := 0
			if i+1 < len(line) {
				free = int(line[i+1] - '0')
			}
			b := block{id: i / 2, size: int(line[i] - '0'), free: free}
			b.previous = prev
			if prev != nil {
				blocks[i/2-1].next = &b
				prev.next = &b
			}
			prev = &b
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
func printBlocks(b *block) {
	fmt.Println("blocks:")
	bb := b
	for {
		fmt.Println(bb)
		if bb.next == nil {
			break
		}
		bb = bb.next
	}
}

func part2(blocks []block) int {
	checksum := 0

	// right := &blocks[len(blocks)-1]
	var foundBlock *block
	pos := len(blocks) - 1

	var right *block
	var tmmp *block
	for {
		right = &blocks[pos]
		if right.previous == nil {
			break
		}
		foundBlock = nil
		b := right
		for {
			b = b.previous
			// fmt.Println(b, "prev:", b.previous)
			maxFree := math.MaxInt
			if right.id == 2 {
				// fmt.Println("b", b, "right", right)
			}
			if b.free >= right.size && b.free < maxFree {
				maxFree = b.free
				foundBlock = b
			}
			if b.previous == nil {
				break
			}
		}

		tmp := right.previous
		if foundBlock != nil {
			fmt.Println("Found block", foundBlock.id, foundBlock.free, "for:", right.id, "size:", right.size)
			right.previous.next = right.next
			if right.next != nil {
				right.next.previous = right.previous
			}

			right.free = foundBlock.free - right.size
			right.next = foundBlock.next
			right.previous = foundBlock

			foundBlock.free = 0
			foundBlock.next.previous = right
			foundBlock.next = right
			tmmp = foundBlock

		}
		right = tmp
		pos--
	}

	for {
		if tmmp.previous == nil {
			break
		}
		tmmp = tmmp.previous
	}

	printBlocks(tmmp)
	b2 := tmmp
	pos = 0
	for {
		if b2.next == nil {
			break
		}
		for i := 0; i < b2.size; i++ {
			checksum += b2.id * pos
			fmt.Print(b2.id)
			pos++
		}
		b2 = b2.next
	}

	return checksum

}
func main() {
	blocks := input()
	// fmt.Println(part1(blocks))
	fmt.Println(part2(blocks))
}
