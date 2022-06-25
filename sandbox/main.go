package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "1234567890"
	b := make([]byte, 3)
	sr := strings.NewReader(s)
	for {
		n, e := sr.Read(b)
		fmt.Println(n, b, e)
		if e != nil {
			break
		}
	}
}
