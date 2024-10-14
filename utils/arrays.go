package utils

import "fmt"

func arrays() {
	// ... significa que no sabemos el numero de elementos que va a haber en el array
	var a = [...]int{5, 10, 30, 40, 50}
	fmt.Println(a)

	a[0] = 10
	a[1] = 20

	for i := 0; i < len(a); i++ {
		fmt.Print(a[i])
	}

	for index, value := range a {
		fmt.Printf("\nIndice: %d, Valor: %d", index, value)
	}

	// matriz de dos dimensiones
	var matriz = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 8}}
	fmt.Println(matriz)
}

// SLICE
func slices() {
	diasSemana := []string{"Domingo", "Lunes", "Martes", "Miercoles", "Jueves", "Viernes", "Sabado"}
	fmt.Println("Array diasSemana", diasSemana)

	diaRebanada := diasSemana[0:5]
	fmt.Println("Array diaRebanada", diaRebanada) // Mostramos los elementos del indice 0 al indice 4

	fmt.Println("Longitud: ", len(diaRebanada))
	fmt.Println("Capacidad: ", cap(diaRebanada)) // capacidad de elementos que puede almacenar "diaRebanada"

	diaRebanada = append(diaRebanada, "Julio")
	diaRebanada = append(diaRebanada, "Marzo")
	fmt.Println(diaRebanada)
	fmt.Println("Longitud: ", len(diaRebanada))
	fmt.Println("Capacidad: ", cap(diaRebanada))
	// Al agregar otro elemento superariamos la maxima capacidad, entonces su capacidad se hace dinamicamente mas grande
	diaRebanada = append(diaRebanada, "Octubre")
	fmt.Println(diaRebanada)
	fmt.Println("Longitud: ", len(diaRebanada))
	fmt.Println("Capacidad: ", cap(diaRebanada))
	// Vamos a eliminar el elemento del indice 2 ...
	diaRebanada = append(diaRebanada[:2], diaRebanada[3:]...) // si iniciamos con 0:n no es necesario escribir el 0
	fmt.Println(diaRebanada)
	fmt.Println("Longitud: ", len(diaRebanada))
	fmt.Println("Capacidad: ", cap(diaRebanada)) // observamos que la capacidad no disminuye una vez se ha incrementado
}

func funcionMake() {
	nombres := make([]string, 5, 10) // indicamos el tipo, longitud y capacidad del slice

	nombres[0] = "Alex"
	fmt.Println(nombres)
	fmt.Println("Longitud: ", len(nombres))
	fmt.Println("Capacidad: ", cap(nombres))

	rebanada1 := []int{1, 2, 3, 4, 5}
	rebanada2 := make([]int, 5)
	// copiamos el contenido de rebanada2 a rebanada1
	copy(rebanada2, rebanada1)
	fmt.Println(rebanada1)
	fmt.Println(rebanada2)
}

func Mapas() {
	// map [clave]valor
	colors := map[string]string{
		"rojo":  "#F00",
		"verde": "#0F0",
		"azul":  "#00F",
	}

	fmt.Println("Colors: ", colors)
	fmt.Println(colors["rojo"])

	colors["negro"] = "#000"
	fmt.Println(colors)

	valor, ok := colors["azul"] // devolvemos el valor y un booleano
	fmt.Println(valor, ok)

	if valor, ok := colors["blanco"]; ok {
		fmt.Println("Si existe la clave: ", valor)
	} else {
		fmt.Println("No existe esta clave")
	}

	delete(colors, "azul")
	fmt.Println("delte (colors, 'azul'): ", colors)

	for clave, valor := range colors {
		fmt.Printf("Clave: %s, Valor: %s\n", clave, valor)
	}

}
