package main

import (
	"bytes"
	"fmt"
	"strings"
	"time"
)

const GROUP int = 3

func main() {
	t0 := time.Now()
	fmt.Println(comma("123461324123123411362347243652345234523452345234234723476234523"))
	t1 := time.Now()
	fmt.Printf("Time - %v\n", t1.Sub(t0))
	fmt.Println(comma2("12346132312341136234724334762.345"))
	t2 := time.Now()
	fmt.Printf("Time - %v\n", t2.Sub(t1))
}

func comma(s string) string {
	n := len(s)
	if n <= GROUP {
		return s
	}

	return comma(s[:n-GROUP]) + "," + s[n-GROUP:]
}

func comma2(s string) string {
	var buf bytes.Buffer
	n := len(s)
	endOfString := ""
	dotPosition := strings.LastIndex(s, ".")
	if dotPosition >= 0 {
		n = dotPosition
		endOfString = s[dotPosition:]
	}

	startPosition := n % GROUP

	if 0 == startPosition {
		buf.WriteString(s[startPosition:GROUP])
		startPosition = GROUP
	} else {
		buf.WriteString(s[:startPosition])
	}

	for i := startPosition; i < n; i += GROUP {
		buf.WriteString("," + s[i:i+GROUP])
	}

	buf.WriteString(endOfString)

	return buf.String()
}
