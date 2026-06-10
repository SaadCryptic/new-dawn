package main
func HasAlphabetOnly(s string)bool {
	for _, char := range s {
		if ((char >= 'A' && char <= 'Z') ||(char >= 'a' && char <= 'z')){
			return true
		}
	}
	return false
}
