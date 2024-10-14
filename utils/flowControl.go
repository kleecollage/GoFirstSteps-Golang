package utils

import (
	"fmt"
	"runtime"
	"time"
)

func flows() {
	t := time.Now()
	hora := t.Hour()

	// en Go podemos inicializar variables dentro del if...
	// if t := time.Now(); t.Hour() < 12 {
	if hora < 12 {
		fmt.Println("¡Mañana!")
	} else if hora < 18 {
		fmt.Println("¡Tarde!")
	} else {
		fmt.Println("¡Noche!")
	}

	// en Go podemos inicializar variables dentro del switch...
	switch t := time.Now(); {
	case t.Hour() < 12:
		fmt.Println("¡Mañana!")
	case t.Hour() < 18:
		fmt.Println("¡Tarde!")
	default:
		fmt.Println("Noche!")
	}

	os := runtime.GOOS
	switch os {
	// switch os := runtime.GOOS; os {
	case "windows":
		fmt.Println("Go run -> Windows")
		// break no es necesario en Go
	case "linux":
		fmt.Println("Go run -> Linux")
	case "darwin":
		fmt.Println("Go run -> macOs")
	default:
		fmt.Println("Go run -> otro Os")
	}

	// BUCLE FOR
	var i int
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	for j := 9; j >= 0; j-- {
		if j == 5 {
			continue // salta a la siguiente iteracion
			// break    // fin del bucle
		}
		fmt.Println(j)
	}

	// FUNCIONES ...
	myFunc()

	saludo := hello("KleeC")
	fmt.Println(saludo)

	sum, mul := calc(2, 4)
	fmt.Println("La suma es: ", sum)
	fmt.Println("La multiplicacion es: ", mul)
}

// FUNCIONES
// sin retorno
func myFunc() {
	fmt.Println("Hola desde la funcion myFunc()")
}

// con retorno
func hello(name string) string {
	return "Hola, " + name
}

// con multiples valores de retorno
func calc(a, b int) (sum, mul int) {
	// sum := a + b // si se declaran las variable como valor de retorno
	// mul := a * b // unicamente se asignan sus valores
	sum = a + b
	mul = a * b

	return // sum, mul // no hace falta declarar los valores de retorno si se delcaran en la funcion
}
