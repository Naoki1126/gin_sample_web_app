package main

import (
	"fmt"
	"gin_sample_web_app/modules"
)

func main() {
	// fmt.Println(111)

	// testFuncStruct := modules.TestFunc{}
	// fmt.Println(context.TODO())
	// fmt.Print(testFuncStruct.TestRun(111))

	// r := config.DefineRoutes()
	// r.Run(":8080")

	// p := &operations.PostUserParams{
	// 	Body: &models.User{
	// 		ID:   1,
	// 		Name: "name_log",
	// 	},
	// }

	// p.SetTimeout(10 * time.Second)

	// transport := httptransport.New("localhost:8000", "api", nil)
	// client := apiclient.New(transport, strfmt.Default)

	// res, _ := client.Operations.PostUser(p)

	// log.Printf("&#v\n", res.Error())

	vs := []modules.Stringfy{
		&modules.Person{Name: "Taro", Age: 21},
		&modules.Car{Number: "123345", Model: "111111"},
	}

	for _, v := range vs {
		fmt.Println(v.ToString())
	}

}
