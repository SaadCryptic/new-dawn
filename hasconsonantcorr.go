package main

import "fmt"
func HasVowel(char rune) bool {
	return (char == 'A' || char == 'E' || char == 'I' || char == 'O' ||char == 'U') ||
		(char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u')
}
func HasConsonant (str string) bool {
	for _, char := range str {
		if (char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z') && !HasVowel(char) {
			return true
		}
	}
	return false
}
func main (){
	fmt.Println(HasConsonant("Hello world"))
	fmt.Println(HasConsonant("aou"))
	fmt.Println(HasConsonant("1234"))
	fmt.Println(HasConsonant("!@#$"))
}
