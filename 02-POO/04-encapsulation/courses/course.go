package courses

import (
	"fmt"
)

type Course struct {
	Name    string
	Slug    string
	Skills  []string
	Price   float64
	IsFree  bool
	UserIDs []uint
	Classes map[uint]string
}

func (c *Course) PrintClasses() {
	text := "Las clases son:\n"
	for _, class := range c.Classes {
		text += class + "\n"
	}
	fmt.Println(text)
}

func (c *Course) ChangePrice(price float64) {
	c.Price = price
}
