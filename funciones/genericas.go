package funciones

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// CONSTRAINTS

// restriccion de aproximacion
type integer int // creamos un nuevo tipo de dato que almacena ints

// podemos crear restricciones con interfaces
type Numbers interface {
	~int | ~float64 | ~float32 | ~uint
}

func ParamsAndRestrics() {
	fmt.Println(sumaArbitraria(4, 5, 6))
	fmt.Println(sumaUnion(4.5, 5, 6))

	// podemos indicar con que tipo de dato vamos a trabajar (no es necesario)
	fmt.Println(sumaUnion[int](10, 25, 36))
	fmt.Println(sumaUnion[float64](10.32, 2.5, 3.6))

	var num1 integer = 100
	var num2 integer = 2100
	fmt.Println(sumaAprox(num1, num2))

	fmt.Println(sumaDefninida(5.32, 5.15, 5.36))

	fmt.Println(sumaConstraint(5, 10, 15, 5.15, 5.36))

	strings := []string{"a", "b", "c", "d", "e"}
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Println(Includes(strings, "a"))
	fmt.Println(Includes(numbers, 2))

	fmt.Println(Filter(numbers, func(value int) bool { return value > 3 }))
	fmt.Println(Filter(strings, func(value string) bool { return value > "b" }))
}

// restriccion arbitrario
func sumaArbitraria[T int](nums ...T) T {
	// vamos a trabajar con T
	var total T
	for _, num := range nums {
		total += num
	}
	return total
}

// restriccion definida
func sumaDefninida[T Numbers](nums ...T) T {
	var total T
	for _, num := range nums {
		total += num
	}
	return total
}

// restriccion de union
func sumaUnion[T int | float64 | integer](nums ...T) T {
	// vamos a trabajar con T que es de tipo int o flotante o integer
	var total T
	for _, num := range nums {
		total += num
	}
	return total
}

// restriccion de aproxmacion
func sumaAprox[T ~int | ~float64](nums ...T) T {
	// podemos trabajar con integer que deriva de int sin tener que declararlo
	var total T
	for _, num := range nums {
		total += num
	}
	return total
}

// usando constraint desde el paquete
// restriccion con paquete constraint de go
func sumaConstraint[T constraints.Integer | constraints.Float](nums ...T) T {
	var total T
	for _, num := range nums {
		total += num
	}
	return total
}

// Restrccion de operadores
func Includes[T comparable](list []T, value T) bool {
	// comparable es una palabra reservada que compara tipos
	for _, item := range list {
		if item == value {
			return true
		}
	}
	return false
}

func Filter[T constraints.Ordered](list []T, callback func(T) bool) []T {
	result := make([]T, 0, len(list))

	for _, item := range list {
		if callback(item) {
			result = append(result, item)
		}
	}
	return result
}
