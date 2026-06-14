package main
func CountLowercase(str string) int {
	count := 0
	for _, char := range str {
		if char >= 'a' && char <= 'z' {
			count++
		}
	}
	return count
}

