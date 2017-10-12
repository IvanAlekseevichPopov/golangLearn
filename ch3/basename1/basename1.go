package main

import (
	"fmt"
)

func main() {
	fmt.Println("asdfasd")
	fmt.Println(basename("asdfaasdf/34254325234.df"))
}

func basename(s string) string {
	//Отбрасываем последний символ '/' и все перед ним
	for i:= len(s) -1 ; i>= 0; i-- {
		if s[i] == '/' {
			s = s[i+1:]
			break
		}
	}
	//Сохраняем все до последней точки
	for i := len(s) -1 ; i >=0; i-- {
		if s[i] == '.' {
			s= s[:i]
			break
		}
	}

	return s
}