package utils

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

// Handle handles errors
func Handle(e error) {
	if e != nil {
		panic(e)
	}
}

// GetInput retreives the input files for the specified year and day
func GetInput(year int, day int) string {
	cmd := exec.Command("curl", fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day), "-H", fmt.Sprintf("Cookie: session=%s", os.Getenv("SESSION")))
	stdout, err := cmd.Output()
	Handle(err)
	return strings.TrimSpace(string(stdout))
}

// Contains determines whether str exists in arr
func Contains(arr []string, str string) bool {
	for _, val := range arr {
		if val == str {
			return true
		}
	}
	return false
}

// MapToInt maps each value of arr to an integer
func MapToInt(arr []string) (res []int) {
	for _, val := range arr {
		val, _ := strconv.Atoi(val)
		res = append(res, val)
	}
	return res
}
