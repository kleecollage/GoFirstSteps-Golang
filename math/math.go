package math

import (
	"fmt"
	"math"
)

func mathOperations() {
	a := 10
	b := 3

	fmt.Println("Suma: ", a+b)
	fmt.Println("Resta ", a-b)
	fmt.Println("Division: ", a/b)
	fmt.Println("Residuo: ", a%b)

	b++
	fmt.Println("Incremento: ", b)
	b--
	fmt.Println("Decremento: ", b)

	a += 5
	fmt.Println("Suma en asignacion: ", a)

	fmt.Println("Constantes de math (Pi): ", math.Pi)
	fmt.Println("Constantes de math (e): ", math.E)

	c := math.Pow(2, 3)
	fmt.Println("Pow: ", c)

	fmt.Println("Sqrt: ", math.Sqrt(64))
	fmt.Println("Cbrt: ", math.Cbrt(64))
}

func trianglePermieter() {
	var lado1, lado2 float32
	const precision = 2

	fmt.Print("Lado 1: ")
	fmt.Scanln(&lado1)
	fmt.Print("Lado 2: ")
	fmt.Scanln(&lado2)

	area := (lado1 * lado2) / 2
	hipotenusa := math.Sqrt(math.Pow(float64(lado1), 2) + math.Pow(float64(lado2), 2))
	perimetro := lado1 + lado2 + float32(hipotenusa)

	fmt.Printf("\nArea: %.*f ", precision, area)
	fmt.Printf("\nPerimetro: %.*f", precision, perimetro)
}
