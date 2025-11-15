package main

import (
	"math/rand"
)

func getRandomNumber(max int) int {
	return rand.Intn(max)
}