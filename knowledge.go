package main

import "fmt"
func CheckSpaces (s string) bool  {
	for _, char := range s {
		if char == ' ' {
			return true
		}
	}
	return false
}
func main() {
	fmt.Println(CheckSpaces("Hello World"))
	fmt.Println(CheckSpaces("HelloWorld"))
	fmt.Println(CheckSpaces(" "))
}
