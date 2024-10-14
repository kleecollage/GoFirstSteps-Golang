package utils

import (
	"fmt"

	"rsc.io/quote"
)

// declarar variables en una sola linea a nivel de paquete
var console, game, version = "Xbox", "WarHammer", 12

// declaracion de constantes a nivel de paquete
// se recomienda declrar constantes con letra Capital
const Pi = 3.1415
const (
	x = 100
	y = 0b1010 // binario
	z = 0o12   // octal
	w = 0xff   // hexadecimal
)

// valores iota
const (
	Domingo   = iota + 1 // 0 + 1
	Lunes                // 1 + 1
	Martes               // 2 + 1
	Miercoles            // iota + 1 ...
	Jueves
	Viernes
	Sabado
)

func myVariables() {
	fmt.Println(quote.Hello())
	fmt.Println(quote.Go())

	// variables
	var firstName, lastName string
	var age int
	// declarar variables en grupo
	var (
		name    = "Klee"
		surname = "Collage"
		year    = 2024
	)
	// declarar variables con := unicamente se adminte esta forma dentro de las funciones
	system, model := "Android", "Cubot"
	stock := 300
	stock = 235

	firstName = "Antonio"
	lastName = "Hernandez"
	age = 37

	fmt.Println(firstName, lastName, age)
	fmt.Println(system, model, stock)
	fmt.Println(name, surname, year)
	fmt.Println(Pi)
	fmt.Println(x, y, z, w)
	fmt.Println(Domingo)
}
