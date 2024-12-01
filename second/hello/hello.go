package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ") // prefixo representando a função sendo usada abaixo
	log.SetFlags(0)

	message, error := greetings.Hello("")
	if error != nil { // se diferente de erro nulo (nil == retorno normal)
		log.Fatal(error)
	}
	fmt.Println(message)
}
