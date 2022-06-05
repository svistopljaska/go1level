package main

import "fmt"

func main() {
	var a, b float32
	fmt.Println("Приложение вычисляет площадь прямоугольника")
	fmt.Print("Введите длину первой стороны: ")
	fmt.Scanln(&a)

	fmt.Print("Введите длину второй стороны: ")
	fmt.Scanln(&b)

	fmt.Printf("Площадь прямоугольника = %.2f\n", a*b)
}
