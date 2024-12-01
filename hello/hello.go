package main

import (
	"fmt"
	"hello/utils"
)

func main() {
	utils.ReusingQuoteGo()

	const name string = "Christian"
	msg := fmt.Sprintf("Olá, %s!", name)

	fmt.Println(msg)
}
