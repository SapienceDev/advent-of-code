package utils

// Handle handles errors
func Handle(e error) {
	if e != nil {
		panic(e)
	}
}
