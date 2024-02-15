package main

import "fmt"

type Animal struct {
	Name string
	Age  int
}

func (a *Animal) Speak() {
	fmt.Println("I am an animal")
}

type Dog struct {
	Animal
	Breed string
}

func (d *Dog) Speak() {
	fmt.Println("I am a dog")
}

func main() {
	dog := Dog{Animal{"Fido", 5}, "Doberman"}
	fmt.Println(dog.Name)
	fmt.Println(dog.Breed)
	dog.Speak()

	animal := Animal{"Fido", 5}
	animal.Speak()
}
