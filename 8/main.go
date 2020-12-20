package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type InstructionType string

const (
	ACC InstructionType = "acc"
	JMP InstructionType = "jmp"
	NOP InstructionType = "nop"
)

type Instruction struct {
	Type       InstructionType
	Value      int
	IsNegative bool
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	instructions := []*Instruction{}
	for scanner.Scan() {
		instructions = append(instructions, parseInstruction(scanner.Text()))
	}
	fmt.Println(part1(instructions))
	fmt.Println(part2(instructions))
}

func part1(instructions []*Instruction) (int, bool) {
	seenInstructions := map[int]bool{}
	i := 0
	accumulator := 0
	for i < len(instructions) {
		if seenInstructions[i] {
			return accumulator, true
		}
		seenInstructions[i] = true
		instruction := instructions[i]

		switch instruction.Type {
		case ACC:
			i++
			if instruction.IsNegative {
				accumulator = accumulator - instruction.Value
			} else {
				accumulator = accumulator + instruction.Value
			}
		case JMP:
			if instruction.IsNegative {
				i = i - instruction.Value
			} else {
				i = i + instruction.Value
			}
		case NOP:
			i++
		}
	}

	return accumulator, false

}

func part2(instructions []*Instruction) int {
	for _, instruction := range instructions {
		typ := instruction.Type
		if instruction.Type == NOP {
			instruction.Type = JMP
		} else if instruction.Type == JMP {
			instruction.Type = NOP
		}
		acc, looped := part1(instructions)
		if !looped {
			return acc
		}
		instruction.Type = typ
	}
	return 0
}

func parseInstruction(line string) *Instruction {
	split := strings.Split(line, " ")
	typ := InstructionType(split[0])
	if typ != ACC && typ != JMP && typ != NOP {
		panic("unsupported instruction type")
	}

	isNegative := string(split[1][0]) == "-"

	value, err := strconv.Atoi(split[1][1:])
	if err != nil {
		panic("unable to parse value")
	}

	return &Instruction{
		Type:       typ,
		Value:      value,
		IsNegative: isNegative,
	}
}
