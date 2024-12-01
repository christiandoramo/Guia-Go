package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func randomFormat() string {
	formats := []string{
		"Oi, %v, Bem-vindo!",
		"Que bom te ver, %v!",
		"Salve, %v! É bom te conhecer",
	}
	return formats[rand.Intn(len(formats))] // retorna uma das strings de 0 até len de formats - 0,1, ou 2
}
func Hello(name string) (string, error) {
	if name == "" {
		return name, errors.New("string passada inválida")
	}
	message := fmt.Sprintf(randomFormat(), name)
	//message := fmt.Sprint(randomFormat())
	return message, nil // significa retorno sem erro
}
func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		msg, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = msg
	}
	return messages, nil // significa retorno sem erro
}
