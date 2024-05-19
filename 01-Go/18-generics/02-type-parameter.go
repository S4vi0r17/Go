package generics

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

type MyInt int

func TypeParameter() {
	fmt.Println(sum(1, 2, 3))
	fmt.Println(sum(1.1, 2.2, 3.3))

	var a MyInt = 1
	var b MyInt = 2

	fmt.Println(sum(a, b))

	fmt.Println(multiply(1, 2, 3))
}

// The sum function can accept int or float64 values.
func sum[T ~int | float64](nums ...T) T {
	var sum T
	for _, n := range nums {
		sum += n
	}
	return sum
}

// type Number interface {
// 	~int | ~float64 | ~float32 | ~uint
// }

// func multiply[T Number](nums ...T) T {
// 	var product T
// 	for _, n := range nums {
// 		product *= n
// 	}
// 	return product
// }

/*
	go get golang.org/x/exp/constraints
*/

func multiply[T constraints.Integer | constraints.Float](nums ...T) T {
	var product T
	for _, n := range nums {
		product *= n
	}
	return product
}
