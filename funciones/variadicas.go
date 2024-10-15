package funciones

import "fmt"

func Variadicas() {
	fmt.Println(suma("Alex", 12, 5, 78, 54))
	fmt.Println(suma("Roel", 10, 20, 30, 40))

	imprimirDatos("Hola", 28, true, 3.14)
	printList("Hola mundo", 55, true, 3.1415)
}

// FUNCION VARIADICA
// recibimos una cantidad variable de argumentos del mismo tipo
func suma(name string, nums ...int) int {
	// fmt.Printf("%T - %v\n", nums, nums)
	var total int
	for _, num := range nums {
		total += num
	}

	fmt.Printf("Hola %s, la suma es: %d\n", name, total)
	return total
}

// recibimos una cantidad variable de argumentos de distinto tipo
func imprimirDatos(datos ...interface{}) {
	for _, dato := range datos {
		fmt.Printf("%T - %v\n", dato, dato)
	}
}

// otra forma de treer valores de distinto tipo sin usar interface{}
func printList(list ...any) {
	for _, value := range list {
		fmt.Println(value)
	}
}
