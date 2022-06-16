package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func insertionSort(sl []int64) {
	slLen := len(sl)
	for j := 1; j < slLen; j++ {
		key := sl[j]
		i := j - 1
		for ; (i > -1) && (sl[i] > key); i-- {
			sl[i+1] = sl[i]
		}
		sl[i+1] = key
	}
}

func main() {

	fmt.Println("Введите целые числа, разделяя их нажатием Enter")
	fmt.Println("После ввода последнего символа нажмите Ctrl+Z")

	var scanner = bufio.NewScanner(os.Stdin)
	var slice = make([]int64, 0)

	for scanner.Scan() {
		line := scanner.Text()
		intNumber, err := strconv.ParseInt(line, 10, 64)
		if err != nil {
			panic(err)
		}

		slice = append(slice, intNumber)
	}

	fmt.Println(slice)
	insertionSort(slice)
	fmt.Println(slice)
}
