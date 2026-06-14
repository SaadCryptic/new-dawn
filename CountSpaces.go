package main
func CountSpacescode(str string) int {
	count := 0
	for _, char := range str {
		if char == ' ' {
			count++
		}
	}
	return count
}

