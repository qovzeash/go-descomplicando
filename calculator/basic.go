package calculator

import "fmt"

func Basic() {
	var numberOne, numberTwo int
	var operation string

	fmt.Scanf("%d %s %d", &numberOne, &operation, &numberTwo)

	switch {
	case operation == "+":
		fmt.Println(numberOne + numberTwo)
	case operation == "-":
		fmt.Println(numberOne - numberTwo)
	case operation == "*":
		fmt.Println(numberOne * numberTwo)
	case operation == "/":
		fmt.Println(numberOne / numberTwo)
	default:
		fmt.Println("The operation should be written as follows: 12 + 1, or 13 - 2...")

	}
}
