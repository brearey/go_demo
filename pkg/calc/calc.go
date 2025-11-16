package calc

import (
	"demo/app-1/pkg/input"
	"demo/app-1/pkg/util"
	"errors"
	"fmt"
)

func Calc() {
	var a, b int
	var result float32
	var operator string
	var err error
	_, err = input.InputIntNumber(&a)
	if err != nil {
		util.PrintError("Calc", err)
		return
	}
	_, err = input.InputOperator(&operator)
	if err != nil {
		util.PrintError("Calc", err)
		return
	}
	_, err = input.InputIntNumber(&b)
	if err != nil {
		util.PrintError("Calc", err)
		return
	}

	result, err = calculate(float32(a), operator, float32(b))
	if err != nil {
		util.PrintError("Calc", err)
		return
	}

	fmt.Printf("Result = %.2f\n", result)
}

func calculate(a float32, operator string, b float32) (float32, error) {
	var result float32
	switch operator {
	case ("+"):
		result = a + b
	case ("-"):
		result = a - b
	case ("*"):
		result = a * b
	case ("/"):
		if b == 0 {
			return 0, errors.New("div by zero")
		}
		result = a / b
	default:
		return 0, errors.New("unknown operator")
	}
	return result, nil
}
