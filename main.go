package main

import (
	"fmt"
	"library/animal"
	"library/book"
	// "library/goroutines"
)

func main() {
	// goroutines.TimeLapse()

	printList("Jhon", "Jane", "Michael", "Brown")
	printList(5, 10, 15, 20, 25, 330)
	printList("Hola mundo", 55, true, 3.1415)
}

func printList(list ...any) {
	for _, value := range list {
		fmt.Println(value)
	}
}

func poo() {
	// No podemos crear el objeto directamente con datos declarados privados
	// var myBook = book.Book{
	// 	Title:  "Moby Dick",
	// 	Author: "Herman Melville",
	// 	Pages:  300,
	// }
	// myBook.PrintInfo()

	myBook2 := book.NewBook("The Process", "Franz Kafka", 255)
	// myBook2.PrintInfo()

	myBook := book.NewBook("MobyDick", "Herman Melville", 640)
	// myBook.PrintInfo()

	myBook.SetTitle("Moby Dick (Special Editon)")
	fmt.Println("\n", myBook.GetTitle())
	// myBook.PrintInfo()

	myTextBook := book.NewTextBook(
		"Computer Networks",
		"Andrew S. Tanenbaum",
		154,
		"DORLING KINDERSLEY (RS)",
		"College")
	// myTextBook.PrintInfo()

	book.Print(myBook)
	book.Print(myBook2)
	book.Print(myTextBook)

	miPerro := animal.Perro{Nombre: "Max"}
	miGato := animal.Gato{Nombre: "Tom"}

	animal.HacerSonido(&miPerro)
	animal.HacerSonido(&miGato)

	animales := []animal.Animal{
		&animal.Perro{Nombre: "Max"},
		&animal.Gato{Nombre: "Tom"},
		&animal.Perro{Nombre: "Buddy"},
		&animal.Gato{Nombre: "Star"},
	}

	for _, animal := range animales {
		animal.Sonido()
	}
}
