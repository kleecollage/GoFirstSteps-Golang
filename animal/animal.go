package animal

import "fmt"

type Animal interface {
	Sonido()
}

type Perro struct {
	Nombre string
}

func (p Perro) Sonido() {
	fmt.Println(p.Nombre + " hace GUAU GUAU")
}

type Gato struct {
	Nombre string
}

func (g Gato) Sonido() {
	fmt.Println(g.Nombre + " hace MIAU MIAU")
}

// Funcion para imprimir sonido
func HacerSonido(animal Animal) {
	animal.Sonido()
}
