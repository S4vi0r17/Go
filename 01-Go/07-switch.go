package main

import "fmt"

func Switch() {
	emoji := "🐶"
	switch emoji {
	case "🐶":
		fmt.Println("It's a dog")
	case "🐱":
		fmt.Println("It's a cat")
	default:
		fmt.Println("It's not a dog or a cat")
	}

	// switch with a short statement
	switch a := 30; {
	case a < 20:
		fmt.Println("a is less than 20")
	default:
		fmt.Println("a is greater than or equal to 20")
	}

	emoji2 := "🐱"
	switch {
	case emoji2 == "🐶" || emoji2 == "🐱":
		fmt.Println("It's a dog or a cat")
	default:
		fmt.Println("It's not a dog or a cat")
	}
}
