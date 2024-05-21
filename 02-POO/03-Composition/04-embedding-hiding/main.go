package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age  uint
}

func NewPerson(name string, age uint) Person {
	return Person{name, age}
}

func (p Person) Greet() string {
	return fmt.Sprintf("Hello, my name is %s and I'm a person", p.Name)
}

type Human struct {
	Age      uint
	Children uint
}

func NewHuman(age, children uint) Human {
	return Human{age, children}
}

type Employee struct {
	Person
	Human
	Salary float64
}

func NewEmployee(name string, age, children uint, salary float64) Employee {
	return Employee{NewPerson(name, age), NewHuman(age, children), salary}
}

func (e Employee) Payroll() string {
	return fmt.Sprintf("%s and my salary is %.2f", e.Greet(), e.Salary)
}

func (e Employee) Greet() string {
	return fmt.Sprintf("Hello, my name is %s and I'm an employee", e.Name)
}

func main() {
	p := NewPerson("John", 30)
	fmt.Println(p.Greet())

	e := NewEmployee("Jane", 25, 2, 1000)
	// fmt.Println(e.Age) // ambiguous selector e.Age
	fmt.Println(e.Greet())
	fmt.Println(e.Person.Greet())
	fmt.Println(e.Payroll())
}
