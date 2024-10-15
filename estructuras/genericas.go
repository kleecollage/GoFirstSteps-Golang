package estructuras

import "fmt"

// modelo de la bd
// estructura generica
type Product[T uint | string] struct {
	Id          T
	Description string
	Price       float32
}

func Genericas() {
	product1 := Product[uint]{1, "Zapatos", 50}
	product2 := Product[string]{"ABC-123", "Zapatos", 50}

	fmt.Println(product1, product2)
}
