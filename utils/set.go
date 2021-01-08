package utils

import "regexp"

// Set is a struct with a single property, values, which contains no repeated elements
type Set struct {
	Values []string
}

// New creates a new set from an array, removing duplicate elements
func (s *Set) New(values []string) *Set {
	for _, val := range values {
		matches, _ := regexp.MatchString("\\s", val)
		if !matches {
			s.Add(val)
		}
	}
	return s
}

// Add appends a string to the end of a set
func (s *Set) Add(str string) {
	if !s.Contains(str) {
		s.Values = append(s.Values, str)
	}
}

// Contains returns a boolean value specifying whether an element is contained the set
func (s *Set) Contains(str string) bool {
	for _, val := range s.Values {
		if val == str {
			return true
		}
	}
	return false
}
