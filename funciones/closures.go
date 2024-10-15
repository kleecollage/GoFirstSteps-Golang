package funciones

import "fmt"

func Closure() {
	nextInt := incrementer() // guarda el estado del resultado
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	// se mantienen las referencias a variables fuera de su propio alcance lexixo
	fmt.Println(nextInt())
}

// FUNCIONES CLOSURE
func incrementer() func() int {
	i := 0
	// significa tener una funcion anonima dentro de otra
	return func() int {
		i++
		return i
	}
}
