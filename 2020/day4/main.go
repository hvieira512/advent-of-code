package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/hvieira512/advent-of-code-2020/utils"
)

var requiredFields = []string{
	"byr", "iyr", "eyr", "hgt",
	"hcl", "ecl", "pid",
}

var validators = map[string]func(string) bool{
	"byr": func(v string) bool {
		n, err := strconv.Atoi(v)
		return err == nil && len(v) == 4 && n >= 1920 && n <= 2002
	},
	"iyr": func(v string) bool {
		n, err := strconv.Atoi(v)
		return err == nil && len(v) == 4 && n >= 2010 && n <= 2020
	},
	"eyr": func(v string) bool {
		n, err := strconv.Atoi(v)
		return err == nil && len(v) == 4 && n >= 2020 && n <= 2030
	},
	"hgt": func(v string) bool {
		if n, err := strconv.Atoi(strings.TrimSuffix(v, "cm")); err == nil && strings.HasSuffix(v, "cm") {
			return n >= 150 && n <= 193
		}

		if n, err := strconv.Atoi(strings.TrimSuffix(v, "in")); err == nil && strings.HasSuffix(v, "in") {
			return n >= 59 && n <= 76
		}

		return false
	},
	"hcl": func(v string) bool {
		matched, _ := regexp.MatchString(`^#[0-9a-f]{6}$`, v)
		return matched
	},
	"ecl": func(v string) bool {
		switch v {
		case "amb", "blu", "brn", "gry", "grn", "hzl", "oth":
			return true
		default:
			return false
		}
	},
	"pid": func(v string) bool {
		matched, _ := regexp.MatchString(`^[0-9]{9}$`, v)
		return matched
	},
}

func parsePassport(passport string) map[string]string {
	fields := make(map[string]string)

	for _, field := range strings.Fields(passport) {
		key, value, ok := strings.Cut(field, ":")
		if ok {
			fields[key] = value
		}
	}

	return fields
}

func hasRequiredFields(passport map[string]string) bool {
	for _, field := range requiredFields {
		if _, exists := passport[field]; !exists {
			return false
		}
	}

	return true
}

func isValidPassport(passport map[string]string) bool {
	for field, validator := range validators {
		value, exists := passport[field]

		if !exists || !validator(value) {
			return false
		}
	}

	return true
}

func part1(data []string) int {
	valid := 0

	for _, raw := range data {
		if hasRequiredFields(parsePassport(raw)) {
			valid++
		}
	}

	return valid
}

func part2(data []string) int {
	valid := 0

	for _, raw := range data {
		passport := parsePassport(raw)

		if hasRequiredFields(passport) && isValidPassport(passport) {
			valid++
		}
	}

	return valid
}

func main() {
	data := utils.ReadLinesBlocks("input.txt")

	fmt.Println("Part 1:", part1(data))
	fmt.Println("Part 2:", part2(data))
}
