package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type productRange struct {
	start int
	end   int
}

func input() []productRange {
	scanner := bufio.NewScanner(os.Stdin)

	var ranges = []productRange{}
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		// Remove the annoying coma in test data
		line = strings.TrimSuffix(line, ",")
		for r := range strings.SplitSeq(line, ",") {
			ids := strings.Split(r, "-")
			idStart, _ := strconv.Atoi(ids[0])
			idEnd, _ := strconv.Atoi(ids[1])
			ranges = append(ranges, productRange{idStart, idEnd})
		}
	}

	if err := scanner.Err(); err != nil {
		panic(fmt.Sprintf("Error reading standard input: %v", err))
	}

	return ranges
}

func solve(input []productRange) (int, int) {
	p1 := 0
	p2 := 0

	for _, pr := range input {
		for numId := pr.start; numId <= pr.end; numId++ {
			strId := strconv.Itoa(numId)
			if strId[0:len(strId)/2] == strId[len(strId)/2:] {
				p1 += numId
			}

			for i := 1; i <= len(strId)/2; i++ {
				if len(strId)%i != 0 {
					continue
				}
				src := strId[0:i]
				j := 0
				for ; j < len(strId)/len(src)-1; j++ {
					startPos := i + j*len(src)
					target := strId[startPos : startPos+len(src)]
					if src != target {
						break
					}
				}
				if j == len(strId)/len(src)-1 {
					p2 += numId
					break
				}
			}
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
