package utils

import "fmt"

type Persona struct {
	nombre string
	edad   int
	correo string
}

// Go recomienda trabajar con punteros
// Creamos un metodo a la estructura Persona
func (p *Persona) diHola() {
	fmt.Println("Hola mi nombre es", p.nombre)
}

func myStructure() {
	// instanciamos la estructura de tipo persona
	var p Persona
	fmt.Println(p) // persona con vlores predeterminados
	// asignamos valores a la estructura
	p.nombre = "Alex"
	p.edad = 23
	p.correo = "alex@google.com"
	fmt.Println(p)
	p.diHola() // Accedemos al metodo de la estructura

	// podemos crear instancias en una sola linea:
	p2 := Persona{"Antonio", 28, "antonio@google.com"} // segunda instancia
	fmt.Println(p2)
	fmt.Println(p2.correo)
	p2.edad = 30
	fmt.Println(p2)
}

func punteros() {
	var x int = 10
	fmt.Println("Valor de x", x)
	var p *int = &x // mandamos la referencia de la memoria de x a p
	fmt.Println("Valor del puntero de x: ", p)
	fmt.Println("Referencia a la memoria de x: ", &x)

	editar(&x)
	fmt.Println("Valor de x", x)
}

// recibimos un puntero en la variable x
func editar(x *int) {
	*x = 20
}
