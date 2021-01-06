package main

import (
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
)

func handle(e error) {
	if e != nil {
		panic(e)
	}
}

func partOne(arr []string) int {
	var valid int = 0
	for i := 0; i < len(arr); i++ {
		re := regexp.MustCompile("([0-9-]+)\\s([a-z]): ([a-z]+)")
		res := re.FindAllStringSubmatch(arr[i], -1)

		num := res[0][1]
		letter := res[0][2]
		password := res[0][3]

		min, err := strconv.Atoi(strings.Split(num, "-")[0])
		handle(err)
		max, err := strconv.Atoi(strings.Split(num, "-")[1])
		handle(err)

		if strings.Count(password, letter) < min || strings.Count(password, letter) > max {
			continue
		}
		valid++
	}
	return valid
}

func partTwo(arr []string) int {
	var valid int = 0
	for i := 0; i < len(arr); i++ {
		re := regexp.MustCompile("([0-9-]+)\\s([a-z]): ([a-z]+)")
		res := re.FindAllStringSubmatch(arr[i], -1)

		num := res[0][1]
		letter := res[0][2]
		password := res[0][3]

		first, err := strconv.Atoi(strings.Split(num, "-")[0])
		handle(err)
		second, err := strconv.Atoi(strings.Split(num, "-")[1])
		handle(err)
		chars := strings.Split(password, "")
		if chars[first-1] == letter && !(chars[second-1] == letter) {
			valid++
		} else if chars[second-1] == letter && !(chars[first-1] == letter) {
			valid++
		}
	}
	return valid
}

func main() {
	data, err := ioutil.ReadFile("input.txt")
	handle(err)
	input := string(data)
	var arr []string = strings.Split(input, "\n")
	partOne(arr)
	partTwo(arr)
}
