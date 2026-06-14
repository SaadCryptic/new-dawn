package main
func CountVowels(str string) int {
	count := 0
	for _, char := range str {
		if (char == 'A' || char == 'E' || char == 'I' || char == 'O' || char == 'U') ||
			(char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u') {
			count++
		}
	}
	return count
}

