package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ") // prefixo representando a função sendo usada abaixo
	log.SetFlags(0)

	names := []string{"Christian", "Cristina", "Paulo", "André"}

	messages, error := greetings.Hellos(names)
	if error != nil { // se diferente de erro nulo (nil == retorno normal)
		log.Fatal(error)
	}
	for _, msg := range messages {
		fmt.Println(msg)
	}
}
