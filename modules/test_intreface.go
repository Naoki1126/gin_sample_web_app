package modules

import "fmt"

type Stringfy interface {
	ToString() string
}

type Person struct {
	Name string
	Age  int
}

func (p *Person) ToString() string {
	return fmt.Sprintf("Name=%v, Age=%v", p.Name, p.Age)
}

type Car struct {
	Number string
	Model  string
}

func (c *Car) ToString() string {
	return fmt.Sprintf("Number=%v, Model=%v", c.Number, c.Model)
}

type TestOverRideInterFace interface {
	error
	ShowText() string
	ShowNumber() int
}

func (c *Car) DoError() error {
	return fmt.Errorf("Error")
}

func (c *Car) ShowText() string {
	fmt.Println("aaaa")
	return "test"
}

func (c *Car) ShowNumber() int {
	fmt.Println("11111")
	return 111
}
