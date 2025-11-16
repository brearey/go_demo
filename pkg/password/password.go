package password

import (
	"errors"
	"math/rand"
)

func Generate(passLength int) (string, error) {
	if passLength < 4 {
		return "", errors.New("длина пароля должна быть больше или равно 4")
	}

	const chars = "abcdefghABCDEFGH!@#$%^&*()1234567890"
	var password string = ""

	for i := 0; i < passLength; i++ {
		var randomSymbol = chars[rand.Intn(len(chars))]
		password += string(randomSymbol)
	}
	return password, nil
}
