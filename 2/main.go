package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	part1 := 0
	part2 := 0
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		splits := strings.Split(scanner.Text(), " ")

		minimum, maximum := parseRule(splits[0])
		l := strings.Replace(splits[1], ":", "", -1)
		pass := splits[2]
		// part 1
		count := int64(strings.Count(pass, l))
		if count >= minimum && count <= maximum {
			part1++
		}

		// part 2
		x := string(pass[minimum-1]) == l
		y := string(pass[maximum-1]) == l
		if (x || y) && !(x && y) {
			part2++
		}

	}

	fmt.Println(part1)
	fmt.Println(part2)
}

func parseRule(s string) (int64, int64) {
	split := strings.Split(s, "-")
	minimum, err := strconv.ParseInt(split[0], 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	maximum, err := strconv.ParseInt(split[1], 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	return minimum, maximum
}
