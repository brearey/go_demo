package main

import "fmt"

func main() {
	// for i := 1; i < len(os.Args); i++ {
	// 	fmt.Printf("%d) arg = %v\n", i, os.Args[i])
	// }
	var arr [3]int = [3]int{1, 2, 3}
	printArray(arr)

	var name string
	var _, err = inputName(&name)
	if (err != nil) {
		fmt.Printf("Error: %s\n", err)
	}
	hello(name)

	var age int
	var _, err2 = inputIntNumber(&age)
	if (err2 != nil) {
		fmt.Printf("Error: %s\n", err2)
	} else {
		fmt.Printf("Your age = %d\n", age);
	}
}