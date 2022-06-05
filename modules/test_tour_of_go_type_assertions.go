package modules

import "fmt"

type TFunc interface {
	error
	Do()
	ToDo()
}

type TFuncHoge struct {
	Err error
}

func (t TFuncHoge) Do() {
	fmt.Println("Do")
}

func (t TFuncHoge) Todo() {
	fmt.Println("Todo")
}

func (t TFuncHoge) Error() string {
	return "HOGEHOGEError"
}
