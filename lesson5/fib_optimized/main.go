package main

import "fmt"

//Cache keeps calculated fibs for speeding up the calculation
type Cache map[int64]int64

var cache Cache

//fib calculates the Nth Fibbonacci number
func fib(n int64) int64 {
	if n == 1 || n == 2 {
		return 1
	}
	val, ok := cache[n]
	if ok {
		return val
	}

	val = fib(n-1) + fib(n-2)
	cache[n] = val
	return val

}

func main() {
	var n int64
	cache = make(Cache)

	fmt.Print("Enter which number of fibbonacci must be calculated: ")

	_, err := fmt.Scanf("%d\n", &n)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Result: %d", fib(n))

}
