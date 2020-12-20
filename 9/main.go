package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const preambleLen = 25

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	nums := []int{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		n, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}
		nums = append(nums, n)
	}

	weakness := part1(nums)
	fmt.Printf("part1: %d\n", weakness)
	fmt.Printf("part2: %d\n", part2(nums, weakness))
}

func part1(nums []int) int {
	stack := []int{}
	contains := map[int]bool{}
	for _, n := range nums {
		if len(stack) < 25 {
			stack = append(stack, n)
			contains[n] = true
			continue
		}
		if !twoSum(stack, contains, n) {
			return n
		}
		contains[stack[0]] = false
		contains[n] = true
		stack = append(stack[1:], n)
	}
	return 0
}

func part2(allNums []int, n int) int {
	sum := 0
	i := 0
	currentNums := []int{}
	for sum != n {
		v := allNums[i]
		currentNums = append(currentNums, v)
		sum = sum + v
		if sum > n {
			for sum > n {
				sum = sum - currentNums[0]
				currentNums = currentNums[1:]
			}
		}
		i++
	}
	min, max := int(^uint(0)>>1), 0
	for _, num := range currentNums {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}
	return min + max
}

func twoSum(nums []int, contains map[int]bool, n int) bool {
	for _, a := range nums {
		if contains[n-a] {
			return true
		}
	}
	return false
}
