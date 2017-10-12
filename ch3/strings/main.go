package main

import (
	"fmt"
	"unicode/utf8"
)

//import "fmt"

func main()  {
	fmt.Println(Contains("12343566", "123"))

	s:= "Hello, 水"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))
	fmt.Println(string(123456))
}
func HasPrefix(s, prefix string) bool {
	return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func Contains(s, substr string) bool {
	for i :=0; i < len(s); i++ {
		if HasPrefix(s[i:], substr) {
			return true
		}
	}
	return false
}