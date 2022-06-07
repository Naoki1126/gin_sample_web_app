package modules

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type MyString string

type MyStringStruct struct {
	X MyString
	Y string
}

func (m MyStringStruct) MyStringFunc() {
	fmt.Println(m.X)
}

type MyInterFace interface {
	GetInfo()
}

type BookDatabase struct {
}

func (b *BookDatabase) GetInfo() {
	fmt.Println("book")
}

type Shop struct{}

func (s *Shop) GetInfo() {
	fmt.Println("show")
}
