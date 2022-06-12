package main

import (
	"bufio"
	"fmt"
	"os"
)

var numbers []int64

func main() {
	br := bufio.NewReader(os.Stdin)
	s, err := br.ReadString('\n')
	if err != nil {
		panic(err)
	}
	var d int64
	_, err = fmt.Sscanf(s, "%d", &d)
	if err != nil {
		panic(err)
	}

	var i int64
	var isSimple bool
	for i = 2; i < d; i++ {
		isSimple = true
		for _, num := range numbers {
			if i%num == 0 {
				isSimple = false
				break
			}
		}
		if isSimple {
			numbers = append(numbers, i)
		}
	}

	fmt.Println("Простые числа 1 < s < N : ", numbers)
}
