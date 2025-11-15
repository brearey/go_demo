// Задание 2: Простой калькулятор
// *Напиши программу, которая принимает два числа и операцию (+, -, , /), затем выводит результат.

package main

import (
	"errors"
	"fmt"
)

func calc() {
	var a, b int
	var result float32
	var operator string
	var err error
	_, err = InputIntNumber(&a)
	if err != nil {
		PrintError("Calc", err)
		return
	}
	_, err = InputOperator(&operator)
	if err != nil {
		PrintError("Calc", err)
		return
	}
	_, err = InputIntNumber(&b)
	if err != nil {
		PrintError("Calc", err)
		return
	}

	switch operator {
	case ("+"):
		result = float32(a + b)
	case ("-"):
		result = float32(a - b)
	case ("*"):
		result = float32(a * b)
	case ("/"):
		if b == 0 {
			err = errors.New("div by zero")
			PrintError("Calc", err)
			return
		}
		result = float32(a) / float32(b)
	default:
		PrintError("Calc", errors.New("unknown operator"))
		return
	}

	fmt.Printf("Result = %.2f\n", result)
}
