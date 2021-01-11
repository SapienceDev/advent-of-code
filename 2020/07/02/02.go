package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/SapienceDev/advent-of-code/utils"
)

func main() {
	data := utils.GetInput(2020, 7)
	input := string(data)
	arr := strings.Split(input, "\n")

	bags := make(map[string]map[string]int)

	for i := range arr {
		re := regexp.MustCompile("([a-z]+ [a-z]+ bags)\\scontain")
		res := re.FindAllStringSubmatch(arr[i], -1)
		name := res[0][1]
		re = regexp.MustCompile("([1-9]) ([a-z]+ [a-z]+) bag")
		res = re.FindAllStringSubmatch(arr[i], -1)
		contains := make(map[string]int)
		for _, match := range res {
			val, _ := strconv.Atoi(match[1])
			contains[match[2]] = val
		}
		bags[name] = contains
	}

	total := find(bags, "shiny gold")
	fmt.Printf("Part 2: %d\n", total)
}

func find(bags map[string]map[string]int, s string) (total int) {
	for k, v := range bags {
		if strings.Contains(k, s) {
			for subk, subv := range v {
				total += subv
				for i := 0; i < subv; i++ {
					val := find(bags, subk)
					total += val
				}
			}
		}
	}
	return total
}
