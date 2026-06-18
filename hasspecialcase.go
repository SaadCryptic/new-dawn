package main
func HasSpecialCase(s string) bool {
	for _, char := range s {
		if !(char >='A' && char <= 'Z' || char >='a' && char <= 'z' || char >='0' && char <= '9' || char == ' ') {
			return true
		}
	}
	return false
}


