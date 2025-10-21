package main

import "fmt"

func main() {
	var userHeight = 1.8 // meters
	var userKg = 100
	var IMT = userKg / userHeight

	fmt.Printf("IMT = %v\n", IMT)
}