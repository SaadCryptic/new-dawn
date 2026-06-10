package main
func HasDigitGreaterThanFive(num string) bool {
	for _, digit := range num {
		if digit >= '6' && digit <= '9' {
			return true
		}
	}
	return false
}

func main (){
	println(HasDigitGreaterThanFive("123485")) // true
	println(HasDigitGreaterThanFive("67890")) // true
}
