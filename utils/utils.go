package utils

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Handle handles errors
func Handle(e error) {
	if e != nil {
		panic(e)
	}
}

// GetInput retreive the input files for the specified year and day
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
