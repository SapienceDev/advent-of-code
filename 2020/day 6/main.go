package main

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/SapienceDev/advent-of-code/utils"
)

func main() {
	input := utils.ReadFile("input.txt")
	re := regexp.MustCompile("\\n\\s")
	arr := re.Split(input, -1)
	var total int = 0
	for _, group := range arr {
		set := utils.Set{}
		set.New(strings.Split(group, ""))
		total += len(set.Values)
	}
	fmt.Printf("Part 1: %d\n", total)

	// -----------------------------------

	var sum int = 0
	for _, group := range arr {
		persons := strings.Split(group, "\n")
		questions := strings.Split(persons[0], "")
		for _, question := range questions {
			var yes int = 0
			for _, person := range persons {
				if contains(strings.Split(person, ""), question) {
					yes++
				}
			}
			if yes == len(persons) {
				sum++
			}
		}
	}
	fmt.Printf("Part 2: %d", sum)
}

func contains(arr []string, str string) bool {
	for _, val := range arr {
		if val == str {
			return true
		}
	}
	return false
}
