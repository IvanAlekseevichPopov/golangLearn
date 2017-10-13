package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(anagram("123555", "255351"))
}

func anagram(str1 string, str2 string) bool {
	if(len(str1) != len(str2)) {
		return false
	}
	var char = []rune(str2)

	for _, v := range char  {
		if !strings.ContainsRune(str1, v) {
			return false
		}
		str1 = strings.Replace(str1, string(v), "", 1)
	}

	return true
}
