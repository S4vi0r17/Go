package generics

import (
	"fmt"
)

// type Product struct {
// 	Id    uint
// 	Name  string
// 	Price float64
// }

// type Product2 struct {
// 	Id    string
// 	Name  string
// 	Price float64
// }

type Product[T uint | string] struct {
	Id    T
	Name  string
	Price float64
}

func GenericTypes() {
	product1 := Product[uint]{Id: 1, Name: "Product 1", Price: 100.0}
	fmt.Println(product1)

	product2 := Product[string]{Id: "1", Name: "Product 2", Price: 200.0}
	fmt.Println(product2)
}
