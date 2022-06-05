package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func factorial(n float64) float64 {
	n = math.Abs(math.Round(n))
	res := n
	for i := res - 1; i > 0; i-- {
		res *= i
	}
	return res
}

func main() {
	var a, b, res float64
	var op string
	fmt.Println("Добро пожаловать в приложение Калькулятор")
	for {
		fmt.Println("Введите входные данные в формате: <дробное число> <операция> <дробное число>")
		fmt.Println("Операция = +|-|/|*|**|!")
		fmt.Println("Для выхода наберите quit")

		br := bufio.NewReader(os.Stdin)
		s, err := br.ReadString('\n')
		if err != nil {
			panic(err)
		}
		if s == "quit\r\n" {
			break
		} else {
			_, err := fmt.Sscanf(s, "%f %s %f", &a, &op, &b)
			if err != nil {
				return
			}
		}

		switch op {
		case "+":
			res = a + b
		case "-":
			res = a - b
		case "*":
			res = a * b
		case "/":
			if b == 0 {
				fmt.Println("Делить на 0 нельзя!")
				continue
			} else {
				res = a / b
			}
		case "**":
			res = math.Pow(a, b)
		case "!":
			res = factorial(a)
		default:
			continue
		}
		fmt.Printf("Результат: %.2f\n\n", res)
	}
}
