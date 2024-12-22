package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func input() []int {
	scanner := bufio.NewScanner(os.Stdin)

	var input = []int{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		n, _ := strconv.Atoi(line)
		input = append(input, n)
	}

	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintf("Error reading standard input: %v", err))
	}

	return input
}

func next(n int) int {
	r := n * 64
	r = n ^ r
	r %= 16777216

	r2 := r / 32
	r2 = r ^ r2
	r2 %= 16777216

	r3 := r2 * 2048
	r3 = r2 ^ r3
	r3 %= 16777216
	return r3
}

func part1(input []int) int {
	total := 0
	for _, nStart := range input {
		n := nStart
		for j := 0; j < 2000; j++ {
			n = next(n)
		}
		total += n
	}
	return total
}

type SeqKey struct {
	a, b, c, d int
}

func part2(input []int) int {
	/*
		{
			"-1,0,0,3": {
				"0": 8,
				"1": 2
				...
	*/
	pricesPerSequencePerNum := make(map[SeqKey]map[int]int)
	for numId, nStart := range input {
		n := nStart
		deltas := []int{}
		// maxPrice := 0
		for j := 0; j < 2000; j++ {
			prev := n
			n = next(n)
			price := n % 10
			delta := n%10 - prev%10
			deltas = append(deltas, delta)
			if len(deltas) >= 4 {
				key := SeqKey{deltas[len(deltas)-4], deltas[len(deltas)-3], deltas[len(deltas)-2], deltas[len(deltas)-1]}
				_, ok := pricesPerSequencePerNum[key]
				if !ok {
					pricesPerSequencePerNum[key] = make(map[int]int)
				}
				if pricesPerSequencePerNum[key][numId] == 0 {
					pricesPerSequencePerNum[key][numId] = price
				}
			}
		}
	}

	bestPrice := 0
	for _, pricePerNum := range pricesPerSequencePerNum {
		priceForSeq := 0
		for _, price := range pricePerNum {
			priceForSeq += price
		}
		if priceForSeq > bestPrice {
			bestPrice = priceForSeq
		}
	}

	return bestPrice
}
func main() {
	input := input()
	fmt.Println(part1(input))
	fmt.Println(part2(input))
}
