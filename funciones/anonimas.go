package funciones

import "fmt"

func Anonimas() {
	// FUNCION ANONIMA
	func() {
		fmt.Println("Hola soy una funcion anonima")
	}() // debemos poner parentesis al final para poder llamarla

	// variable que contiene una funcion anonima
	saludo := func(name string) {
		fmt.Printf("¡Hola, %s!\n", name)
	}
	saludar("Alex", saludo)

	// pasamos funciones como valor
	r1 := aplicar(duplicar, 5)
	r2 := aplicar(triplicar, 5)

	fmt.Println(r1, r2)
}

func saludar(name string, f func(string)) {
	f(name)
}

func duplicar(n int) int {
	return n * 2
}

func triplicar(n int) int {
	return n * 3
}

func aplicar(f func(int) int, n int) int {
	return f(n)
}
