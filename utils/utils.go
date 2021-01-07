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
