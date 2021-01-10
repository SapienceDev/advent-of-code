package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/SapienceDev/advent-of-code/utils"
)

func main() {
	data := utils.ReadFile("input.txt")
	input := string(data)
	bags := strings.Split(input, "\n")

	var valid []string

	// get all direct matches
	for i, bag := range bags {
		re := regexp.MustCompile("([a-z]+ [a-z]+ bags) .+ [1-9] shiny gold bag")
		if re.MatchString(bag) {
			res := re.FindAllStringSubmatch(bags[i], -1)
			name := res[0][1]
			valid = append(valid, fmt.Sprintf("%s?", name))
		}
	}

	// get all indirect matches
	for {
		var oldCount int = len(valid)
		for i, bag := range bags {
			re := regexp.MustCompile(fmt.Sprintf("([a-z]+ [a-z]+ [a-z]+) .+ [1-9] (%s)", strings.Join(valid, "|")))
			if re.MatchString(bag) {
				res := re.FindAllStringSubmatch(bags[i], -1)
				name := res[0][1]
				if !contains(valid, fmt.Sprintf("%s?", name)) {
					valid = append(valid, fmt.Sprintf("%s?", name))
				}
			}
		}
		if oldCount == len(valid) {
			break
		}
	}
	fmt.Printf("Part 1: %d", len(valid))
}

func contains(arr []string, str string) bool {
	for _, val := range arr {
		if val == str {
			return true
		}
	}
	return false
}
