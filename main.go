package main

import (
	"fmt"
)

func main() {
	// for i := 1; i < len(os.Args); i++ {
	// 	fmt.Printf("%d) arg = %v\n", i, os.Args[i])
	// }
	var arr [3]int = [3]int{1, 2, 3}
	for idx, value := range arr {
		fmt.Printf("%d) %d\n", idx, value)
	}
	echo()
}