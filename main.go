package main

import (
	"fmt"
	"log"

	"github.com/kleecollage/GoFirstSteps-Golang/greetings"
)

func main() {
	log.SetPrefix("greetings(): ")
	log.SetFlags(0) // las banderas de formato determinan que info adicional se incluyen en el regitro

	names := []string{"Alex", "Jhon", "Jane"}
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err) // manejador de errores
	}
	fmt.Println(messages)

	message, err := greetings.Hello("Juan")
	// manejador de errores
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(message)
}
