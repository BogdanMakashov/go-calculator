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

func add(x, y int) int {
	return x + y
}

func subtract(x, y int) int {
	return x - y
}

func multiply(x, y int) int {
	return x * y
}

func divide(x, y int) int {
	return x / y
}

func runCalculation() {
	var x, y int
	var operation string

	fmt.Println("Введите левый операнд:")
	fmt.Scan(&x)
	fmt.Println("Введите правый операнд:")
	fmt.Scan(&y)
	fmt.Println("Введите требуемую операцию(+,-,*,/):")
	fmt.Scan(&operation)

	switch operation {
	case "+":
		fmt.Println(add(x, y))
	case "-":
		fmt.Println(subtract(x, y))
	case "*":
		fmt.Println(multiply(x, y))
	case "/":
		fmt.Println(divide(x, y))
	default:
		fmt.Println("Недопустимый оператор")
	}
}
