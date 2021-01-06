package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func handle(e error) {
	if e != nil {
		panic(e)
	}
}

type result struct {
	First  int
	Second int
	Third  int
}

func retreiveValues(arr []string) result {
	for i := 0; i < len(arr); i++ {
		first, err := strconv.Atoi(arr[i])
		handle(err)
		for i := 0; i < len(arr); i++ {
			second, err := strconv.Atoi(arr[i])
			handle(err)
			for i := 0; i < len(arr); i++ {
				third, err := strconv.Atoi(arr[i])
				handle(err)
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
	data, err := ioutil.ReadFile("input.txt")
	handle(err)
	input := string(data)
	arr := strings.Split(input, "\n")
	result := retreiveValues(arr)
	fmt.Println(result.First * result.Second * result.Third)
}
