package modules

import (
	"fmt"
)

type TestFunc struct {
}

func (tFunction TestFunc) TestRun(myNum int) string {

	fmt.Println(myNum)

	return "aaaaaa"
}
