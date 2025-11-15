// Задание 2: Простой калькулятор
// *Напиши программу, которая принимает два числа и операцию (+, -, , /), затем выводит результат.
// Подсказка: strconv.Atoi() для конвертации строк в числа.

package main

import (
	"errors"
	"fmt"
	"os"
)

func calc() {
	var a, b int
	var result float32
	var operator string
	var err error
	_, err = InputIntNumber(&a)
	_, err = InputOperator(&operator)
	_, err = InputIntNumber(&b)

	switch(operator) {
		case("+"):
			result = float32(a + b)
		case("-"):
			result = float32(a - b)
		case("*"):
			result = float32(a * b)
		case("/"):
			result = float32(a / b)
		default:
			err = errors.New("unknown operator")
	}

	if (err != nil) {
		fmt.Fprintf(os.Stderr, "Calc error: %s\n", err)
	}

	fmt.Printf("Result = %.2f\n", result)
}