package main

import "fmt"
func HasVowel(char rune) bool {
	return (char == 'A' || char == 'E' || char == 'I' || char == 'O' ||char == 'U') ||
		(char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u')
}
func HasConsonant (str string) bool {
	for _, char := range str {
		if !HasVowel(char) && char != ' ' {
			return true
		}
	}
	return false
}
func main (){
	fmt.Println(HasConsonant("aou"))
}
// this is great but if we are to pass a special case into the tester function, it will return true whch is not what we want.
//Create another file that correct this.
