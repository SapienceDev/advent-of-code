package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/SapienceDev/advent-of-code/utils"
)

func slice(arr []int, key string) []int {
	if key == "F" || key == "L" {
		return arr[0 : len(arr)/2]
	}
	return arr[len(arr)/2:]
}

func get(rows []int, keys []string) int {
	var arr []int = rows[:]
	for i := 0; i < len(keys); i++ {
		arr = slice(arr[:], keys[i])
	}
	return arr[0]
}

func main() {
	input := utils.GetInput(2020, 5)
	passes := strings.Split(input, "\n")

	var rows [128]int
	for index := range rows {
		rows[index] = index
	}

	var columns [8]int
	for index := range columns {
		columns[index] = index
	}

	var highest int = 0
	var ids [872]int
	for index, pass := range passes {
		var rowKeys []string = strings.Split(pass, "")[0:7]
		var row int = get(rows[:], rowKeys)

		var columnKeys []string = strings.Split(pass, "")[7:]
		var column int = get(columns[:], columnKeys)

		var seatID int = row*8 + column
		ids[index] = seatID
		if seatID > highest {
			highest = seatID
		}
	}
	fmt.Printf("Highest Seat ID: %d\n", highest)
	sort.Ints(ids[:])
	for index, id := range ids {
		if index == 871 {
			break
		} else if ids[index+1] != id+1 {
			fmt.Printf("My Seat ID: %d", id+1)
		}
	}
}
