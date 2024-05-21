package main

import (
	"fmt"
)

type Greeter interface {
	Greet()
}

type Byer interface {
	Bye()
}

type GreeterByer interface {
	Greeter
	Byer
}

type Person struct {
	Name string
}

func (p Person) Greet() {
	fmt.Println("Hello, my name is", p.Name)
}

func (p Person) Bye() {
	fmt.Println("Goodbye, my name is", p.Name)
}

func (p Person) String() string {
	return "Hello, my name is " + p.Name
}

type Text string

func (t Text) Greet() {
	fmt.Println("Hello, I'm a text message:", t)
}

func (t Text) Bye() {
	fmt.Println("Goodbye, I'm a text message:", t)
}

func GreetAll(greeters ...Greeter) {
	for _, greeter := range greeters {
		greeter.Greet()
	}
}

func ByeAll(byers ...Byer) {
	for _, byer := range byers {
		byer.Bye()
	}
}

func GreetByeAll(greeterByers ...GreeterByer) {
	for _, greeterByer := range greeterByers {
		greeterByer.Greet()
		greeterByer.Bye()
	}
}

func main() {
	// var g Greeter = Person{"Eder"}
	// g.Greet()

	// g = Text("Blah blah blah...")
	// g.Greet()

	p := Person{"Eder"}
	t := Text("Blah blah blah...")
	GreetAll(p, t)
	ByeAll(p, t)

	GreetByeAll(p, t)

	fmt.Println(p)
}
