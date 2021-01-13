package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/SapienceDev/advent-of-code/utils"
)

func main() {
	input := utils.GetInput(2020, 10)
	arr := utils.MapToInt(strings.Split(input, "\n"))
	sort.Ints(arr)
	var three int = 0
	var one int = 1
	for index, val := range arr {
		var next int = 0
		if index+1 >= len(arr) {
			next = arr[len(arr)-1] + 3
		} else {
			next = arr[index+1]
		}
		diff := next - val
		if diff == 3 {
			three++
		} else if diff == 1 {
			one++
		}
	}

	fmt.Printf("Part 1: %d\n", one*three)
}
