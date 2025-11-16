package main

import "testing"

func TestCalculate(t *testing.T) {
	var a float32 = 2
	var b float32 = 3
	var operator = "+"
	var except = a + b
	var actual, err = calculate(a, operator, b)

	if err != nil {
		t.Errorf("calculate выбросил ошибку %s", err)
	}

	if actual != except {
		t.Error("ответ calculate не совпадает с ожидаемым")
	}
}
