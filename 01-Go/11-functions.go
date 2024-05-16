package main

import "fmt"

func Functions3() {
	nums := []int{1, 10, 70, 20, 15}
	evens := filter(nums, func(n int) bool {
		return n%2 == 0
	})
	fmt.Println(evens)

	x := hello("Gustavo")
	fmt.Println(x("Benites"))

	// x := hello("Gustavo")("Benites")
	// fmt.Println(x)
}

func filter(nums []int, callback func(int) bool) []int {
	var result []int
	for _, num := range nums {
		if callback(num) {
			result = append(result, num)
		}
	}
	return result
}

func hello(name string) func(string) string {
	return func(lastName string) string {
		return "Hello, " + name + " " + lastName + "! How are you?"
	}
}
