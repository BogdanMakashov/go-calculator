package main

import "fmt"

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

func add(x, y float64) float64 {
	return x + y
}

func subtract(x, y float64) float64 {
	return x - y
}

func multiply(x, y float64) float64 {
	return x * y
}

func divide(x, y float64) float64 {
	return x / y
}

func runCalculation() {
	var x, y float64
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
		fmt.Println("Результат: ", divide(x, y))
	default:
		fmt.Println("Недопустимый оператор")
	}
}
