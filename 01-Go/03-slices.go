package main

import "fmt"

func Slices() {
	// Slices are like arrays, but they can grow and shrink, also they don't data, they are just references(like pointers) to an array
	set := []int{1, 2, 3, 4, 5}
	evens := set[1:3]
	fmt.Println(evens) // [2, 3]
	evens[0] = 10
	fmt.Println(evens) // [10, 3]
	fmt.Println(set)   // [1, 10, 3, 4, 5]

	set2 := []string{"🐈", "🐁", "🐇", "🐦", "🦋", "🦆"}
	// terrarium := set2[0:3]
	terrarium := set2[:3]
	fmt.Println(terrarium) // [🐈, 🐁, 🐇]

	fly := set2[3:]
	fmt.Println(fly) // [🐦, 🦋, 🦆]

	fly[2] = "🦅"
	fmt.Println(set2) // [🐈, 🐁, 🐇, 🐦, 🦋, 🦅]

	fmt.Println()

	food := [5]string{"🍕", "🍓", "🍌", "🍑", "🍔"}
	fruits := food[1:4] // [🍓, 🍌, 🍑]

	fmt.Println("food: ", food)
	fmt.Println("fruits: ", fruits)
	fmt.Println("len(fruits): ", len(fruits))
	fmt.Println("cap(fruits): ", cap(fruits)) // from the first element in the slice to the last element in the array

	fruits = append(fruits, "🍇", "🍐", "🥑")
	// * If we append more elements than the capacity of the slice, the slice will be reallocated and the reference will change *

	fmt.Println("\nfood: ", food)
	fmt.Println("fruits: ", fruits)
	fmt.Println("len(fruits): ", len(fruits))
	fmt.Println("cap(fruits): ", cap(fruits))

	fmt.Println("\nfood: ", food)

	fmt.Println()

	// Other way to declare a slice
	emoji := make([]string, 3)
	/*
		s := make([]int, 5) // Crea un slice de enteros con longitud 5 y capacidad 5
		s := make([]int, 5, 10) // Crea un slice de enteros con longitud 5 y capacidad 10
	*/
	emoji[0] = "🥵"
	emoji[1] = "🥶"
	emoji[2] = "🤮"
	emoji = append(emoji, "🤯", "🤬")
	fmt.Println(emoji)
	fmt.Println("len(emoji): ", len(emoji))
	fmt.Println("cap(emoji): ", cap(emoji))
}
