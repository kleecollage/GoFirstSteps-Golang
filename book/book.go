package book

import "fmt"

type Printable interface {
	PrintInfo()
}

func Print(p Printable) {
	p.PrintInfo()
}

type Book struct {
	// declarar las variables en minusculas las convierte en privadas
	title  string
	author string
	pages  int
}

// empezar con letra capital hace las funciones publicas y accesibles desde otros puntos de la app
func NewBook(title string, author string, pages int) *Book {
	return &Book{
		title:  title,
		author: author,
		pages:  pages,
	}
}

// GETTERS AND SETTER para estructura con datos privados
func (b *Book) SetTitle(title string) {
	b.title = title
}
func (b *Book) GetTitle() string {
	return b.title
}

func (b *Book) SetAuthor(author string) {
	b.author = author
}
func (b *Book) GetAuthor() string {
	return b.author
}

func (b *Book) SetPages(pages int) {
	b.pages = pages
}
func (b *Book) GetPages() int {
	return b.pages
}

// Imprimir la informacion
func (b *Book) PrintInfo() {
	fmt.Printf("\nTitle: %s\nAuthor: %s\nPages: %d\n", b.title, b.author, b.pages)
}

type TextBook struct {
	Book
	editorial string
	level     string
}

// Estructura con herencia
func NewTextBook(title string, author string, pages int, editorial string, level string) *TextBook {
	return &TextBook{
		Book:      Book{title, author, pages},
		editorial: editorial,
		level:     level,
	}
}

func (b *TextBook) PrintInfo() {
	fmt.Printf("\nTitle: %s\nAuthor: %s\nPages: %d\nEditorial: %s\nLevel: %s\n",
		b.title, b.author, b.pages, b.editorial, b.level)
}
