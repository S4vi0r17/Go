package main

import "fmt"

type Person struct {
	name string
	age  int
}

func Structs() {
	p1 := Person{"Alice", 25} // This is the only way to declare a struct without using the field names
	p2 := Person{name: "Bob", age: 30}
	p3 := Person{name: "Charlie"}
	p4 := Person{}

	fmt.Println(p1, p2, p3, p4)
	fmt.Println(p1.name, p2.age)

	// Pointers to structs
	// var p *Person
	// p = &p1
	p := &p1
	fmt.Println(p)
	fmt.Println(*p)

	// (*p).name = "David"
	p.name = "David" // This is a shorthand for the above line
	fmt.Println(p1)
}
