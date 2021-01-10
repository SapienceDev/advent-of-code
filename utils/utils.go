package utils

import "io/ioutil"

// Handle handles errors
func Handle(e error) {
	if e != nil {
		panic(e)
	}
}

// ReadFile reads an input file and returns a string representation of that file
func ReadFile(file string) string {
	data, err := ioutil.ReadFile(file)
	Handle(err)
	input := string(data)
	return input
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
