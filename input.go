package main

import "fmt"

func inputName(name *string) (int, error)  {
	fmt.Println("Please enter your name")
	return fmt.Scanf("%s", name)
}

func inputIntNumber(num *int) (int, error) {
	fmt.Println("Please enter any integer number")
	return fmt.Scanf("%d", num)
}