package main

import (
	"fmt"
)

func main() {
	var number int
	fmt.Print("Введите трехзначное число: ")
	fmt.Scanln(&number)
	if number < 0 {
		number = -number
	}
	if number < 100 || number > 999 {
		fmt.Print("Требуется ввести трехзначное число!")
		return
	}

	fmt.Printf("Сотни: %d\n", number/100)
	fmt.Printf("Десятки: %d\n", (number/10)%10)
	fmt.Printf("Единицы: %d\n", number%10)
}
