package main

import (
	"context"
	"fmt"
	"gin_sample_web_app/config"
	"gin_sample_web_app/modules"
)

func main() {
	// fmt.Println(111)

	testFuncStruct := modules.TestFunc{}
	fmt.Println(context.TODO())
	fmt.Print(testFuncStruct.TestRun(111))

	r := config.DefineRoutes()
	r.Run(":8080")
}
