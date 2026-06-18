package main
func CountConsonants(arg string) int {
	count := 0
	for _, char := range arg {
		if (char >= 'a' && char <= 'z') || (char >= 'A' && char <= 'Z') {
			if !(char == 'A' || char == 'E' || char == 'I' || char == 'O' || char == 'U' ||
				char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u') {
					count++
				}
		}
	}
	return count
}

