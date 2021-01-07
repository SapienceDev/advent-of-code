package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/SapienceDev/advent-of-code/utils"
)

type result struct {
	First  int
	Second int
	Third  int
}

func retreiveValues(arr []string) result {
	for i := 0; i < len(arr); i++ {
		first, err := strconv.Atoi(arr[i])
		utils.Handle(err)
		for i := 0; i < len(arr); i++ {
			second, err := strconv.Atoi(arr[i])
			utils.Handle(err)
			for i := 0; i < len(arr); i++ {
				third, err := strconv.Atoi(arr[i])
				utils.Handle(err)
				if first+second+third == 2020 {
					return result{
						First:  first,
						Second: second,
						Third:  third,
					}
				}
			}
		}
	}
	return result{}
}

func main() {
	input := utils.ReadFile("input.txt")
	arr := strings.Split(input, "\n")
	result := retreiveValues(arr)
	fmt.Println(result.First * result.Second * result.Third)
}
