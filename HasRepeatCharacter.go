package main

func HasRepeatedCharacter(str string) bool {
	dim := make(map[rune]bool)
	for _, char := range str {
		if dim[char] {
			return true
		}
		dim[char]=true
	}
	return false
}

