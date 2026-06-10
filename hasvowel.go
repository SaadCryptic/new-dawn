package main

import "fmt"
func HasVowel(s string) bool {
	for _, char := range s {
		if char == 'A' || char == 'E' || char == 'I' || char == 'O' || char == 'U' || char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(HasVowel("Hello World"))
	fmt.Println(HasVowel("Hll Wrld"))
}
