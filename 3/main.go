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

	scanner := bufio.NewScanner(f)
	board := [][]string{}
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), "")
		board = append(board, split)
	}

	total := numTrees(board, 1, 1)
	total = total * numTrees(board, 3, 1)
	total = total * numTrees(board, 5, 1)
	total = total * numTrees(board, 7, 1)
	total = total * numTrees(board, 1, 2)
	fmt.Println(total)
}

func numTrees(board [][]string, right, down int) int {
	trees := 0
	y, x := 0, 0
	for y < len(board) {
		if board[y][x] == "#" {
			trees++
		}
		x = (x + right) % len(board[y])
		y = y + down
	}
	return trees
}
