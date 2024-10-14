package utils

import "fmt"

func fmtUses() {
	name := "Antonio"
	age := 32
	// %s trae strings %d trae ints
	fmt.Printf("Hola, me llamo %s y tengo %d años.\n", name, age)
	// %v trae un valor default
	greeting := fmt.Sprintf("Hola, me llamo %v y tengo %v años", name, age)
	fmt.Println(greeting)

	fmt.Printf("El tipo de name es: %T \n", name)
	fmt.Printf("El tipo de age es: %T \n", age)

	var alias string
	var alias2 string
	var year int
	fmt.Print("Ingrese su alias: ")
	fmt.Scanln(&alias, &alias2)
	fmt.Print("Ingrese su generacion: ")
	fmt.Scanln(&year)
	fmt.Printf("Hola, mi alias es %s %s y soy de la generacion %d.\n", alias, alias2, year)
}
