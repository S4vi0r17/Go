package main

import "fmt"

func For() {
	// Classic for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// While loop
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	// Infinite loop
	// for {
	// 	fmt.Println("Infinite loop")
	// }

	// Looping through a slice
	slice := []int{1, 2, 3}
	for i, v := range slice {
		fmt.Println("Index:", i, "Value:", v)
	}

	// Looping through a map
	m := map[string]int{"a": 1, "b": 2}
	for k, v := range m {
		fmt.Println("Key:", k, "Value:", v)
	}

	// Looping through a string
	for i, v := range "Hello" {
		fmt.Println("Index:", i, "Value:", string(v))
	}

	// Looping through a channel
	ch := make(chan int)
	go func() {
		ch <- 1
		ch <- 2
		ch <- 3
		close(ch)
	}()
	for v := range ch {
		fmt.Println("Value:", v)
	}
}
