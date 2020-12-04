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

	scanner := bufio.NewScanner(f)
	passports := []string{}
	p := &strings.Builder{}
	for scanner.Scan() {
		s := scanner.Text()
		if len(s) > 0 {
			p.WriteString(fmt.Sprintf(" %s", s))
		} else {
			passports = append(passports, p.String())
			p.Reset()
		}
	}
	if p.Len() > 0 {
		passports = append(passports, p.String())
		p.Reset()
	}

	count := 0
	for _, passport := range passports {
		if validPassport2(passport) {
			count++
		}
	}
	fmt.Println(count)
}

func validPassport(p string) bool {
	return strings.Contains(p, "byr:") &&
		strings.Contains(p, "iyr:") &&
		strings.Contains(p, "eyr:") &&
		strings.Contains(p, "hgt:") &&
		strings.Contains(p, "hcl:") &&
		strings.Contains(p, "ecl:") &&
		strings.Contains(p, "pid:")
}

func validPassport2(p string) bool {
	if !validPassport(p) {
		return false
	}

	lookup := map[string]string{}
	split := strings.Split(p, " ")
	for _, v := range split {
		kv := strings.Split(strings.TrimSpace(v), ":")
		if len(kv) < 2 {
			continue
		}
		lookup[kv[0]] = kv[1]
	}

	for k, v := range lookup {
		switch k {
		case "byr":
			year, _ := strconv.ParseInt(v, 10, 64)
			if year < 1920 || year > 2002 {
				return false
			}
		case "iyr":
			year, _ := strconv.ParseInt(v, 10, 64)
			if year < 2010 || year > 2020 {
				return false
			}
		case "eyr":
			year, _ := strconv.ParseInt(v, 10, 64)
			if year < 2020 || year > 2030 {
				return false
			}
		case "hgt":
			unit := v[len(v)-2:]
			hgt, _ := strconv.ParseInt(v[:len(v)-2], 10, 64)
			// in
			if unit == "in" {
				if hgt < 59 || hgt > 76 {
					return false
				}
				// cm
			} else if unit == "cm" {
				if hgt < 150 || hgt > 193 {
					return false
				}
			} else {
				return false
			}
		case "hcl":
			if !(strings.Contains(string(v[0]), "#") && len(v[1:]) == 6) {
				return false
			}
		case "ecl":
			if !validEyeColor(v) {
				return false
			}
		case "pid":
			if len(v) != 9 {
				return false
			}
		}
	}

	return true
}

func validEyeColor(color string) bool {
	validColors := map[string]bool{
		"amb": true,
		"blu": true,
		"brn": true,
		"gry": true,
		"grn": true,
		"hzl": true,
		"oth": true,
	}
	return validColors[color]
}
