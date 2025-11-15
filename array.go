package main

import "fmt"

func printArray(arr [3]int) {
	for idx, value := range arr {
		fmt.Printf("%d) %d\n", idx, value)
	}
}
