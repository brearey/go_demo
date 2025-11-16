package password

import (
	"reflect"
	"testing"
)

func TestGeneratePassword(t *testing.T) {
	const passLength = 10
	var password, err = Generate(passLength)

	if err != nil {
		t.Error("функция Generate не должна вернуть ошибку")
	}

	if len(password) != passLength {
		t.Errorf("длина пароля ожидалась %d, но получили %d", passLength, len(password))
	}

	if reflect.TypeOf(password).Kind() != reflect.String {
		t.Errorf("ожидалась строка но получили %v", reflect.TypeOf(password))
	}
}
