package utils

import (
	"fmt"
	"math"
)

func dataTypes() {
	// impresion de caracteres especiales
	fullName := "Antonio Hernandez \t(alias \"kleeC\")\n"
	fmt.Println(fullName)

	// accedemos a la posision del caracter
	var a byte = 'a'
	fmt.Println("byte a", a)
	// accedemos al primer caracter
	s := "hola"
	fmt.Println("Byte de hola[0]: ", s[0]) // se muestra su posicion byte

	// nos imprime su valor de tipo unicode
	var r rune = '❦'
	fmt.Println("unicode: ", r)

	// mostramos las longitudes de los algunos tipos de datos numericos
	fmt.Println("Max Int: ", math.MaxInt)
	fmt.Println("Max Float32: ", math.MaxFloat32)
	fmt.Println("Max Int8: ", math.MaxInt8)

	// TIPO DE DATOS PREDETERMINADOS
	var (
		dafaultInt    int
		defaultUint   uint
		defaultFloat  float32
		defaultBool   bool
		defaultString string
	)
	fmt.Println(
		"Defatults: \n",
		"Int: ", dafaultInt,
		"Uint: ", defaultUint,
		"Float: ", defaultFloat,
		"Bool: ", defaultBool,
		"String", defaultString)
}
