package main

import "fmt"

func Array() {
	var array [5]string

	array[0] = "🦎"
	array[1] = "🐶"
	array[2] = "🐢"
	array[3] = "🐈"
	array[4] = "💩"

	fmt.Println(array)

	// Other way to declare an array
	// array2 := [5]int{1, 2, 3, 4, 5}
	array2 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array2)
}
