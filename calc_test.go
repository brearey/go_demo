package main

import "testing"

func TestCalculateAdd(t *testing.T) {
	var a float32 = 2
	var b float32 = 3
	var operator = "+"
	var except = a + b
	var actual, err = calculate(a, operator, b)

	if err != nil {
		t.Errorf("calculate выбросил ошибку %s", err)
	}

	if actual != except {
		t.Errorf("ответ calculate %.6f не совпадает с ожидаемым %.6f", actual, except)
	}
}

func TestCalculateSub(t *testing.T) {
	var a float32 = 3333.22313
	var b float32 = 1000000.330292
	var operator = "-"
	var except = a - b
	var actual, err = calculate(a, operator, b)

	if err != nil {
		t.Errorf("calculate выбросил ошибку %s", err)
	}

	if actual != except {
		t.Errorf("ответ calculate %.6f не совпадает с ожидаемым %.6f", actual, except)
	}
}

func TestCalculateMulti(t *testing.T) {
	var a float32 = 24.05
	var b float32 = 34223.1023
	var operator = "*"
	var except = a * b
	var actual, err = calculate(a, operator, b)

	if err != nil {
		t.Errorf("calculate выбросил ошибку %s", err)
	}

	if actual != except {
		t.Errorf("ответ calculate %.6f не совпадает с ожидаемым %.6f", actual, except)
	}
}

func TestCalculateDiv(t *testing.T) {
	var a float32 = -2323123.2323
	var b float32 = -23223.023232
	var operator = "/"
	var except = a / b
	var actual, err = calculate(a, operator, b)

	if err != nil {
		t.Errorf("calculate выбросил ошибку %s", err)
	}

	if actual != except {
		t.Errorf("ответ calculate %.6f не совпадает с ожидаемым %.6f", actual, except)
	}
}