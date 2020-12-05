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

	highest := 0
	allSeatIDs := map[int]bool{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		leftRow, rightRow := 0, 127
		leftColumn, rightColumn := 0, 7
		split := strings.Split(scanner.Text(), "")
		for _, s := range split {
			switch s {
			case "F":
				rightRow = int((rightRow + leftRow) / 2)
			case "B":
				leftRow = int((rightRow+leftRow)/2) + 1
			case "R":
				leftColumn = int((rightColumn+leftColumn)/2) + 1
			case "L":
				rightColumn = int((rightColumn + leftColumn) / 2)
			}
		}
		if leftRow != rightRow || leftColumn != rightColumn {
			log.Fatal("wut")
		}
		seatID := leftRow*8 + leftColumn
		allSeatIDs[seatID] = true
		if seatID > highest {
			highest = seatID
		}
	}

	fmt.Printf("highest: %d\n", highest)
	// we know we are not at the very front or back
	for i := 1; i < highest; i++ {
		if !allSeatIDs[i] && allSeatIDs[i-1] && allSeatIDs[i+1] {
			fmt.Printf("my seat: %d\n", i)
			break
		}
	}
}
