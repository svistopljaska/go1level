package main

import "fmt"

//fib calculates the Nth Fibbonacci number
func fib(n int64) int64 {
	if n == 1 || n == 2 {
		return 1
	} else {
		return fib(n-1) + fib(n-2)
	}
}

func main() {
	var n int64

	fmt.Print("Enter which number of fibbonacci must be calculated: ")

	_, err := fmt.Scanf("%d\n", &n)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Result: %d", fib(n))

}
