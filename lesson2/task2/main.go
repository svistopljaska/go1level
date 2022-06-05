package main

import (
	"fmt"
	"math"
)

func main() {
	var s float64
	fmt.Printf("Введите площадь круга: ")
	fmt.Scanln(&s)

	//Вычислим радиус
	var radius float64
	radius = math.Sqrt(s / math.Pi)

	//Выведем диаметр
	fmt.Printf("Диаметр окружности: %.2f\n", radius*2)

	//Выведем длинну окружности
	fmt.Printf("Длинна окружности: %.2f\n", 2*math.Pi*radius)
}
