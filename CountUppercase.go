package main
func CountUppercase(str string) int {
	count := 0
	for _, char := range str {
		if char >= 'A' && char <= 'Z' {
			count++
		}
	}
	return count
}

