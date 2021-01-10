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
	arr := strings.Split(input, "\n")

	bags := make(map[string]string)

	for i := range arr {
		re := regexp.MustCompile("([a-z]+ [a-z]+ bags)\\scontain")
		res := re.FindAllStringSubmatch(arr[i], -1)
		name := res[0][1]
		re = regexp.MustCompile("([1-9]) ([a-z]+ [a-z]+) bag")
		res = re.FindAllStringSubmatch(arr[i], -1)
		contains := ""
		for _, match := range res {
			contains += fmt.Sprintf(" %s bags", match[2])
		}
		bags[name] = strings.Trim(contains, " ")
	}

	ans := find(bags, "shiny gold")
	fmt.Printf("Part 1: %d\n", len(ans))
}

func find(bags map[string]string, s string) (res []string) {
	for k, v := range bags {
		if strings.Contains(v, s) {
			res = append(res, k)
			ans := find(bags, k)
			if len(ans) > 0 {
				for _, val := range ans {
					if !contains(res, val) {
						res = append(res, val)
					}
				}
			}
		}
	}
	return res
}

func contains(arr []string, str string) bool {
	for _, val := range arr {
		if val == str {
			return true
		}
	}
	return false
}
