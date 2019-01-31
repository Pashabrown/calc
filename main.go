package main

// Here's a change

import (
	"fmt"
	"math"
)

func Add(a int, b int) int {
	return a + b
}

func Substract(a int, b int) int {
	return a - b
}

func Multiply(a int, b int) int {
	return a * b
}

func Divide(a int, b int) int {
	if b == 0 {
		fmt.Println("You cant divide by 0")
		return 0
	}

	return a / b
}

func Raise(a int, b int) float64 {
	return math.Pow(float64(a), float64(b))
}
func main() {

	for {

		var num1 int
		var num2 int
		var operator string
		var result int

		fmt.Scanf("%d %s %d", &num1, &operator, &num2)
		//fmt.Println(num2, operator, num1)

		switch operator {
		case "+":
			result = Add(num1, num2)
		case "-":
			result = Substract(num1, num2)
		case "*":
			result = Multiply(num1, num2)
		case "/":
			result = Divide(num1, num2)
		case "^":
			result = int(Raise(num1, num2))

		default:
			fmt.Println("Unknown operator")

		}

		fmt.Println(result)
		//sum := Add(3, 2)
		//fmt.Println(sum)
	}
}
