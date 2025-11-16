package number

import (
	"math/rand"
)

func GetRandomNumber(max int) int {
	return rand.Intn(max)
}
