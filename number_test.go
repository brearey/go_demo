package main

import "testing"

func TestGetRandomNumber(t *testing.T) {
	var max = 5
	var actual = GetRandomNumber(max)

	if actual < 0 {
		t.Error("GetRandomNumber should be greater or equal 0")
	} else if actual > max {
		t.Errorf("GetRandomNumber should be smaller or equal %d number", max)
	}
}
