package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

const sum = 2020

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	nums := []int64{}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		i, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, i)
	}

	twoSum(nums)
	threeSum(nums)
}

func twoSum(nums []int64) {
	m := map[int64]bool{}
	for _, v := range nums {
		m[v] = true
		if m[sum-v] {
			log.Print(v * (sum - v))
			return
		}
	}
	log.Fatal("no solution")
}

func threeSum(nums []int64) {
	for i, a := range nums {
		m := map[int64]bool{}
		twoSum := sum - a
		for j := i + 1; j < len(nums); j++ {
			b := nums[j]
			m[b] = true
			if m[twoSum-b] {
				log.Print(a * b * (twoSum - b))
				return
			}
		}
	}
	log.Fatal("no solution")
}
