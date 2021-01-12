package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/SapienceDev/advent-of-code/utils"
)

func main() {
	input := utils.GetInput(2020, 9)
	arr := mapToInt(strings.Split(string(input), "\n"))
	findInvalidNum(arr, 25)
	for i := 2; i < len(arr); i++ {
		findSum(arr, 0, i)
	}
}

func findSum(arr []int, start int, end int) {
	if end > len(arr) {
		return
	}
	subarr := arr[start:end]
	var sum int = 0
	for _, val := range subarr {
		sum += val
	}
	if sum == 15690279 {
		sort.Ints(subarr)
		fmt.Printf("Part 2: %d\n", subarr[0]+subarr[len(subarr)-1])
		return
	}
	findSum(arr, start+1, end+1)
}

func findInvalidNum(arr []int, size int) (invalid int) {
	preamble := arr[0:size]
	for _, val := range arr[size:] {
		var valid bool
		for _, x := range preamble {
			for _, y := range preamble {
				if x != y && x+y == val {
					valid = true
				}
			}
		}
		if !valid {
			fmt.Printf("Part 1: %d\n", val)
		}
		preamble = append(preamble[1:size], val)
	}
	return invalid
}

func mapToInt(arr []string) (res []int) {
	for _, val := range arr {
		val, _ := strconv.Atoi(val)
		res = append(res, val)
	}
	return res
}
