package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	count := 0
	part2 := 0
	groupSize := 0
	uniqueQuestions := map[string]int{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			count += len(uniqueQuestions)
			for _, c := range uniqueQuestions {
				if c == groupSize {
					part2 = part2 + 1
				}
			}
			uniqueQuestions = map[string]int{}
			groupSize = 0
			continue
		}
		groupSize = groupSize + 1
		splits := strings.Split(line, "")
		for _, s := range splits {
			uniqueQuestions[s] = uniqueQuestions[s] + 1
		}
	}

	fmt.Printf("count: %d\n", count)
	fmt.Printf("part2: %d\n", part2)
}
