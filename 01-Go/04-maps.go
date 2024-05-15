package main

import "fmt"

func Maps() {
	animals := map[string]string{
		"cat":     "meow",
		"dog":     "woof",
		"cow":     "moo",
		"pig":     "oink",
		"duck":    "quack",
		"sheep":   "baa",
	}
	fmt.Println(animals)

	// Other way to declare a map
	var fruits = make(map[string]string)
	fruits["apple"] = "🍎"
	fruits["banana"] = "🍌"
	fruits["peach"] = "🍑"
	fmt.Println(fruits)

	// Delete an element from a map
	delete(animals, "sheep")
	fmt.Println(animals)

	// Get a value from a map
	animal := "cat"
	sound := animals[animal]
	fmt.Println(animal, ":", sound)

	// Check if a key exists in a map
	animal = "sheep"

	sound, ok := animals[animal]
	if ok {
		fmt.Println(animal, ":", sound)
	} else {
		fmt.Println(animal, "is not in the map")
	}

	// In one line
	if sound, ok := animals[animal]; ok {
		fmt.Println(animal, ":", sound)
	} else {
		fmt.Println(animal, "is not in the map")
	}
}
