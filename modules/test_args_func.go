package modules

import (
	"fmt"
	"strconv"
)

type TestArgsFunc struct {
}

type TestArgsFunc2 struct {
	Name string
	age  int
}

func (ta TestArgsFunc) Func2(num int) (*TestArgsFunc2, error) {
	var test1 TestArgsFunc
	fmt.Println(test1)

	test2 := TestArgsFunc2{"aaaa", 211111}

	str := "1111111"

	cvStr, err := strconv.Atoi(str)
	if err != nil {
		return nil, err
	}

	fmt.Println(cvStr)

	return &test2, nil

}
