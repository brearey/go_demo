package main

import (
	"fmt"
	"math"
)

func main() {
	var userHeight = 1.73 // meters
	var userKg float64 = 75.0
	var IMT = userKg / math.Pow(userHeight, 2)

	fmt.Printf("IMT = %v\n", IMT)
}