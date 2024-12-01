package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("string passada inválida")
	}
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil // significa retorno sem erro
}
