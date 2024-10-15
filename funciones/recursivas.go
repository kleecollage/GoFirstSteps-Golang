package funciones

import "fmt"

func Recursivas() {
	fmt.Println(factorial(5))
}

// RECURIVAS
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
