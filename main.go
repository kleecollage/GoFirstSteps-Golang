package main

import (
	"fmt"

	"github.com/kleecollage/GoFirstSteps-Golang/greetings"
)

func main() {
	fmt.Println("Hola Mundo")
	message, _ := greetings.Hello("Pablo")
	fmt.Println(message)
}
