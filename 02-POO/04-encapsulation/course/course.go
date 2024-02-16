package course

import "fmt"

type course struct {
	name    string
	Slug    string
	Skills  []string
	price   float64
	isFree  bool
	userIDs []uint
	classes map[uint]string
}

func New(name, slug string, skills []string, price float64, isFree bool, userIDs []uint, classes map[uint]string) *course {
	if price == 0 {
		price = 30
	}
	return &course{
		name:    name,
		Slug:    slug,
		Skills:  skills,
		price:   price,
		isFree:  isFree,
		userIDs: userIDs,
		classes: classes,
	}
}

func (c *course) SetClasses(classes map[uint]string) {
	c.classes = classes
}

func (c *course) SetName(name string) {
	c.name = name
}

func (c *course) Name() string {
	return c.name
}

func (c *course) SetPrice(price float64) {
	c.price = price
}

func (c *course) Price() float64 {
	return c.price
}

func (c *course) SetIsFree(isFree bool) {
	c.isFree = isFree
}

func (c *course) IsFree() bool {
	return c.isFree
}

func (c *course) SetUserIDs(userIDs []uint) {
	c.userIDs = userIDs
}

func (c *course) UserIDs() []uint {
	return c.userIDs
}

func (c *course) PrintClasses() {
	text := "Las clases son:\n"
	for _, class := range c.classes {
		text += class + "\n"
	}
	fmt.Println(text)
}
