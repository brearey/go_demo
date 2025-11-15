package main

import (
	"fmt"
	"os"
)

func echo() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("%s ", os.Args[i])
	}
	fmt.Printf("\n")
}

// Задание 1: Утилита echo
// Напиши программу, которая выводит все свои аргументы в одну строку.
