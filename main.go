package main

import (
	"fmt"
	"gin_sample_web_app/modules"
)

func main() {
	fmt.Println(111)

	testFuncStruct := modules.TestFunc{}

	fmt.Print(testFuncStruct.TestRun(111))
}
