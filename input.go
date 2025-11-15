package main

import "fmt"

func InputString(str *string) (int, error) {
	fmt.Println("Please enter string")
	return fmt.Scanf("%s", str)
}

func InputIntNumber(num *int) (int, error) {
	fmt.Println("Please enter any integer number")
	return fmt.Scanf("%d", num)
}

func InputOperator(operator *string) (int, error) {
	fmt.Println("Please enter operator: +, -, *, /")
	return fmt.Scanf("%s", operator)
}
