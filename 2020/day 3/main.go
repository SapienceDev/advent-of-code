package main

import (
	"fmt"
	"strings"

	"github.com/SapienceDev/advent-of-code/utils"
)

func create2D(arr []string) [323][31]string {
	var twoD [323][31]string
	col := 0
	row := 0
	for i := 0; i < len(arr); i++ {
		if col == 31 {
			col = 0
			row++
		} else {
			twoD[row][col] = arr[i]
			col++
		}
	}
	return twoD
}

func walk(arr [323][31]string, right int, down int) {
	col := 0
	trees := 0
	for row := down; row < 323; row += down {
		col += right
		if col > 10 {
			col %= 31
		}
		if arr[row][col] == "#" {
			trees++
		}
	}
	fmt.Printf("slope: right %d, down %d | %d\n", right, down, trees)
}

func main() {
	input := utils.ReadFile("input.txt")
	arr := strings.Split(input, "")
	twoD := create2D(arr)
	walk(twoD, 1, 1)
	walk(twoD, 3, 1)
	walk(twoD, 5, 1)
	walk(twoD, 7, 1)
	walk(twoD, 1, 2)
}
