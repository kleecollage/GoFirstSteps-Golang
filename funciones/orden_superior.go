package funciones

import "fmt"

func OrdSuperior() {
	r := double(addOne, 3)
	fmt.Println("Resultado: ", r)
}

// FUNCIONES DE ORDEN SUPERIOR
func double(f func(int) int, x int) int {
	return f(x * 2)
}

func addOne(x int) int {
	return x + 1
}
