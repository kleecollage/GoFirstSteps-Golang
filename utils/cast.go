package utils

import (
	"fmt"
	"strconv"
)

func casting() {
	var integer16 int16 = 50
	var integer32 int32 = 100
	// casteo de datos
	fmt.Println(int32(integer16) + integer32)

	// casting de string a int con strconv de GO
	s := "150"
	i, _ := strconv.Atoi(s) // Atoi devuelve un int y un error, el error no lo manejamos, por eso _
	fmt.Println(i + i)      // suma

	// casting de int a string
	n := 32
	s = strconv.Itoa(n) // Iota devuelve un string
	fmt.Println(s + s)  // concatena
}
