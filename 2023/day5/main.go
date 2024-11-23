package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func input() []string {
	scanner := bufio.NewScanner(os.Stdin)

	var input = []string{}
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" {
			input = append(input, line)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintf("Error reading standard input: %v", err))
	}

	return input
}

func main() {
	input := input()

	ids := make([]int, 0)
	ids2 := make([]int, 0)
	var visited map[int]bool
	var visited2 map[int]bool

	for n, line := range input {
		if n == 0 {
			for _, seedIdStr := range strings.Split(strings.TrimLeft(line, "seeds: "), " ") {
				seedIdInt, _ := strconv.Atoi(seedIdStr)
				ids = append(ids, seedIdInt)
				ids2 = append(ids2, seedIdInt)
			}

			continue
		}

		if line == "" {
			continue
		}

		if strings.Contains(line, "map:") {
			visited = make(map[int]bool)
			visited2 = make(map[int]bool)
			continue
		}

		rangeTokens := strings.Split(line, " ")
		dst, _ := strconv.Atoi(rangeTokens[0])
		src, _ := strconv.Atoi(rangeTokens[1])
		rangeLen, _ := strconv.Atoi(rangeTokens[2])

		for i, id := range ids {
			if !visited[i] && id >= src && id < src+rangeLen {
				ids[i] = dst + (id - src)
				visited[i] = true
			}

		}

		// part 2
		for i := 0; i < len(ids2); i += 2 {
			idStart := ids2[i]
			idEnd := idStart + ids2[i+1] - 1
			srcStart := src
			srcEnd := src + rangeLen - 1

			// Within range
			if !visited2[i] && idStart < srcEnd && idEnd > srcStart {
				startRange := idStart
				endRange := idEnd
				if idStart < srcStart {
					ids2 = append(ids2, idStart, srcStart-idStart)
					startRange = src
				}

				if idEnd > srcEnd {
					ids2 = append(ids2, srcEnd+1, idEnd-srcEnd)
					endRange = srcEnd
				}

				ids2[i] = dst + (startRange - srcStart)
				ids2[i+1] = endRange - startRange + 1
				visited2[i] = true
			}
		}
	}

	part1 := math.MaxInt
	for _, id := range ids {
		part1 = min(id, part1)
	}

	part2 := math.MaxInt
	for i := 0; i < len(ids2); i += 2 {
		part2 = min(ids2[i], part2)
	}

	fmt.Println(part1)
	fmt.Println(part2)
}
