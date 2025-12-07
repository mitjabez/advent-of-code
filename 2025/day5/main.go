package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type IdRange struct {
	start, end int
	ignore     bool
}

type ByStart []IdRange

func (a ByStart) Len() int           { return len(a) }
func (a ByStart) Less(i, j int) bool { return a[i].start < a[j].start }
func (a ByStart) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func input() ([]IdRange, []int) {
	scanner := bufio.NewScanner(os.Stdin)

	ranges := []IdRange{}
	ingredients := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		rs := strings.Split(line, "-")
		r1, _ := strconv.Atoi(rs[0])
		r2, _ := strconv.Atoi(rs[1])
		ranges = append(ranges, IdRange{r1, r2, false})
	}

	for scanner.Scan() {
		line := scanner.Text()
		ing, _ := strconv.Atoi(line)
		ingredients = append(ingredients, ing)
		if line == "" {
			continue
		}
	}

	sort.Sort(ByStart(ranges))

	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintf("Error reading standard input: %v", err))
	}

	return ranges, ingredients
}

// func binarySearch(id int, ranges []IdRange) bool {
// 	start := 0
// 	end := len(ranges) - 1
//
// 	for {
// 		if id >= ranges[start].start && id <= ranges[start].end {
// 			return true
// 		} else if id >= ranges[end].start && id <= ranges[end].end {
// 			return true
// 		}
// 		// fmt.Printf("start: %d, end: %d\n", start, end)
// 		if end-start <= 1 {
// 			// if start == end {
// 			break
// 		}
// 		// 0 2
// 		// mid = 1
// 		// 0 1 or 1 2
//
// 		// 0 1
// 		// mid 0
// 		mid := (end - start) / 2
// 		if id < ranges[start+mid].start {
// 			end = start + mid
// 		} else {
// 			start = start + mid
// 		}
// 	}
//
// 	return false
// }

func search(id int, ranges []IdRange) bool {
	for _, r := range ranges {
		if id >= r.start && id <= r.end {
			return true
		}
	}
	return false
}

func solve(ranges []IdRange, ingredients []int) (int, int) {
	p1 := 0
	p2 := 0

	// part1
	for _, ingredient := range ingredients {
		if search(ingredient, ranges) {
			p1++
		}
	}

	// part2
	for i, r := range ranges {
		iPrev := i - 1
		for iPrev >= 0 && ranges[iPrev].ignore {
			iPrev--
		}

		if iPrev >= 0 {
			if r.start <= ranges[iPrev].end {
				newStart := ranges[iPrev].end + 1
				if newStart > r.end {
					r.ignore = true
					ranges[i] = r
					continue
				}
				r.start = newStart
			}
		}

		p2 += r.end - r.start + 1
	}

	return p1, p2
}

func main() {
	ranges, ingredients := input()
	p1, p2 := solve(ranges, ingredients)
	fmt.Println(p1)
	fmt.Println(p2)
}
