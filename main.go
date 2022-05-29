package main

import "gin_sample_web_app/modules"

func IsOne(i int) bool {
	if i == 1 {
		return true
	} else {
		return false
	}
}

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

	// vs := []modules.Stringfy{
	// 	&modules.Person{Name: "Taro", Age: 21},
	// 	&modules.Car{Number: "123345", Model: "111111"},
	// }

	// for _, v := range vs {
	// 	fmt.Println(v.ToString())
	// }

	// err := modules.RaiseError()
	// fmt.Println(err.Error())

	// e, ok := err.(*modules.MyError)
	// if ok {
	// 	fmt.Println(e.ErrCode)
	// }

	// s := []int{2, 3, 4, 5, 6}

	// fmt.Println(modules.Average(s))

	// u := new(modules.User)
	// u.Id = 1
	// u.Name = "test User"
	// u.Email = "example@example.com"
	// u.Created = time.Now()

	// bs, err := json.Marshal(u)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(string(bs))

	// fmt.Printf("%T\n", bs)

	// u2 := new(modules.User)

	// if err := json.Unmarshal(bs, &u2); err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(u2)

	// modules.PrintError(modules.DoError(modules.HOGEHOGE))
	// modules.PrintError(modules.DoError(modules.FUGAFUGA))
	// modules.PrintError(modules.DoError(modules.HOGEFUGA))
	// modules.CallPanic()
	// modules.PanicTest()
	// modules.RecoverPanic()
	modules.RecoverPanicTestFunc()
}
