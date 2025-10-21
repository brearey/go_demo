package main

import (
	"fmt"
	"math"
)

func main() {
	userHeight := 1.73 // meters
	var userKg float64 = 75 // kg
	IMT := userKg / math.Pow(userHeight, 2)

	fmt.Printf("IMT = %v\n", IMT)
}