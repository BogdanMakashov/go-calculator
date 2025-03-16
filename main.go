package main

import (
	"errors"
	"fmt"

	"github.com/fatih/color"
)

func main() {
	calculateAgain := true

	for {
		if calculateAgain {
			runCalculation()
			calculateAgain = askToCalculateAgain()
		} else {
			break
		}
	}

}

func askToCalculateAgain() bool {
	var calculateAgain string

	fmt.Println("Произвести ещё одно вычисление?")
	fmt.Scan(&calculateAgain)

	return calculateAgain == "Да"
}

func add(x, y float32) float32 {
	return x + y
}

func subtract(x, y float32) float32 {
	return x - y
}

func multiply(x, y float32) float32 {
	return x * y
}

func divide(x, y float32) (float32, error) {
	if y == 0 {
		return 0, errors.New("Делить на ноль нельзя!")
	}
	return x / y, nil
}

func runCalculation() {
	var x, y float32
	var operation string

	fmt.Println("Введите левый операнд:")
	fmt.Scan(&x)
	fmt.Println("Введите правый операнд:")
	fmt.Scan(&y)
	fmt.Println("Введите требуемую операцию(+,-,*,/):")
	fmt.Scan(&operation)

	switch operation {
	case "+":
		fmt.Println("Результат: ", add(x, y))
	case "-":
		fmt.Println("Результат: ", subtract(x, y))
	case "*":
		fmt.Println("Результат: ", multiply(x, y))
	case "/":
		result, error := divide(x, y)

		if error != nil {
			color.Red(error.Error())
		}

		fmt.Println("Результат: ", result)
	default:
		fmt.Println("Недопустимый оператор")
	}
}
