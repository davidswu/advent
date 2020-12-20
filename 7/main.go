package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Color struct {
	Name   string
	Amount int
}

func main() {

	rules := parseRules()

	// reverseRules simply flips rules without regard to amount
	reverseRules := map[string][]string{}
	for color, colors := range rules {
		for _, c := range colors {
			reverseRules[c.Name] = append(reverseRules[c.Name], color)
		}
	}
	search := []string{"shiny gold"}
	containsGold := map[string]bool{}
	for len(search) > 0 {
		c := search[0]
		search = search[1:]

		for _, color := range reverseRules[c] {
			containsGold[color] = true
			search = append(search, color)
		}
	}

	fmt.Println(len(containsGold))
	fmt.Println(part2(rules, "shiny gold"))
}

func part2(rules map[string][]Color, c string) int {
	sum := 0
	for _, inside := range rules[c] {
		sum = sum + inside.Amount + (inside.Amount * part2(rules, inside.Name))
	}

	return sum
}

func parseRules() map[string][]Color {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	rules := map[string][]Color{}

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " ")
		c := strings.Join(split[:2], " ")
		colors := []Color{}

		for i := 4; i < len(split)-2; i += 4 {
			amount, err := strconv.Atoi(split[i])
			if err != nil {
				continue
			}
			colors = append(colors, Color{
				Name:   strings.Join(split[i+1:i+3], " "),
				Amount: amount,
			})
		}

		rules[c] = colors
	}

	return rules
}
