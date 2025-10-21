package main

import (
	"fmt"
	"math"
)

func main() {
	userHeight := 1.73
	userKg := 75.0
	IMT := userKg / math.Pow(userHeight, 2)

	fmt.Printf("IMT = %v\n", IMT)
}