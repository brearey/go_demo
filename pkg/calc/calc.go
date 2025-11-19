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
	if _, err = input.InputIntNumber(&a); err != nil {
		util.PrintError("Calc", err)
		return
	}
	if _, err = input.InputOperator(&operator); err != nil {
		util.PrintError("Calc", err)
		return
	}
	if _, err = input.InputIntNumber(&b); err != nil {
		util.PrintError("Calc", err)
		return
	}

	if result, err = calculate(float32(a), operator, float32(b)); err != nil {
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
