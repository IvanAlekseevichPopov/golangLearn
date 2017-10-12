package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("asdfasd")
	fmt.Println(basename2("asdfaasdf/34254325234.df"))
}

func basename2(s string) string {
	slash := strings.LastIndex(s, "/") // -1, если / не найден
	s = s[slash+1:]

	if dotPosition := strings.LastIndex(s, "."); dotPosition >= 0 {
		s = s[:dotPosition]
	}

	return s
}